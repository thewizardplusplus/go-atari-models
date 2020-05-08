package sgf

import (
	"errors"
	"fmt"

	models "github.com/thewizardplusplus/go-atari-models"
)

const (
	// it equals to the length of the low range
	highRangeShift = 'z' - 'a' + 1
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
		axis = int(symbol-'A') + highRangeShift
	default:
		return 0, errors.New("incorrect axis")
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
		return models.Point{},
			fmt.Errorf("incorrect column: %s", err)
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
