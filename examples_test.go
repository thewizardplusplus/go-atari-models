package atarimodels_test

import (
	"fmt"

	models "github.com/thewizardplusplus/go-atari-models"
)

func ExampleBoard_CheckMove_withLegalSelfcapture() {
	// +-+-+-+-+-+
	// | | | | | |
	// +-+-+-+-+-+
	// | |B|W| | |
	// +-+-+-+-+-+
	// |B|W|X|W| |
	// +-+-+-+-+-+
	// | |B|W| | |
	// +-+-+-+-+-+
	// | | | | | |
	// +-+-+-+-+-+
	board := models.NewBoard(models.Size{Width: 5, Height: 5})
	for _, move := range []models.Move{
		{Color: models.Black, Point: models.Point{Column: 1, Row: 1}},
		{Color: models.White, Point: models.Point{Column: 2, Row: 1}},
		{Color: models.Black, Point: models.Point{Column: 0, Row: 2}},
		{Color: models.White, Point: models.Point{Column: 1, Row: 2}},
		{Color: models.White, Point: models.Point{Column: 3, Row: 2}},
		{Color: models.Black, Point: models.Point{Column: 1, Row: 3}},
		{Color: models.White, Point: models.Point{Column: 2, Row: 3}},
	} {
		board = board.ApplyMove(move)
	}

	move := models.Move{
		Color: models.Black,
		Point: models.Point{Column: 2, Row: 2},
	}
	fmt.Printf("%+v: %v\n", move, board.CheckMove(move))

	// Output: {Color:0 Point:{Column:2 Row:2}}: <nil>
}

func ExampleBoard_CheckMove_withIllegalSelfcapture() {
	// +-+-+-+-+-+
	// | | | | | |
	// +-+-+-+-+-+
	// | | |W| | |
	// +-+-+-+-+-+
	// | |W|X|W| |
	// +-+-+-+-+-+
	// | | |W| | |
	// +-+-+-+-+-+
	// | | | | | |
	// +-+-+-+-+-+
	board := models.NewBoard(models.Size{Width: 5, Height: 5})
	for _, move := range []models.Move{
		{Color: models.White, Point: models.Point{Column: 2, Row: 1}},
		{Color: models.White, Point: models.Point{Column: 1, Row: 2}},
		{Color: models.White, Point: models.Point{Column: 3, Row: 2}},
		{Color: models.White, Point: models.Point{Column: 2, Row: 3}},
	} {
		board = board.ApplyMove(move)
	}

	move := models.Move{
		Color: models.Black,
		Point: models.Point{Column: 2, Row: 2},
	}
	fmt.Printf("%+v: %v\n", move, board.CheckMove(move))

	// Output: {Color:0 Point:{Column:2 Row:2}}: self-capture
}
