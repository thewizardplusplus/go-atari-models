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

func TestBoardStoneNeighbors(
	test *testing.T,
) {
	type fields struct {
		size   Size
		stones stoneGroup
	}
	type args struct {
		point Point
	}
	type data struct {
		fields       fields
		args         args
		wantEmpty    []Point
		wantOccupied []Point
	}

	for _, data := range []data{
		data{
			fields: fields{
				size:   Size{5, 5},
				stones: nil,
			},
			args:         args{Point{2, 3}},
			wantEmpty:    nil,
			wantOccupied: nil,
		},
	} {
		board := Board{
			size:   data.fields.size,
			stones: data.fields.stones,
		}
		gotEmpty, gotOccupied := board.
			StoneNeighbors(data.args.point)

		if !reflect.DeepEqual(
			gotEmpty,
			data.wantEmpty,
		) {
			test.Fail()
		}
		if !reflect.DeepEqual(
			gotOccupied,
			data.wantOccupied,
		) {
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

func TestBoardCheckMove(test *testing.T) {
	type fields struct {
		size   Size
		stones stoneGroup
	}
	type args struct {
		move Move
	}
	type data struct {
		fields fields
		args   args
		want   error
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
			args: args{
				move: Move{White, Point{-1, -1}},
			},
			want: ErrOutOfSize,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 3}: Black,
					Point{3, 2}: White,
				},
			},
			args: args{
				move: Move{White, Point{2, 3}},
			},
			want: ErrOccupiedPoint,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 3}: Black,
					Point{3, 2}: White,
				},
			},
			args: args{
				move: Move{White, Point{3, 2}},
			},
			want: ErrOccupiedPoint,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 3}: Black,
				},
			},
			args: args{
				move: Move{White, Point{3, 2}},
			},
			want: nil,
		},
	} {
		board := Board{
			size:   data.fields.size,
			stones: data.fields.stones,
		}
		got := board.CheckMove(data.args.move)

		if got != data.want {
			test.Fail()
		}
	}
}

func TestBoardMovesForColor(
	test *testing.T,
) {
	type fields struct {
		size   Size
		stones stoneGroup
	}
	type args struct {
		color Color
	}
	type data struct {
		fields fields
		args   args
		want   []Move
	}

	for _, data := range []data{
		data{
			fields: fields{
				size:   Size{3, 3},
				stones: nil,
			},
			args: args{White},
			want: []Move{
				Move{White, Point{0, 0}},
				Move{White, Point{1, 0}},
				Move{White, Point{2, 0}},
				Move{White, Point{0, 1}},
				Move{White, Point{1, 1}},
				Move{White, Point{2, 1}},
				Move{White, Point{0, 2}},
				Move{White, Point{1, 2}},
				Move{White, Point{2, 2}},
			},
		},
		data{
			fields: fields{
				size: Size{3, 3},
				stones: stoneGroup{
					Point{0, 2}: Black,
					Point{2, 0}: White,
				},
			},
			args: args{White},
			want: []Move{
				Move{White, Point{0, 0}},
				Move{White, Point{1, 0}},
				Move{White, Point{0, 1}},
				Move{White, Point{1, 1}},
				Move{White, Point{2, 1}},
				Move{White, Point{1, 2}},
				Move{White, Point{2, 2}},
			},
		},
	} {
		board := Board{
			size:   data.fields.size,
			stones: data.fields.stones,
		}
		got := board.MovesForColor(
			data.args.color,
		)

		if !reflect.DeepEqual(
			got,
			data.want,
		) {
			test.Fail()
		}
	}
}
