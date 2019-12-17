package atarimodels

// Size ...
type Size struct {
	Width  int
	Height int
}

// Board ...
type Board struct {
	size   Size
	stones stoneGroup
}

// NewBoard ...
func NewBoard(size Size) Board {
	stones := make(stoneGroup)
	return Board{size, stones}
}

// Size ...
func (board Board) Size() Size {
	return board.size
}

// Piece ...
func (board Board) Stone(
	point Point,
) (color Color, ok bool) {
	color, ok = board.stones[point]
	return color, ok
}

// ApplyMove ...
//
// It doesn't check that the move
// is correct.
func (board Board) ApplyMove(
	move Move,
) Board {
	stones := board.stones.Copy()
	stones.Move(move)

	return Board{board.size, stones}
}
