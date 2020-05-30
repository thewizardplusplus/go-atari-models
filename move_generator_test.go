package atarimodels

import (
	"reflect"
	"testing"
)

func TestMoveGeneratorPseudolegalMoves(test *testing.T) {
	type args struct {
		storage StoneStorage
		color   Color
	}
	type data struct {
		args args
		want []Move
	}

	for _, data := range []data{
		{
			args: args{
				storage: Board{
					size:   Size{3, 3},
					stones: nil,
				},
				color: White,
			},
			want: []Move{
				{White, Point{0, 0}},
				{White, Point{1, 0}},
				{White, Point{2, 0}},
				{White, Point{0, 1}},
				{White, Point{1, 1}},
				{White, Point{2, 1}},
				{White, Point{0, 2}},
				{White, Point{1, 2}},
				{White, Point{2, 2}},
			},
		},
		{
			args: args{
				storage: Board{
					size: Size{3, 3},
					stones: StoneGroup{
						Point{0, 2}: Black,
						Point{2, 0}: White,
					},
				},
				color: White,
			},
			want: []Move{
				{White, Point{0, 0}},
				{White, Point{1, 0}},
				{White, Point{0, 1}},
				{White, Point{1, 1}},
				{White, Point{2, 1}},
				{White, Point{1, 2}},
				{White, Point{2, 2}},
			},
		},
		{
			args: args{
				storage: Board{
					size: Size{3, 3},
					stones: StoneGroup{
						Point{0, 1}: Black,
						Point{1, 0}: Black,
					},
				},
				color: White,
			},
			want: []Move{
				{White, Point{2, 0}},
				{White, Point{1, 1}},
				{White, Point{2, 1}},
				{White, Point{0, 2}},
				{White, Point{1, 2}},
				{White, Point{2, 2}},
			},
		},
	} {
		got := MoveGenerator{}.PseudolegalMoves(data.args.storage, data.args.color)

		if !reflect.DeepEqual(got, data.want) {
			test.Fail()
		}
	}
}

func TestMoveGeneratorLegalMoves(test *testing.T) {
	type args struct {
		storage      StoneStorage
		previousMove Move
	}
	type data struct {
		args      args
		wantMoves []Move
		wantErr   error
	}

	for _, data := range []data{
		{
			args: args{
				storage: Board{
					size: Size{3, 3},
					stones: StoneGroup{
						Point{0, 2}: Black,
						Point{2, 0}: White,
					},
				},
				previousMove: Move{
					Color: Black,
					Point: Point{0, 2},
				},
			},
			wantMoves: []Move{
				{White, Point{0, 0}},
				{White, Point{1, 0}},
				{White, Point{0, 1}},
				{White, Point{1, 1}},
				{White, Point{2, 1}},
				{White, Point{1, 2}},
				{White, Point{2, 2}},
			},
			wantErr: nil,
		},
		{
			args: args{
				storage: Board{
					size: Size{3, 3},
					stones: StoneGroup{
						Point{0, 0}: Black,
						Point{0, 1}: White,
						Point{1, 0}: White,
					},
				},
				previousMove: Move{
					Color: White,
					Point: Point{1, 0},
				},
			},
			wantMoves: nil,
			wantErr:   ErrAlreadyLoss,
		},
		{
			args: args{
				storage: Board{
					size: Size{3, 3},
					stones: StoneGroup{
						Point{0, 0}: Black,
						Point{0, 1}: White,
						Point{1, 0}: White,
					},
				},
				previousMove: Move{
					Color: Black,
					Point: Point{0, 0},
				},
			},
			wantMoves: nil,
			wantErr:   ErrAlreadyWin,
		},
		{
			args: args{
				storage: Board{
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
				previousMove: Move{
					Color: White,
					Point: Point{1, 2},
				},
			},
			wantMoves: nil,
			wantErr:   ErrAlreadyLoss,
		},
		{
			args: args{
				storage: Board{
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
				previousMove: Move{
					Color: White,
					Point: Point{2, 2},
				},
			},
			wantMoves: nil,
			wantErr:   ErrAlreadyWin,
		},
	} {
		gotMoves, gotErr :=
			MoveGenerator{}.LegalMoves(data.args.storage, data.args.previousMove)

		if !reflect.DeepEqual(gotMoves, data.wantMoves) {
			test.Fail()
		}
		if gotErr != data.wantErr {
			test.Fail()
		}
	}
}
