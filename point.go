package atarimodels

// Point ...
type Point struct {
	Column int
	Row    int
}

// Translate ...
func (point Point) Translate(
	translation Point,
) Point {
	return Point{
		Column: point.Column +
			translation.Column,
		Row: point.Row + translation.Row,
	}
}
