package sgf

import (
	"testing"

	models "github.com/thewizardplusplus/go-atari-models"
)

func TestEncodeAxis(test *testing.T) {
	type args struct {
		axis int
	}
	type data struct {
		args       args
		wantSymbol byte
		wantPanic  bool
	}

	for _, data := range []data{
		data{
			args:       args{4},
			wantSymbol: 'e',
			wantPanic:  false,
		},
		data{
			args:       args{30},
			wantSymbol: 'E',
			wantPanic:  false,
		},
		data{
			args:       args{-1},
			wantSymbol: 0,
			wantPanic:  true,
		},
	} {
		var gotSymbol byte
		var hasPanic bool
		func() {
			defer func() {
				if err := recover(); err != nil {
					hasPanic = true
				}
			}()

			gotSymbol = EncodeAxis(data.args.axis)
		}()

		if gotSymbol != data.wantSymbol {
			test.Fail()
		}

		if hasPanic != data.wantPanic {
			test.Fail()
		}
	}
}

func TestEncodePoint(test *testing.T) {
	type args struct {
		point models.Point
	}
	type data struct {
		args      args
		wantText  string
		wantPanic bool
	}

	for _, data := range []data{
		data{
			args: args{
				point: models.Point{
					Column: 4,
					Row:    30,
				},
			},
			wantText:  "eE",
			wantPanic: false,
		},
		data{
			args: args{
				point: models.Point{
					Column: -1,
					Row:    4,
				},
			},
			wantText:  "",
			wantPanic: true,
		},
		data{
			args: args{
				point: models.Point{
					Column: 4,
					Row:    -1,
				},
			},
			wantText:  "",
			wantPanic: true,
		},
	} {
		var gotText string
		var hasPanic bool
		func() {
			defer func() {
				if err := recover(); err != nil {
					hasPanic = true
				}
			}()

			gotText = EncodePoint(data.args.point)
		}()

		if gotText != data.wantText {
			test.Fail()
		}

		if hasPanic != data.wantPanic {
			test.Fail()
		}
	}
}
