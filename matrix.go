package gogame

// AddCell adds a single Cell
func (m Map) AddCell(pos Vector, content rune) {
	m[pos] = Cell{pos, content}
}

// AddRow adds an entire Row to the Map
func (m Map) AddRow(row string, y int) {
	for x, cell := range []rune(row) {
		cell1 := Cell{Vector{x, y}, cell}
		m[cell1.Pos] = cell1
	}
}

// GetCell let you read the Content of a Cell at a specific location on the grid
func (m Map) GetCell(y int, x int) string {
	return string(m[Vector{x, y}].Content)
}

// FindCell allows you to find out a Cells Vector, using its Content, it will always return the first result
func (m Map) FindCell(content rune) Vector {
	for _, cell := range m {
		if cell.Content == content {
			return cell.Pos
		}
	}
	return Vector{0, 0}
}
