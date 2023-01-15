package gogame

type Cell struct {
	Pos     Vector
	Content rune
}

func NewCell(v Vector, content rune) Cell {
	return Cell{v, content}
}

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
