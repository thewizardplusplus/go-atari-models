package sgf_test

import (
	"fmt"

	models "github.com/thewizardplusplus/go-atari-models"
	"github.com/thewizardplusplus/go-atari-models/encoding/sgf"
)

func ExampleDecodeColor() {
	color, _ := sgf.DecodeColor('W')
	fmt.Printf("%d\n", color)

	// Output: 1
}

func ExampleEncodeColor() {
	color := sgf.EncodeColor(models.White)
	fmt.Printf("%c\n", color)

	// Output: W
}

func ExampleDecodeAxis() {
	axis, _ := sgf.DecodeAxis('E')
	fmt.Printf("%d\n", axis)

	// Output: 30
}

func ExampleEncodeAxis() {
	axis := sgf.EncodeAxis(30)
	fmt.Printf("%c\n", axis)

	// Output: E
}

func ExampleDecodePoint() {
	point, _ := sgf.DecodePoint("eE")
	fmt.Printf("%+v\n", point)

	// Output: {Column:4 Row:30}
}

func ExampleEncodePoint() {
	point := sgf.EncodePoint(models.Point{Column: 4, Row: 30})
	fmt.Printf("%s\n", point)

	// Output: eE
}

func ExampleDecodeStoneStorage() {
	storage, _ := sgf.DecodeStoneStorage(
		"(;FF[4]SZ[7:9]GN[test];B[aa](;W[gi]N[test]))",
		models.NewBoard,
	)

	var moves []models.Move
	for _, point := range storage.Size().Points() {
		if color, ok := storage.Stone(point); ok {
			moves = append(moves, models.Move{Color: color, Point: point})
		}
	}

	fmt.Printf("%+v\n", storage.Size())
	fmt.Printf("%+v\n", moves)

	// Output:
	// {Width:7 Height:9}
	// [{Color:0 Point:{Column:0 Row:0}} {Color:1 Point:{Column:6 Row:8}}]
}
