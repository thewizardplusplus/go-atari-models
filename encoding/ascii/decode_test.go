package ascii

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-atari-models"
)

func TestDecodePoint(test *testing.T) {
	type args struct {
		text string
	}
	type data struct {
		args      args
		wantPoint models.Point
		wantErr   bool
	}

	for _, data := range []data{
		data{
			args: args{"e2"},
			wantPoint: models.Point{
				Column: 4,
				Row:    1,
			},
			wantErr: false,
		},
		data{
			args:      args{"e"},
			wantPoint: models.Point{},
			wantErr:   true,
		},
		data{
			args:      args{"e23"},
			wantPoint: models.Point{},
			wantErr:   true,
		},
		data{
			args:      args{"\n2"},
			wantPoint: models.Point{},
			wantErr:   true,
		},
		data{
			args:      args{"e\n"},
			wantPoint: models.Point{},
			wantErr:   true,
		},
	} {
		gotPoint, gotErr :=
			DecodePoint(data.args.text)

		if !reflect.DeepEqual(
			gotPoint,
			data.wantPoint,
		) {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}
