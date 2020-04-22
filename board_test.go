package atarimodels

import (
	"reflect"
	"testing"
)

func TestNewBoard(test *testing.T) {
	board := NewBoard(Size{5, 5})

	expectedBoard := Board{
		size:   Size{5, 5},
		stones: make(StoneGroup),
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
		stones StoneGroup
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
				stones: StoneGroup{
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
				stones: StoneGroup{
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
		stones StoneGroup
	}
	type args struct {
		point Point
	}
	type data struct {
		fields                fields
		args                  args
		wantNeighbors         StoneGroup
		wantHasStoneLiberties bool
	}

	for _, data := range []data{
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
					Point{0, 0}: Black,
					Point{1, 0}: White,
				},
			},
			args: args{
				point: Point{0, 0},
			},
			wantNeighbors: StoneGroup{
				Point{1, 0}: White,
			},
			wantHasStoneLiberties: true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
					Point{2, 0}: Black,
					Point{3, 0}: White,
				},
			},
			args: args{
				point: Point{2, 0},
			},
			wantNeighbors: StoneGroup{
				Point{3, 0}: White,
			},
			wantHasStoneLiberties: true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
					Point{2, 2}: Black,
					Point{3, 2}: White,
					Point{2, 3}: White,
				},
			},
			args: args{
				point: Point{2, 2},
			},
			wantNeighbors: StoneGroup{
				Point{3, 2}: White,
				Point{2, 3}: White,
			},
			wantHasStoneLiberties: true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
					Point{2, 2}: Black,
				},
			},
			args: args{
				point: Point{2, 2},
			},
			wantNeighbors:         StoneGroup{},
			wantHasStoneLiberties: true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
					Point{2, 1}: White,
					Point{1, 2}: White,
					Point{2, 2}: Black,
					Point{3, 2}: White,
					Point{2, 3}: White,
				},
			},
			args: args{
				point: Point{2, 2},
			},
			wantNeighbors: StoneGroup{
				Point{2, 1}: White,
				Point{1, 2}: White,
				Point{3, 2}: White,
				Point{2, 3}: White,
			},
			wantHasStoneLiberties: false,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
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
			args: args{
				point: Point{2, 2},
			},
			wantNeighbors: StoneGroup{
				Point{2, 1}: White,
				Point{1, 2}: White,
				Point{3, 2}: White,
				Point{2, 3}: White,
			},
			wantHasStoneLiberties: false,
		},
	} {
		board := Board{
			size:   data.fields.size,
			stones: data.fields.stones,
		}
		gotNeighbors, gotHasStoneLiberties :=
			board.StoneNeighbors(data.args.point)

		if !reflect.DeepEqual(
			gotNeighbors,
			data.wantNeighbors,
		) {
			test.Fail()
		}
		if gotHasStoneLiberties !=
			data.wantHasStoneLiberties {
			test.Fail()
		}
	}
}

