package sgf

import (
	"strconv"

	models "github.com/thewizardplusplus/go-atari-models"
)

const (
	minColumnCode = 97
)

// EncodeAxis ...
//
// It performs the inverse transformation
// for DecodeAxis. See the latter
// for details.
//
// It panics, if the axis out of ranges.
//
func EncodeAxis(axis int) byte {
	var symbol byte
	switch {
	case axis >= 0 && axis < alphabetLength:
		symbol = byte(axis) + 'a'
	case axis >= alphabetLength &&
		axis < 2*alphabetLength:
		symbol =
			byte(axis) - alphabetLength + 'A'
	default:
		panic(
			"sgf.EncodeAxis: " +
				"axis out of ranges",
		)
	}

	return symbol
}

// EncodePoint ...
func EncodePoint(point models.Point) string {
	column :=
		string(point.Column + minColumnCode)
	row := strconv.Itoa(point.Row + 1)
	return column + row
}
