package sgf

import (
	"errors"
	"fmt"
	"strconv"

	models "github.com/thewizardplusplus/go-atari-models"
)

const (
	// it equals to the length of the low range
	highRangeShift = 'z' - 'a' + 1
	minColumnCode  = 97
)

// DecodeAxis ...
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
func DecodePoint(text string) (
	point models.Point,
	err error,
) {
	if len(text) != 2 {
		return models.Point{},
			errors.New("incorrect length")
	}

	column := int(text[0]) - minColumnCode
	if column < 0 {
		return models.Point{},
			errors.New("incorrect column")
	}

	row, err := strconv.Atoi(text[1:])
	if err != nil {
		return models.Point{},
			fmt.Errorf("incorrect row: %s", err)
	}
	row--

	point = models.Point{
		Column: column,
		Row:    row,
	}
	return point, nil
}