func TestBoardHasCapture(test *testing.T) {
	type fields struct {
		size   Size
		stones StoneGroup
	}
	type args struct {
		options []HasCaptureOption
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
				stones: StoneGroup{
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
				options: []HasCaptureOption{
					WithColor(Black),
				},
			},
			wantColor: 0,
			wantOk:    false,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
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
				options: []HasCaptureOption{
					WithColor(White),
				},
			},
			wantColor: 0,
			wantOk:    false,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
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
				options: []HasCaptureOption{
					WithColor(Black),
				},
			},
			wantColor: Black,
			wantOk:    true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
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
				options: []HasCaptureOption{
					WithColor(White),
				},
			},
			wantColor: 0,
			wantOk:    false,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
					Point{0, 0}: Black,
					Point{0, 1}: White,
					Point{1, 0}: White,
					Point{4, 3}: Black,
					Point{3, 4}: Black,
					Point{4, 4}: White,
				},
			},
			args: args{
				options: []HasCaptureOption{
					WithColor(Black),
				},
			},
			wantColor: Black,
			wantOk:    true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
					Point{0, 0}: Black,
					Point{0, 1}: White,
					Point{1, 0}: White,
					Point{4, 3}: Black,
					Point{3, 4}: Black,
					Point{4, 4}: White,
				},
			},
			args: args{
				options: []HasCaptureOption{
					WithColor(White),
				},
			},
			wantColor: White,
			wantOk:    true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
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
				options: []HasCaptureOption{
					WithOrigin(Point{4, 2}),
				},
			},
			wantColor: 0,
			wantOk:    false,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
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
				options: []HasCaptureOption{
					WithOrigin(Point{4, 2}),
				},
			},
			wantColor: Black,
			wantOk:    true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
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
				options: []HasCaptureOption{
					WithColor(Black),
					WithOrigin(Point{4, 2}),
				},
			},
			wantColor: 0,
			wantOk:    false,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
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
				options: []HasCaptureOption{
					WithColor(White),
					WithOrigin(Point{4, 2}),
				},
			},
			wantColor: 0,
			wantOk:    false,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
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
				options: []HasCaptureOption{
					WithColor(Black),
					WithOrigin(Point{4, 2}),
				},
			},
			wantColor: Black,
			wantOk:    true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
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
				options: []HasCaptureOption{
					WithColor(White),
					WithOrigin(Point{4, 2}),
				},
			},
			wantColor: 0,
			wantOk:    false,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
					Point{0, 0}: Black,
					Point{0, 1}: White,
					Point{1, 0}: White,
				},
			},
			args: args{
				options: nil,
			},
			wantColor: Black,
			wantOk:    true,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
					Point{4, 3}: Black,
					Point{3, 4}: Black,
					Point{4, 4}: White,
				},
			},
			args: args{
				options: nil,
			},
			wantColor: White,
			wantOk:    true,
		},
	} {
		board := Board{
			size:   data.fields.size,
			stones: data.fields.stones,
		}
		gotColor, gotOk := board.HasCapture(
			data.args.options...,
		)

		if gotColor != data.wantColor {
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
		stones: StoneGroup{
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
		stones StoneGroup
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
				stones: StoneGroup{
					Point{2, 3}: Black,
					Point{3, 2}: White,
				},
			},
			args: args{
				move: Move{White, Point{-2, -2}},
			},
			want: ErrOutOfSize,
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
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
				stones: StoneGroup{
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
				stones: StoneGroup{
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
				stones: StoneGroup{
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
				stones: StoneGroup{
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
		stones StoneGroup
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
				stones: StoneGroup{
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
				stones: StoneGroup{
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

func TestBoardLegalMoves(test *testing.T) {
	type fields struct {
		size   Size
		stones StoneGroup
	}
	type args struct {
		previousMove Move
	}
	type data struct {
		fields    fields
		args      args
		wantMoves []Move
		wantErr   error
	}

	for _, data := range []data{
		data{
			fields: fields{
				size: Size{3, 3},
				stones: StoneGroup{
					Point{0, 2}: Black,
					Point{2, 0}: White,
				},
			},
			args: args{
				previousMove: Move{
					Color: Black,
					Point: Point{0, 2},
				},
			},
			wantMoves: []Move{
				Move{White, Point{0, 0}},
				Move{White, Point{1, 0}},
				Move{White, Point{0, 1}},
				Move{White, Point{1, 1}},
				Move{White, Point{2, 1}},
				Move{White, Point{1, 2}},
				Move{White, Point{2, 2}},
			},
			wantErr: nil,
		},
		data{
			fields: fields{
				size: Size{3, 3},
				stones: StoneGroup{
					Point{0, 0}: Black,
					Point{0, 1}: White,
					Point{1, 0}: White,
				},
			},
			args: args{
				previousMove: Move{
					Color: White,
					Point: Point{1, 0},
				},
			},
			wantMoves: nil,
			wantErr:   ErrAlreadyLoss,
		},
		data{
			fields: fields{
				size: Size{3, 3},
				stones: StoneGroup{
					Point{0, 0}: Black,
					Point{0, 1}: White,
					Point{1, 0}: White,
				},
			},
			args: args{
				previousMove: Move{
					Color: Black,
					Point: Point{0, 0},
				},
			},
			wantMoves: nil,
			wantErr:   ErrAlreadyWin,
		},
		data{
			fields: fields{
				size: Size{3, 3},
				stones: StoneGroup{
					Point{1, 0}: White,
					Point{2, 0}: White,
					Point{0, 1}: White,
					Point{1, 1}: White,
					Point{2, 1}: White,
					Point{0, 2}: White,
					Point{1, 2}: White,
				},
			},
			args: args{
				previousMove: Move{
					Color: White,
					Point: Point{1, 2},
				},
			},
			wantMoves: nil,
			wantErr:   ErrAlreadyLoss,
		},
		data{
			fields: fields{
				size: Size{3, 3},
				stones: StoneGroup{
					Point{0, 0}: White,
					Point{1, 0}: White,
					Point{2, 0}: White,
					Point{0, 1}: White,
					Point{1, 1}: White,
					Point{2, 1}: White,
					Point{0, 2}: White,
					Point{1, 2}: White,
					Point{2, 2}: White,
				},
			},
			args: args{
				previousMove: Move{
					Color: White,
					Point: Point{2, 2},
				},
			},
			wantMoves: nil,
			wantErr:   ErrAlreadyWin,
		},
	} {
		board := Board{
			size:   data.fields.size,
			stones: data.fields.stones,
		}
		gotMoves, gotErr := board.LegalMoves(
			data.args.previousMove,
		)

		if !reflect.DeepEqual(
			gotMoves,
			data.wantMoves,
		) {
			test.Fail()
		}
		if gotErr != data.wantErr {
			test.Fail()
		}
	}
}

func TestBoardHasLiberties(
	test *testing.T,
) {
	type fields struct {
		size   Size
		stones StoneGroup
	}
	type args struct {
		point Point
		chain pointGroup
	}
	type data struct {
		fields           fields
		args             args
		wantHasLiberties bool
		wantChain        []pointGroup
	}

	for _, data := range []data{
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
					Point{2, 2}: Black,
				},
			},
			args: args{
				point: Point{2, 2},
				chain: make(pointGroup),
			},
			wantHasLiberties: true,
			wantChain: []pointGroup{
				pointGroup{
					Point{2, 2}: struct{}{},
				},
			},
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
					Point{2, 2}: Black,
					Point{3, 2}: White,
					Point{2, 3}: White,
				},
			},
			args: args{
				point: Point{2, 2},
				chain: make(pointGroup),
			},
			wantHasLiberties: true,
			wantChain: []pointGroup{
				pointGroup{
					Point{2, 2}: struct{}{},
				},
			},
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
					Point{2, 1}: White,
					Point{1, 2}: White,
					Point{2, 2}: Black,
					Point{3, 2}: White,
					Point{2, 3}: White,
				},
			},
			args: args{
				point: Point{2, 2},
				chain: make(pointGroup),
			},
			wantHasLiberties: false,
			wantChain: []pointGroup{
				pointGroup{
					Point{2, 2}: struct{}{},
				},
			},
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
					Point{2, 1}: Black,
					Point{1, 2}: Black,
					Point{2, 2}: Black,
					Point{3, 2}: Black,
					Point{2, 3}: Black,
				},
			},
			args: args{
				point: Point{2, 2},
				chain: make(pointGroup),
			},
			wantHasLiberties: true,
			wantChain: []pointGroup{
				pointGroup{
					Point{2, 1}: struct{}{},
					Point{2, 2}: struct{}{},
				},
				pointGroup{
					Point{1, 2}: struct{}{},
					Point{2, 2}: struct{}{},
				},
				pointGroup{
					Point{2, 2}: struct{}{},
					Point{3, 2}: struct{}{},
				},
				pointGroup{
					Point{2, 2}: struct{}{},
					Point{2, 3}: struct{}{},
				},
			},
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
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
				chain: make(pointGroup),
			},
			wantHasLiberties: true,
			wantChain: []pointGroup{
				pointGroup{
					Point{2, 1}: struct{}{},
					Point{2, 2}: struct{}{},
				},
				pointGroup{
					Point{1, 2}: struct{}{},
					Point{2, 2}: struct{}{},
				},
				pointGroup{
					Point{2, 2}: struct{}{},
					Point{3, 2}: struct{}{},
				},
				pointGroup{
					Point{2, 2}: struct{}{},
					Point{2, 3}: struct{}{},
				},
			},
		},
		data{
			fields: fields{
				size: Size{5, 5},
				stones: StoneGroup{
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
				chain: make(pointGroup),
			},
			wantHasLiberties: false,
			wantChain: []pointGroup{
				pointGroup{
					Point{2, 1}: struct{}{},
					Point{1, 2}: struct{}{},
					Point{2, 2}: struct{}{},
					Point{3, 2}: struct{}{},
					Point{2, 3}: struct{}{},
				},
			},
		},
	} {
		board := Board{
			size:   data.fields.size,
			stones: data.fields.stones,
		}
		gotHasLiberties := board.hasLiberties(
			data.args.point,
			data.args.chain,
		)

		if gotHasLiberties !=
			data.wantHasLiberties {
			test.Fail()
		}

		var hasChainMatch bool
		for _, chain := range data.wantChain {
			hasChainMatch = reflect.DeepEqual(
				data.args.chain,
				chain,
			)
			if hasChainMatch {
				break
			}
		}
		if !hasChainMatch {
			test.Fail()
		}
	}
}
