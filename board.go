package atarimodels

import (
	"errors"
)

// ...
var (
	ErrOutOfSize = errors.New(
		"out of size",
	)
	ErrOccupiedPoint = errors.New(
		"occupied point",
	)
)

// Board ...
type Board struct {
	size   Size
	stones stoneGroup
}

// NewBoard ...
func NewBoard(size Size) Board {
	stones := make(stoneGroup)
	return Board{size, stones}
}

// Size ...
func (board Board) Size() Size {
	return board.size
}

// Stone ...
func (board Board) Stone(
	point Point,
) (color Color, ok bool) {
	color, ok = board.stones[point]
	return color, ok
}

// StoneNeighbors ...
func (board Board) StoneNeighbors(
	point Point,
) []Point {
	var neighbors []Point
	for _, shift := range []Point{
		Point{0, -1},
		Point{-1, 0},
		Point{1, 0},
		Point{0, 1},
	} {
		neighbor := Point{
			Column: point.Column + shift.Column,
			Row:    point.Row + shift.Row,
		}

		if _, ok := board.stones[neighbor]; ok {
			neighbors = append(
				neighbors,
				neighbor,
			)
		}
	}

	return neighbors
}

// ApplyMove ...
//
// It doesn't check that the move
// is correct.
func (board Board) ApplyMove(
	move Move,
) Board {
	stones := board.stones.Copy()
	stones.Move(move)

	return Board{board.size, stones}
}

// CheckMove ...
func (board Board) CheckMove(
	move Move,
) error {
	if !board.size.HasPoint(move.Point) {
		return ErrOutOfSize
	}

	if _, ok := board.stones[move.Point]; ok {
		return ErrOccupiedPoint
	}

	return nil
}

// MovesForColor ...
func (board Board) MovesForColor(
	color Color,
) []Move {
	var moves []Move
	points := board.size.Points()
	for _, point := range points {
		move := Move{color, point}
		err := board.CheckMove(move)
		if err != nil {
			continue
		}

		moves = append(moves, move)
	}

	return moves
}