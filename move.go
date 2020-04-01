package atarimodels

// Move ...
type Move struct {
	Color Color
	Point Point
}

// NewPreliminaryMove ...
//
// It creates the move
// from only the negated passed color.
func NewPreliminaryMove(color Color) Move {
	return Move{
		Color: color.Negative(),
	}
}
