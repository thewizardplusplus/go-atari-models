package atarimodels

// Size ...
type Size struct {
	Width  int
	Height int
}

// HasPoint ...
func (size Size) HasPoint(point Point) bool {
	return less(point.Column, size.Width) &&
		less(point.Row, size.Height)
}

func less(value int, limit int) bool {
	return 0 <= value && value < limit
}
