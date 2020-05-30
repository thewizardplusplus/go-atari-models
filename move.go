package atarimodels

// Move ...
type Move struct {
	Color Color
	Point Point
}

// NewPreliminaryMove ...
//
// It creates the move from the negated passed color and the nil point.
//
func NewPreliminaryMove(color Color) Move {
	return Move{
		Color: color.Negative(),
		Point: NilPoint,
	}
}
