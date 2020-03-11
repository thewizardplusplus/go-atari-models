package atarimodels

import (
	"reflect"
	"testing"
)

func TestNewMove(test *testing.T) {
	got := NewMove(Black)

	want := Move{Color: White}
	if !reflect.DeepEqual(got, want) {
		test.Fail()
	}
}
