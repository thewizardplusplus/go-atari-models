package sgf

import (
	"errors"
	"fmt"
	"strconv"

	models "github.com/thewizardplusplus/go-atari-models"
)

const (
	minColumnCode = 97
)

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
