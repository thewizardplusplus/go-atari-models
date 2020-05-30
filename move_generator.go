package atarimodels

import (
	"errors"
)

// ...
var (
	ErrAlreadyLoss = errors.New("already loss")
	ErrAlreadyWin  = errors.New("already win")
)

// Generator ...
type Generator interface {
	LegalMoves(storage StoneStorage, previousMove Move) ([]Move, error)
}

// MoveGenerator ...
type MoveGenerator struct{}

// PseudolegalMoves ...
func (generator MoveGenerator) PseudolegalMoves(
	storage StoneStorage,
	color Color,
) []Move {
	var moves []Move
	for _, point := range storage.Size().Points() {
		move := Move{color, point}
		if err := storage.CheckMove(move); err != nil {
			continue
		}

		moves = append(moves, move)
	}

	return moves
}

// LegalMoves ...
//
// Returned error can be ErrAlreadyLoss or ErrAlreadyWin only.
//
func (generator MoveGenerator) LegalMoves(
	storage StoneStorage,
	previousMove Move,
) ([]Move, error) {
	nextColor := previousMove.Color.Negative()
	if captureColor, ok := storage.HasCapture(WithOrigin(previousMove.Point)); ok {
		var err error
		if captureColor == nextColor {
			err = ErrAlreadyLoss
		} else {
			err = ErrAlreadyWin
		}

		return nil, err
	}

	moves := generator.PseudolegalMoves(storage, nextColor)
	if len(moves) == 0 {
		// game result in this case depends on used game rules
		return nil, ErrAlreadyLoss
	}

	return moves, nil
}
