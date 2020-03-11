package atarimodels

// Point ...
type Point struct {
	Column int
	Row    int
}

// Move ...
type Move struct {
	Color Color
	Point Point
}

// NewMove ...
func NewMove(color Color) Move {
	return Move{
		Color: color.Negative(),
	}
}
