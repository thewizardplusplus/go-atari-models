package atarimodels

// Point ...
type Point struct {
	Column int
	Row    int
}

// PointGroup ...
type PointGroup map[Point]struct{}

// ...
//
// nolint: gochecknoglobals
//
var (
	NilPoint = Point{-1, -1}
)

// IsNil ...
func (point Point) IsNil() bool {
	return point == NilPoint
}

// Translate ...
func (point Point) Translate(translation Point) Point {
	return Point{
		Column: point.Column + translation.Column,
		Row:    point.Row + translation.Row,
	}
}
