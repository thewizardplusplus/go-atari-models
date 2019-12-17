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
