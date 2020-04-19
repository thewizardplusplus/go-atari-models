package atarimodels

import (
	"reflect"
	"testing"
)

func TestNewPreliminaryMove(
	test *testing.T,
) {
	got := NewPreliminaryMove(Black)

	want := Move{
		Color: White,
		Point: NilPoint,
	}
	if !reflect.DeepEqual(got, want) {
		test.Fail()
	}
}
