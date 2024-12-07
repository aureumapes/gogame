package gogame

// Cell is a structure, which stores the Position and Content of a Cell inside the Grid
type Cell struct {
	Pos     Vector
	Content rune
}

// NewCell generates and returns a new Cell
func NewCell(v Vector, content rune) Cell {
	return Cell{v, content}
}

// GetUpper GetDowner GetLeft and GetRight functions will return the position of neighbor Cell
func (cell Cell) GetUpper() Vector {
	return Vector{X: cell.Pos.X, Y: cell.Pos.Y - 1}
}

func (cell Cell) GetDowner() Vector {
	return Vector{X: cell.Pos.X, Y: cell.Pos.Y + 1}
}

func (cell Cell) GetLeft() Vector {
	return Vector{X: cell.Pos.X - 1, Y: cell.Pos.Y}
}

func (cell Cell) GetRight() Vector {
	return Vector{X: cell.Pos.X + 1, Y: cell.Pos.Y}
}
