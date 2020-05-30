package atarimodels

// StoneGroup ...
type StoneGroup map[Point]Color

// Move ...
//
// It doesn't check that the move is correct.
//
func (group StoneGroup) Move(move Move) {
	group[move.Point] = move.Color
}

// Copy ...
func (group StoneGroup) Copy() StoneGroup {
	groupCopy := make(StoneGroup)
	for point, color := range group {
		groupCopy[point] = color
	}

	return groupCopy
}
