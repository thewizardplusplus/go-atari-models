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
	ErrSelfcapture = errors.New(
		"self-capture",
	)
	ErrAlreadyLoss = errors.New(
		"already loss",
	)
	ErrAlreadyWin = errors.New("already win")
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
) (empty []Point, occupied []Point) {
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
		if !board.size.HasPoint(neighbor) {
			continue
		}

		if _, ok := board.stones[neighbor]; ok {
			occupied = append(occupied, neighbor)
		} else {
			empty = append(empty, neighbor)
		}
	}

	return empty, occupied
}

// StoneLiberties ...
//
// There should be a stone at the point.
//
// The chain shouldn't be nil.
//
// After finishing the function
// the chain will be completed.
func (board Board) StoneLiberties(
	point Point,
	chain map[Point]struct{},
) int {
	if _, ok := chain[point]; ok {
		return 0
	}

	baseColor := board.stones[point]
	chain[point] = struct{}{}

	var liberties int
	empty, occupied := board.
		StoneNeighbors(point)
	liberties += len(empty)

	for _, neighbor := range occupied {
		color := board.stones[neighbor]
		if color != baseColor {
			continue
		}

		liberties +=
			board.StoneLiberties(neighbor, chain)
	}

	return liberties
}

// HasCapture ...
func (board Board) HasCapture(
	color ...Color,
) (Color, bool) {
	var filterByColor bool
	var sampleColor Color
	if len(color) != 0 {
		filterByColor = true
		sampleColor = color[0]
	}

	for point, color := range board.stones {
		if filterByColor &&
			color != sampleColor {
			continue
		}

		liberties := board.StoneLiberties(
			point,
			make(map[Point]struct{}),
		)
		if liberties == 0 {
			return color, true
		}
	}

	return 0, false
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

	nextBoard := board.ApplyMove(move)
	nextColor := move.Color.Negative()
	_, selfcapture :=
		nextBoard.HasCapture(move.Color)
	_, opponentCapture :=
		nextBoard.HasCapture(nextColor)
	if selfcapture && !opponentCapture {
		return ErrSelfcapture
	}

	return nil
}

// PseudolegalMoves ...
func (board Board) PseudolegalMoves(
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

// LegalMoves ...
func (board Board) LegalMoves(
	color Color,
) ([]Move, error) {
	moves := board.PseudolegalMoves(color)
	if len(moves) == 0 {
		// game result in this case
		// depends on used game rules
		return nil, ErrAlreadyLoss
	}

	captureColor, ok := board.HasCapture()
	if ok {
		var err error
		if captureColor == color {
			err = ErrAlreadyLoss
		} else {
			err = ErrAlreadyWin
		}

		return nil, err
	}

	return moves, nil
}
