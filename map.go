package gogame

// Map contains Level Data
type Map map[Vector]Cell

// NewMap is a function to create a new empty Map
func NewMap() Map {
	return make(map[Vector]Cell)
}

// MoveCell allows to move a Cell to another position, replacing its old Content with the Content, of the Cell, it moved to
func (m Map) MoveCell(cell Cell, moveTo Vector) (Map, Cell) {
	oldCell := cell
	newCell, exists := m[moveTo]
	if !exists {
		return m, cell
	}
	cell.Pos = newCell.Pos
	m.AddCell(oldCell.Pos, newCell.Content)
	m.AddCell(newCell.Pos, oldCell.Content)
	return m, cell
}
