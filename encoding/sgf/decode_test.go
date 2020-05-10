package sgf

import (
	"fmt"
	"reflect"
	"strings"
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

func TestFindAndDecodeSize(
	test *testing.T,
) {
	type args struct {
		text string
	}
	type data struct {
		args     args
		wantSize models.Size
		wantErr  bool
	}

	for _, data := range []data{
		data{
			args:     args{"(;FF[4]GN[test])"},
			wantSize: defaultSize,
			wantErr:  false,
		},
		data{
			args: args{"(;FF[4]SZ[7]GN[test])"},
			wantSize: models.Size{
				Width:  7,
				Height: 7,
			},
			wantErr: false,
		},
		data{
			args: args{"(;FF[4]SZ[7:9]GN[test])"},
			wantSize: models.Size{
				Width:  7,
				Height: 9,
			},
			wantErr: false,
		},
		data{
			args: args{
				text: "(;FF[4]SZ[23:42]GN[test])",
			},
			wantSize: models.Size{
				Width:  23,
				Height: 42,
			},
			wantErr: false,
		},
		data{
			args: args{
				text: "(;FF[4]SZ[100:7]GN[test])",
			},
			wantSize: models.Size{},
			wantErr:  true,
		},
		data{
			args: args{
				text: "(;FF[4]SZ[7:100]GN[test])",
			},
			wantSize: models.Size{},
			wantErr:  true,
		},
		data{
			args: args{
				text: fmt.Sprintf(
					"(;FF[4]SZ[%s:7]GN[test])",
					strings.Repeat("9", 23),
				),
			},
			wantSize: models.Size{},
			wantErr:  true,
		},
		data{
			args: args{
				text: fmt.Sprintf(
					"(;FF[4]SZ[7:%s]GN[test])",
					strings.Repeat("9", 23),
				),
			},
			wantSize: models.Size{},
			wantErr:  true,
		},
	} {
		gotSize, gotErr :=
			FindAndDecodeSize(data.args.text)

		if !reflect.DeepEqual(
			gotSize,
			data.wantSize,
		) {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}
