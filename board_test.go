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

		if gotColor != data.wantColor {
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
				size: Size{5, 5},
				stones: stoneGroup{
					Point{0, 0}: Black,
					Point{1, 0}: White,
				},
			},
			args:         args{Point{0, 0}},
			wantEmpty:    []Point{Point{0, 1}},
			wantOccupied: []Point{Point{1, 0}},
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 0}: Black,
					Point{3, 0}: White,
				},
			},
			args: args{Point{2, 0}},
			wantEmpty: []Point{
				Point{1, 0},
				Point{2, 1},
			},
			wantOccupied: []Point{Point{3, 0}},
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 2}: Black,
					Point{3, 2}: White,
					Point{2, 3}: White,
				},
			},
			args: args{Point{2, 2}},
			wantEmpty: []Point{
				Point{2, 1},
				Point{1, 2},
			},
			wantOccupied: []Point{
				Point{3, 2},
				Point{2, 3},
			},
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 2}: Black,
				},
			},
			args: args{Point{2, 2}},
			wantEmpty: []Point{
				Point{2, 1},
				Point{1, 2},
				Point{3, 2},
				Point{2, 3},
			},
			wantOccupied: nil,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 1}: White,
					Point{1, 2}: White,
					Point{2, 2}: Black,
					Point{3, 2}: White,
					Point{2, 3}: White,
				},
			},
			args:      args{Point{2, 2}},
			wantEmpty: nil,
			wantOccupied: []Point{
				Point{2, 1},
				Point{1, 2},
				Point{3, 2},
				Point{2, 3},
			},
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{1, 1}: White,
					Point{2, 1}: White,
					Point{3, 1}: White,
					Point{1, 2}: White,
					Point{2, 2}: Black,
					Point{3, 2}: White,
					Point{1, 3}: White,
					Point{2, 3}: White,
					Point{3, 3}: White,
				},
			},
			args:      args{Point{2, 2}},
			wantEmpty: nil,
			wantOccupied: []Point{
				Point{2, 1},
				Point{1, 2},
				Point{3, 2},
				Point{2, 3},
			},
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

func TestBoardStoneLiberties(
	test *testing.T,
) {
	type fields struct {
		size   Size
		stones stoneGroup
	}
	type args struct {
		point Point
		chain map[Point]struct{}
	}
	type data struct {
		fields        fields
		args          args
		wantLiberties int
		wantChain     map[Point]struct{}
	}

	for _, data := range []data{
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 2}: Black,
				},
			},
			args: args{
				point: Point{2, 2},
				chain: make(map[Point]struct{}),
			},
			wantLiberties: 4,
			wantChain: map[Point]struct{}{
				Point{2, 2}: struct{}{},
			},
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 2}: Black,
					Point{3, 2}: White,
					Point{2, 3}: White,
				},
			},
			args: args{
				point: Point{2, 2},
				chain: make(map[Point]struct{}),
			},
			wantLiberties: 2,
			wantChain: map[Point]struct{}{
				Point{2, 2}: struct{}{},
			},
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 1}: White,
					Point{1, 2}: White,
					Point{2, 2}: Black,
					Point{3, 2}: White,
					Point{2, 3}: White,
				},
			},
			args: args{
				point: Point{2, 2},
				chain: make(map[Point]struct{}),
			},
			wantLiberties: 0,
			wantChain: map[Point]struct{}{
				Point{2, 2}: struct{}{},
			},
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 1}: Black,
					Point{1, 2}: Black,
					Point{2, 2}: Black,
					Point{3, 2}: Black,
					Point{2, 3}: Black,
				},
			},
			args: args{
				point: Point{2, 2},
				chain: make(map[Point]struct{}),
			},
			wantLiberties: 12,
			wantChain: map[Point]struct{}{
				Point{2, 1}: struct{}{},
				Point{1, 2}: struct{}{},
				Point{2, 2}: struct{}{},
				Point{3, 2}: struct{}{},
				Point{2, 3}: struct{}{},
			},
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 1}: Black,
					Point{1, 2}: Black,
					Point{2, 2}: Black,
					Point{3, 2}: Black,
					Point{4, 2}: White,
					Point{2, 3}: Black,
					Point{3, 3}: White,
					Point{2, 4}: White,
				},
			},
			args: args{
				point: Point{2, 2},
				chain: make(map[Point]struct{}),
			},
			wantLiberties: 8,
			wantChain: map[Point]struct{}{
				Point{2, 1}: struct{}{},
				Point{1, 2}: struct{}{},
				Point{2, 2}: struct{}{},
				Point{3, 2}: struct{}{},
				Point{2, 3}: struct{}{},
			},
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 0}: White,
					Point{1, 1}: White,
					Point{2, 1}: Black,
					Point{3, 1}: White,
					Point{0, 2}: White,
					Point{1, 2}: Black,
					Point{2, 2}: Black,
					Point{3, 2}: Black,
					Point{4, 2}: White,
					Point{1, 3}: White,
					Point{2, 3}: Black,
					Point{3, 3}: White,
					Point{2, 4}: White,
				},
			},
			args: args{
				point: Point{2, 2},
				chain: make(map[Point]struct{}),
			},
			wantLiberties: 0,
			wantChain: map[Point]struct{}{
				Point{2, 1}: struct{}{},
				Point{1, 2}: struct{}{},
				Point{2, 2}: struct{}{},
				Point{3, 2}: struct{}{},
				Point{2, 3}: struct{}{},
			},
		},
	} {
		board := Board{
			size:   data.fields.size,
			stones: data.fields.stones,
		}
		gotLiberties := board.StoneLiberties(
			data.args.point,
			data.args.chain,
		)

		if gotLiberties != data.wantLiberties {
			test.Fail()
		}
		if !reflect.DeepEqual(
			data.args.chain,
			data.wantChain,
		) {
			test.Fail()
		}
	}
}

