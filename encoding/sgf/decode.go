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

var (
	defaultSize = models.Size{
		Width:  5,
		Height: 5,
	}
	sizePattern = regexp.MustCompile(
		`\bSZ\[(\d+)(?::(\d+))?\]`,
	)
)

// DecodeAxis ...
//
// It decodes an axis in accordance
// with SGF (FF[4]).
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
		return 0,
			errors.New("symbol out of ranges")
	}

	return axis, nil
}

// DecodePoint ...
//
// It decodes a point in accordance
// with SGF (FF[4]).
//
// See DecodeAxis for details.
//
func DecodePoint(text string) (
	point models.Point,
	err error,
) {
	if len(text) != 2 {
		return models.Point{},
			errors.New("incorrect length")
	}

	column, err := DecodeAxis(text[0])
	if err != nil {
		return models.Point{}, fmt.Errorf(
			"incorrect column: %s",
			err,
		)
	}

	row, err := DecodeAxis(text[1])
	if err != nil {
		return models.Point{},
			fmt.Errorf("incorrect row: %s", err)
	}

	point = models.Point{
		Column: column,
		Row:    row,
	}
	return point, nil
}

// FindAndDecodeSize ...
//
// It finds in the provided text and decodes
// a size property in accordance with SGF
// (FF[4]).
//
func FindAndDecodeSize(text string) (
	size models.Size,
	err error,
) {
	match :=
		sizePattern.FindStringSubmatch(text)
	if match == nil {
		return defaultSize, nil
	}

	widthText, heightText :=
		match[1], match[2]
	width, err := strconv.Atoi(widthText)
	if err != nil {
		return models.Size{},
			fmt.Errorf("incorrect width: %s", err)
	}
	if !checkSideRange(width) {
		return models.Size{},
			errors.New("width out of range")
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
		return models.Size{}, fmt.Errorf(
			"incorrect height: %s",
			err,
		)
	}
	if !checkSideRange(height) {
		return models.Size{},
			errors.New("height out of range")
	}

	size = models.Size{
		Width:  width,
		Height: height,
	}
	return size, nil
}

// DecodeBoard ...
func DecodeBoard(text string) (
	board models.Board,
	err error,
) {
	return models.Board{}, err
}

func checkSideRange(side int) bool {
	return side >= 1 &&
		side <= 2*alphabetLength
}
