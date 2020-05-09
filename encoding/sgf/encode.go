package sgf

import (
	models "github.com/thewizardplusplus/go-atari-models"
)

// EncodeAxis ...
//
// It encodes an axis in accordance
// with SGF (FF[4]).
//
// It performs the inverse transformation
// for DecodeAxis.
//
// See DecodeAxis for details.
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
//
// It encodes a point in accordance
// with SGF (FF[4]).
//
// It performs the inverse transformation
// for DecodePoint.
//
// See DecodeAxis for details.
//
// It panics, if the axis out of ranges.
//
func EncodePoint(point models.Point) string {
	column := EncodeAxis(point.Column)
	row := EncodeAxis(point.Row)
	return string([]byte{column, row})
}
