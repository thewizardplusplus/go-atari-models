package ascii

import (
	"testing"

	models "github.com/thewizardplusplus/go-atari-models"
)

func TestEncodePoint(test *testing.T) {
	type args struct {
		point models.Point
	}
	type data struct {
		args args
		want string
	}

	for _, data := range []data{
		data{
			args: args{
				point: models.Point{
					Column: 2,
					Row:    1,
				},
			},
			want: "c2",
		},
		data{
			args: args{
				point: models.Point{
					Column: 5,
					Row:    6,
				},
			},
			want: "f7",
		},
	} {
		got := EncodePoint(data.args.point)

		if got != data.want {
			test.Fail()
		}
	}
}
