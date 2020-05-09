package sgf

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-atari-models"
)

func TestDecodeAxis(test *testing.T) {
	type args struct {
		symbol byte
	}
	type data struct {
		args     args
		wantAxis int
		wantErr  bool
	}

	for _, data := range []data{
		data{
			args:     args{'e'},
			wantAxis: 4,
			wantErr:  false,
		},
		data{
			args:     args{'E'},
			wantAxis: 30,
			wantErr:  false,
		},
		data{
			args:     args{'\n'},
			wantAxis: 0,
			wantErr:  true,
		},
	} {
		gotAxis, gotErr :=
			DecodeAxis(data.args.symbol)

		if !reflect.DeepEqual(
			gotAxis,
			data.wantAxis,
		) {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}

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
			args: args{"eE"},
			wantPoint: models.Point{
				Column: 4,
				Row:    30,
			},
			wantErr: false,
		},
		data{
			args:      args{"e"},
			wantPoint: models.Point{},
			wantErr:   true,
		},
		data{
			args:      args{"eee"},
			wantPoint: models.Point{},
			wantErr:   true,
		},
		data{
			args:      args{"\ne"},
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
