package sgf

import (
	"strconv"

	models "github.com/thewizardplusplus/go-atari-models"
)

const (
	minColumnCode = 97
)

// EncodePoint ...
func EncodePoint(point models.Point) string {
	column :=
		string(point.Column + minColumnCode)
	row := strconv.Itoa(point.Row + 1)
	return column + row
}
