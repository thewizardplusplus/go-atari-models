package atarimodels

import (
	"testing"
)

func TestColorNegative(test *testing.T) {
	type data struct {
		color Color
		want  Color
	}

	for _, data := range []data{
		data{
			color: Black,
			want:  White,
		},
		data{
			color: White,
			want:  Black,
		},
	} {
		got := data.color.Negative()

		if got != data.want {
			test.Fail()
		}
	}
}
