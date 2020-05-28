package sgf

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	models "github.com/thewizardplusplus/go-atari-models"
)

func TestDecodeColor(test *testing.T) {
	type args struct {
		symbol byte
	}
	type data struct {
		args      args
		wantColor models.Color
		wantErr   bool
	}

	for _, data := range []data{
		data{
			args:      args{'B'},
			wantColor: models.Black,
			wantErr:   false,
		},
		data{
			args:      args{'W'},
			wantColor: models.White,
			wantErr:   false,
		},
		data{
			args:      args{'\n'},
			wantColor: 0,
			wantErr:   true,
		},
	} {
		gotColor, gotErr :=
			DecodeColor(data.args.symbol)

		if !reflect.DeepEqual(
			gotColor,
			data.wantColor,
		) {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}

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

func TestFindAndDecodeMove(
	test *testing.T,
) {
	type args struct {
		text string
	}
	type data struct {
		args          args
		wantMove      models.Move
		wantLastIndex int
		wantOk        bool
	}

	for _, data := range []data{
		data{
			args: args{
				text: "(;FF[4]GN[test])",
			},
			wantMove:      models.Move{},
			wantLastIndex: 0,
			wantOk:        false,
		},
		data{
			args: args{
				text: "(;FF[4]B[eE]GN[test])",
			},
			wantMove: models.Move{
				Color: models.Black,
				Point: models.Point{
					Column: 4,
					Row:    30,
				},
			},
			wantLastIndex: 12,
			wantOk:        true,
		},
		data{
			args: args{
				text: "(;FF[4]AB[eE]GN[test])",
			},
			wantMove: models.Move{
				Color: models.Black,
				Point: models.Point{
					Column: 4,
					Row:    30,
				},
			},
			wantLastIndex: 13,
			wantOk:        true,
		},
		data{
			args: args{
				text: "(;FF[4]W[eE]GN[test])",
			},
			wantMove: models.Move{
				Color: models.White,
				Point: models.Point{
					Column: 4,
					Row:    30,
				},
			},
			wantLastIndex: 12,
			wantOk:        true,
		},
		data{
			args: args{
				text: "(;FF[4]AW[eE]GN[test])",
			},
			wantMove: models.Move{
				Color: models.White,
				Point: models.Point{
					Column: 4,
					Row:    30,
				},
			},
			wantLastIndex: 13,
			wantOk:        true,
		},
	} {
		gotMove, gotLastIndex, gotOk :=
			FindAndDecodeMove(data.args.text)

		if !reflect.DeepEqual(
			gotMove,
			data.wantMove,
		) {
			test.Fail()
		}

		if gotLastIndex != data.wantLastIndex {
			test.Fail()
		}

		if gotOk != data.wantOk {
			test.Fail()
		}
	}
}

func TestDecodeStoneStorage(
	test *testing.T,
) {
	type args struct {
		text    string
		factory StoneStorageFactory
	}
	type data struct {
		args        args
		wantStorage models.StoneStorage
		wantErr     bool
	}

	for _, data := range []data{
		data{
			args: args{
				text:    "(;FF[4]SZ[7:9]GN[test])",
				factory: models.NewBoard,
			},
			wantStorage: models.NewBoard(
				models.Size{
					Width:  7,
					Height: 9,
				},
			),
			wantErr: false,
		},
		data{
			args: args{
				text: "(;FF[4]SZ[7:9]GN[test]" +
					";B[aa](;W[gi]N[test]))",
				factory: models.NewBoard,
			},
			wantStorage: func() models.StoneStorage {
				board := models.NewBoard(
					models.Size{
						Width:  7,
						Height: 9,
					},
				)

				moves := []models.Move{
					models.Move{
						Color: models.Black,
						Point: models.Point{
							Column: 0,
							Row:    0,
						},
					},
					models.Move{
						Color: models.White,
						Point: models.Point{
							Column: 6,
							Row:    8,
						},
					},
				}
				for _, move := range moves {
					board = board.ApplyMove(move)
				}

				return board
			}(),
			wantErr: false,
		},
		data{
			args: args{
				text:    "(;FF[4]SZ[100:7]GN[test])",
				factory: models.NewBoard,
			},
			wantStorage: nil,
			wantErr:     true,
		},
	} {
		gotStorage, gotErr :=
			DecodeStoneStorage(
				data.args.text,
				data.args.factory,
			)

		if !reflect.DeepEqual(
			gotStorage,
			data.wantStorage,
		) {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}
