package atarimodels

import (
	"reflect"
	"testing"
)

func TestNewBoard(test *testing.T) {
	board := NewBoard(Size{5, 5})

	expectedBoard := Board{
		size:   Size{5, 5},
		stones: make(stoneGroup),
	}
	if !reflect.DeepEqual(
		board,
		expectedBoard,
	) {
		test.Fail()
	}
}

func TestBoardSize(test *testing.T) {
	board := NewBoard(Size{5, 5})
	size := board.Size()

	if !reflect.DeepEqual(size, Size{5, 5}) {
		test.Fail()
	}
}

func TestBoardStone(test *testing.T) {
	type fields struct {
		size   Size
		stones stoneGroup
	}
	type args struct {
		point Point
	}
	type data struct {
		fields    fields
		args      args
		wantColor Color
		wantOk    bool
	}

	for _, data := range []data{
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 3}: Black,
					Point{3, 2}: White,
				},
			},
			args:      args{Point{3, 2}},
			wantColor: White,
			wantOk:    true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 3}: Black,
					Point{3, 2}: White,
				},
			},
			args:      args{Point{2, 2}},
			wantColor: 0,
			wantOk:    false,
		},
	} {
		board := Board{
			size:   data.fields.size,
			stones: data.fields.stones,
		}
		gotColor, gotOk := board.
			Stone(data.args.point)

		if !reflect.DeepEqual(
			gotColor,
			data.wantColor,
		) {
			test.Fail()
		}
		if gotOk != data.wantOk {
			test.Fail()
		}
	}
}

func TestBoardApplyMove(test *testing.T) {
	board := NewBoard(Size{5, 5})
	for _, move := range []Move{
		Move{Black, Point{2, 3}},
		Move{White, Point{3, 2}},
	} {
		board = board.ApplyMove(move)
	}

	expectedBoard := Board{
		size: Size{5, 5},
		stones: stoneGroup{
			Point{2, 3}: Black,
			Point{3, 2}: White,
		},
	}
	if !reflect.DeepEqual(
		board,
		expectedBoard,
	) {
		test.Fail()
	}
}
