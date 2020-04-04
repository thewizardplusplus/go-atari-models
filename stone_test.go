package atarimodels

import (
	"reflect"
	"testing"
)

func TestStoneGroupMove(test *testing.T) {
	stones := make(stoneGroup)
	stones.Move(Move{Black, Point{2, 3}})
	stones.Move(Move{White, Point{3, 2}})

	expectedStones := stoneGroup{
		Point{2, 3}: Black,
		Point{3, 2}: White,
	}
	if !reflect.DeepEqual(
		stones,
		expectedStones,
	) {
		test.Fail()
	}
}

func TestStoneGroupCopy(test *testing.T) {
	stones := make(stoneGroup)
	stones.Move(Move{Black, Point{2, 3}})

	stonesCopy := stones.Copy()
	stones.Move(Move{White, Point{3, 2}})

	expectedStonesCopy := stoneGroup{
		Point{2, 3}: Black,
	}
	if !reflect.DeepEqual(
		stonesCopy,
		expectedStonesCopy,
	) {
		test.Fail()
	}
}

func TestStoneGroupCopyByPoints(
	test *testing.T,
) {
	type args struct {
		points []Point
	}
	type data struct {
		stones  stoneGroup
		args    args
		process func(stones stoneGroup)
		want    stoneGroup
	}

	for _, data := range []data{
		data{
			stones: stoneGroup{
				Point{0, 0}: Black,
				Point{2, 0}: White,
				Point{0, 2}: Black,
				Point{2, 2}: White,
			},
			args: args{
				points: []Point{
					Point{0, 0},
					Point{2, 2},
				},
			},
			process: func(stones stoneGroup) {
				stones.Move(Move{
					Color: Black,
					Point: Point{0, 1},
				})
				stones.Move(Move{
					Color: White,
					Point: Point{2, 1},
				})
			},
			want: stoneGroup{
				Point{0, 0}: Black,
				Point{2, 2}: White,
			},
		},
		data{
			stones: stoneGroup{
				Point{0, 0}: Black,
				Point{2, 0}: White,
				Point{0, 2}: Black,
				Point{2, 2}: White,
			},
			args: args{
				points: []Point{
					Point{0, 0},
					Point{1, 0},
					Point{1, 2},
					Point{2, 2},
				},
			},
			process: func(stones stoneGroup) {
				stones.Move(Move{
					Color: Black,
					Point: Point{0, 1},
				})
				stones.Move(Move{
					Color: White,
					Point: Point{2, 1},
				})
			},
			want: stoneGroup{
				Point{0, 0}: Black,
				Point{2, 2}: White,
			},
		},
		data{
			stones: stoneGroup{
				Point{0, 0}: Black,
				Point{2, 0}: White,
				Point{0, 2}: Black,
				Point{2, 2}: White,
			},
			args: args{
				points: nil,
			},
			process: func(stones stoneGroup) {
				stones.Move(Move{
					Color: Black,
					Point: Point{0, 1},
				})
				stones.Move(Move{
					Color: White,
					Point: Point{2, 1},
				})
			},
			want: stoneGroup{},
		},
	} {
		got := data.stones.CopyByPoints(
			data.args.points,
		)
		data.process(data.stones)

		if !reflect.DeepEqual(
			got,
			data.want,
		) {
			test.Fail()
		}
	}
}
