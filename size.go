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

// Points ...
func (size Size) Points() []Point {
	var points []Point
	width, height := size.Width, size.Height
	row := 0
	for ; row < height; row++ {
		column := 0
		for ; column < width; column++ {
			point := Point{column, row}
			points = append(points, point)
		}
	}

	return points
}

func less(value int, limit int) bool {
	return 0 <= value && value < limit
}
