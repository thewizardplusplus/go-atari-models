package atarimodels

import (
	"reflect"
	"testing"
)

func TestMoveGeneratorPseudolegalMoves(
	test *testing.T,
) {
	type args struct {
		storage StoneStorage
		color   Color
	}
	type data struct {
		args args
		want []Move
	}

	for _, data := range []data{
		data{
			args: args{
				storage: Board{
					size:   Size{3, 3},
					stones: nil,
				},
				color: White,
			},
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
				Move{White, Point{2, 0}},
				Move{White, Point{1, 1}},
				Move{White, Point{2, 1}},
				Move{White, Point{0, 2}},
				Move{White, Point{1, 2}},
				Move{White, Point{2, 2}},
			},
		},
	} {
		var generator MoveGenerator
		got := generator.PseudolegalMoves(
			data.args.storage,
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

func TestMoveGeneratorLegalMoves(
	test *testing.T,
) {
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
		data{
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
		data{
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
		data{
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
		data{
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
		var generator MoveGenerator
		gotMoves, gotErr :=
			generator.LegalMoves(
				data.args.storage,
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
