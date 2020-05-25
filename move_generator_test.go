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
