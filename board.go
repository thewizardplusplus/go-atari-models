package atarimodels

import (
	"errors"
)

const (
	maximalNeighborCount = 4
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
)

// ...
var (
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
	chain PointGroup,
) int {
	if _, ok := chain[point]; ok {
		return 0
	}

	baseColor := board.stones[point]
	chain[point] = struct{}{}

	var liberties int
	neighbors, _ := board.neighbors(point)
	liberties += maximalNeighborCount -
		len(neighbors)

	for neighbor := range neighbors {
		color := board.stones[neighbor]
		if color != baseColor {
			continue
		}

		liberties +=
			board.StoneLiberties(neighbor, chain)
	}

	return liberties
}

// HasCaptureConfiguration ...
type HasCaptureConfiguration struct {
	filterByColor  bool
	filterByOrigin bool
	sampleColor    Color
	origin         Point
}

// HasCaptureOption ...
type HasCaptureOption func(
	configuration *HasCaptureConfiguration,
)

// WithColor ...
func WithColor(
	color Color,
) HasCaptureOption {
	return func(
		configuration *HasCaptureConfiguration,
	) {
		configuration.filterByColor = true
		configuration.sampleColor = color
	}
}

// WithOrigin ...
func WithOrigin(
	origin Point,
) HasCaptureOption {
	return func(
		configuration *HasCaptureConfiguration,
	) {
		configuration.filterByOrigin = true
		configuration.origin = origin
	}
}

// HasCapture ...
func (board Board) HasCapture(
	options ...HasCaptureOption,
) (Color, bool) {
	var configuration HasCaptureConfiguration
	for _, option := range options {
		option(&configuration)
	}

	var stones stoneGroup
	/*if configuration.filterByOrigin &&
	    !configuration.origin.IsNil() {
	    neighbors := board.
	      StoneNeighbors(configuration.origin)
	    stones = board.stones.
	      CopyByPoints(neighbors)
	  } else {
	    stones = board.stones
	  }*/
	stones = board.stones

	for point, color := range stones {
		if configuration.filterByColor &&
			color != configuration.sampleColor {
			continue
		}

		liberties := board.StoneLiberties(
			point,
			make(PointGroup),
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
		nextBoard.HasCapture(
			WithColor(move.Color),
		)
	_, opponentCapture :=
		nextBoard.HasCapture(
			WithColor(nextColor),
		)
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
//
// Returned error can be
// ErrAlreadyLoss or ErrAlreadyWin only.
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

func (board Board) neighbors(
	point Point,
) (
	neighbors stoneGroup,
	hasLiberties bool,
) {
	neighbors = make(stoneGroup)
	for _, shift := range []Point{
		Point{0, -1},
		Point{-1, 0},
		Point{1, 0},
		Point{0, 1},
	} {
		neighbor := point.Translate(shift)
		if !board.size.HasPoint(neighbor) {
			continue
		}

		color, ok := board.stones[neighbor]
		if ok {
			neighbors[neighbor] = color
		} else {
			hasLiberties = true
		}
	}

	return neighbors, hasLiberties
}