func TestBoardHasCapture(test *testing.T) {
	type fields struct {
		size   Size
		stones stoneGroup
	}
	type args struct {
		color []Color
	}
	type data struct {
		fields fields
		args   args
		want   bool
	}

	for _, data := range []data{
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 1}: Black,
					Point{1, 2}: Black,
					Point{2, 2}: Black,
					Point{3, 2}: Black,
					Point{4, 2}: White,
					Point{2, 3}: Black,
					Point{3, 3}: White,
					Point{2, 4}: White,
				},
			},
			args: args{[]Color{Black}},
			want: false,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 1}: Black,
					Point{1, 2}: Black,
					Point{2, 2}: Black,
					Point{3, 2}: Black,
					Point{4, 2}: White,
					Point{2, 3}: Black,
					Point{3, 3}: White,
					Point{2, 4}: White,
				},
			},
			args: args{[]Color{White}},
			want: false,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 0}: White,
					Point{1, 1}: White,
					Point{2, 1}: Black,
					Point{3, 1}: White,
					Point{0, 2}: White,
					Point{1, 2}: Black,
					Point{2, 2}: Black,
					Point{3, 2}: Black,
					Point{4, 2}: White,
					Point{1, 3}: White,
					Point{2, 3}: Black,
					Point{3, 3}: White,
					Point{2, 4}: White,
				},
			},
			args: args{[]Color{Black}},
			want: true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{2, 0}: White,
					Point{1, 1}: White,
					Point{2, 1}: Black,
					Point{3, 1}: White,
					Point{0, 2}: White,
					Point{1, 2}: Black,
					Point{2, 2}: Black,
					Point{3, 2}: Black,
					Point{4, 2}: White,
					Point{1, 3}: White,
					Point{2, 3}: Black,
					Point{3, 3}: White,
					Point{2, 4}: White,
				},
			},
			args: args{[]Color{White}},
			want: false,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{0, 0}: Black,
					Point{0, 1}: White,
					Point{1, 0}: White,
					Point{4, 3}: Black,
					Point{3, 4}: Black,
					Point{4, 4}: White,
				},
			},
			args: args{[]Color{Black}},
			want: true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{0, 0}: Black,
					Point{0, 1}: White,
					Point{1, 0}: White,
					Point{4, 3}: Black,
					Point{3, 4}: Black,
					Point{4, 4}: White,
				},
			},
			args: args{[]Color{White}},
			want: true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{0, 0}: Black,
					Point{0, 1}: White,
					Point{1, 0}: White,
				},
			},
			args: args{nil},
			want: true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{4, 3}: Black,
					Point{3, 4}: Black,
					Point{4, 4}: White,
				},
			},
			args: args{nil},
			want: true,
		},
	} {
		board := Board{
			size:   data.fields.size,
			stones: data.fields.stones,
		}
		got :=
			board.HasCapture(data.args.color...)

		if got != data.want {
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
					Point{2, 1}: White,
					Point{1, 2}: White,
					Point{3, 2}: White,
					Point{2, 3}: White,
				},
			},
			args: args{
				move: Move{Black, Point{2, 2}},
			},
			want: ErrSelfcapture,
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
		// a move with self-capture is allowed,
		// if the opponent will be captured
		data{
			fields: fields{
				size: Size{5, 5},
				stones: stoneGroup{
					Point{1, 1}: Black,
					Point{2, 1}: White,
					Point{0, 2}: Black,
					Point{1, 2}: White,
					Point{3, 2}: White,
					Point{1, 3}: Black,
					Point{2, 3}: White,
				},
			},
			args: args{
				move: Move{Black, Point{2, 2}},
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

func TestBoardPseudolegalMoves(
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
		data{
			fields: fields{
				size: Size{3, 3},
				stones: stoneGroup{
					Point{0, 1}: Black,
					Point{1, 0}: Black,
				},
			},
			args: args{White},
			want: []Move{
				Move{White, Point{2, 0}},
				Move{White, Point{1, 1}},
				Move{White, Point{2, 1}},
				Move{White, Point{0, 2}},
				Move{White, Point{1, 2}},
				Move{White, Point{2, 2}},
			},
		},
	} {
		board := Board{
			size:   data.fields.size,
			stones: data.fields.stones,
		}
		got := board.PseudolegalMoves(
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
