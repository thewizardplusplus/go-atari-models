package atarimodels

type stoneFilter func(Point, Color) bool

type stoneGroup map[Point]Color

// It doesn't check that the move is correct.
func (group stoneGroup) Move(move Move) {
	group[move.Point] = move.Color
}

func (group stoneGroup) Copy() stoneGroup {
	groupCopy := make(stoneGroup)
	for point, color := range group {
		groupCopy[point] = color
	}

	return groupCopy
}
