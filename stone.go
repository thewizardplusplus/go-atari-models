package atarimodels

// Color ...
type Color int

// ...
const (
	Black Color = iota
	White
)

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

type stoneGroup map[Point]Color

// It doesn't check that the move is correct.
func (group stoneGroup) Move(move Move) {
	group[move.Point] = move.Color
}

func (group stoneGroup) Copy() stoneGroup {
	groupCopy := make(stoneGroup)
	for point, color := range group {
		move := Move{color, point}
		groupCopy.Move(move)
	}

	return groupCopy
}
