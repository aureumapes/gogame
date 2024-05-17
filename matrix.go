package gogame

// AddCell adds a single Cell
func (g Game) AddCell(pos Vector, content rune) {
	g.GameMap[pos] = Cell{pos, content}
}

// AddRow adds an entire Row to the Game
func (g Game) AddRow(row string, y int) {
	for x, cell := range []rune(row) {
		cell1 := Cell{Vector{x, y}, cell}
		g.GameMap[cell1.Pos] = cell1
	}
}

// GetCell let you read the Content of a Cell at a specific location on the grid
func (g Game) GetCell(y int, x int) string {
	return string(g.GameMap[Vector{x, y}].Content)
}

// FindCell allows you to find out a Cells Vector, using its Content, it will always return the first result
func (g Game) FindCell(content rune) Vector {
	for _, cell := range g.GameMap {
		if cell.Content == content {
			return cell.Pos
		}
	}
	return Vector{0, 0}
}
