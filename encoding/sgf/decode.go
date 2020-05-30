package sgf

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	models "github.com/thewizardplusplus/go-atari-models"
)

const (
	alphabetLength = 'z' - 'a' + 1
)

// nolint: gochecknoglobals
var (
	defaultSize = models.Size{
		Width:  5,
		Height: 5,
	}
	sizePattern = regexp.MustCompile(`\bSZ\[(\d+)(?::(\d+))?\]`)
	movePattern = regexp.MustCompile(`\bA?([BW])\[([[:alpha:]]{2})\]`)
)

// StoneStorageFactory ...
type StoneStorageFactory func(size models.Size) models.StoneStorage

// DecodeColor ...
//
// It decodes a color in accordance with SGF (FF[4]).
//
func DecodeColor(symbol byte) (color models.Color, err error) {
	switch symbol {
	case 'B':
		color = models.Black
	case 'W':
		color = models.White
	default:
		return 0, errors.New("incorrect symbol")
	}

	return color, nil
}

// DecodeAxis ...
//
// It decodes an axis in accordance with SGF (FF[4]).
//
// Symbols from 'a' to 'z' maps to [0; 25].
//
// Symbols from 'A' to 'Z' maps to [26; 51].
//
func DecodeAxis(symbol byte) (int, error) {
	var axis int
	switch {
	case symbol >= 'a' && symbol <= 'z':
		axis = int(symbol - 'a')
	case symbol >= 'A' && symbol <= 'Z':
		axis = int(symbol-'A') + alphabetLength
	default:
		return 0, errors.New("symbol out of ranges")
	}

	return axis, nil
}

// DecodePoint ...
//
// It decodes a point in accordance with SGF (FF[4]).
//
// See DecodeAxis for details.
//
func DecodePoint(text string) (point models.Point, err error) {
	if len(text) != 2 {
		return models.Point{}, errors.New("incorrect length")
	}

	column, err := DecodeAxis(text[0])
	if err != nil {
		return models.Point{}, fmt.Errorf("incorrect column: %s", err)
	}

	row, err := DecodeAxis(text[1])
	if err != nil {
		return models.Point{}, fmt.Errorf("incorrect row: %s", err)
	}

	point = models.Point{
		Column: column,
		Row:    row,
	}
	return point, nil
}

// FindAndDecodeSize ...
//
// It finds in the provided text and decodes a size property in accordance
// with SGF (FF[4]).
//
// By default the resulting size equals to 5x5.
//
func FindAndDecodeSize(text string) (size models.Size, err error) {
	match := sizePattern.FindStringSubmatch(text)
	if match == nil {
		return defaultSize, nil
	}

	widthText, heightText := match[1], match[2]
	width, err := strconv.Atoi(widthText)
	if err != nil {
		return models.Size{}, fmt.Errorf("incorrect width: %s", err)
	}
	if !checkSideRange(width) {
		return models.Size{}, errors.New("width out of range")
	}
	if len(heightText) == 0 {
		// square board
		size = models.Size{
			Width:  width,
			Height: width,
		}
		return size, nil
	}

	height, err := strconv.Atoi(heightText)
	if err != nil {
		return models.Size{}, fmt.Errorf("incorrect height: %s", err)
	}
	if !checkSideRange(height) {
		return models.Size{}, errors.New("height out of range")
	}

	size = models.Size{
		Width:  width,
		Height: height,
	}
	return size, nil
}

// FindAndDecodeMove ...
//
// It finds in the provided text and decodes a move property in accordance
// with SGF (FF[4]).
//
// A tree structure is ignored, moves are searched simply sequentially
// in the provided text.
//
// Both move and setup properties are supported and are considered equivalent.
//
// See DecodeAxis for details.
//
func FindAndDecodeMove(text string) (move models.Move, lastIndex int, ok bool) {
	match := movePattern.FindStringSubmatchIndex(text)
	if match == nil {
		return models.Move{}, 0, false
	}

	color, _ := DecodeColor(text[match[2]])          // nolint: gosec
	point, _ := DecodePoint(text[match[4]:match[5]]) // nolint: gosec
	move = models.Move{
		Color: color,
		Point: point,
	}
	return move, match[1], true
}

// DecodeStoneStorage ...
//
// It decodes a stone storage in accordance with SGF (FF[4]).
//
// Size and move properties are supported.
//
// See FindAndDecodeSize and FindAndDecodeMove for details.
//
func DecodeStoneStorage(text string, factory StoneStorageFactory) (
	storage models.StoneStorage,
	err error,
) {
	size, err := FindAndDecodeSize(text)
	if err != nil {
		return nil, fmt.Errorf("incorrect size: %s", err)
	}

	storage = factory(size)
	for {
		move, lastIndex, ok := FindAndDecodeMove(text)
		if !ok {
			break
		}

		storage = storage.ApplyMove(move)
		text = text[lastIndex:]
	}

	return storage, nil
}

func checkSideRange(side int) bool {
	return side >= 1 && side <= 2*alphabetLength
}
