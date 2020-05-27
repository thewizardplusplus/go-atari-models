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
)

// StoneStorage ...
type StoneStorage interface {
	Size() Size
	HasCapture(
		options ...HasCaptureOption,
	) (Color, bool)
	ApplyMove(move Move) StoneStorage
	CheckMove(move Move) error
}

// Board ...
type Board struct {
	size   Size
	stones StoneGroup
}

// NewBoard ...
func NewBoard(size Size) StoneStorage {
	stones := make(StoneGroup)
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
) (
	neighbors StoneGroup,
	hasStoneLiberties bool,
) {
	neighbors = make(StoneGroup)
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
			hasStoneLiberties = true
		}
	}

	return neighbors, hasStoneLiberties
}

// HasChainLiberties ...
//
// There should be a stone at the point.
//
// The chain shouldn't be nil.
//
// After finishing the function
// the chain will be filled
// (partially, if the result is true).
func (board Board) HasChainLiberties(
	point Point,
	chain PointGroup,
) bool {
	if _, ok := chain[point]; ok {
		return false
	}

	baseColor := board.stones[point]
	chain[point] = struct{}{}

	neighbors, hasStoneLiberties :=
		board.StoneNeighbors(point)
	if hasStoneLiberties {
		return true
	}

	for neighbor := range neighbors {
		color := board.stones[neighbor]
		if color != baseColor {
			continue
		}

		hasLiberties := board.HasChainLiberties(
			neighbor,
			chain,
		)
		if hasLiberties {
			return true
		}
	}

	return false
}

// HasCaptureConfiguration ...
type HasCaptureConfiguration struct {
	FilterByColor  bool
	FilterByOrigin bool
	SampleColor    Color
	Origin         Point
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
		configuration.FilterByColor = true
		configuration.SampleColor = color
	}
}

// WithOrigin ...
//
// There should be a stone
// at the origin point.
//
// If the origin is NilPoint,
// then it'll be ignored.
func WithOrigin(
	origin Point,
) HasCaptureOption {
	return func(
		configuration *HasCaptureConfiguration,
	) {
		configuration.FilterByOrigin = true
		configuration.Origin = origin
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

	var stones StoneGroup
	if configuration.FilterByOrigin &&
		!configuration.Origin.IsNil() {
		stones, _ = board.
			StoneNeighbors(configuration.Origin)

		// copy the origin stone
		stones[configuration.Origin] =
			board.stones[configuration.Origin]
	} else {
		stones = board.stones
	}

	for point, color := range stones {
		if configuration.FilterByColor &&
			color != configuration.SampleColor {
			continue
		}

		hasLiberties := board.HasChainLiberties(
			point,
			make(PointGroup),
		)
		if !hasLiberties {
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
) StoneStorage {
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
			WithOrigin(move.Point),
		)
	_, opponentCapture :=
		nextBoard.HasCapture(
			WithColor(nextColor),
			WithOrigin(move.Point),
		)
	if selfcapture && !opponentCapture {
		return ErrSelfcapture
	}

	return nil
}
