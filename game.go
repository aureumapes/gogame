package gogame

// Game is a struct containing Level Data
type Game struct {
	GameMap map[Vector]Cell
}

// NewGame is a function to create a new empty Game
func NewGame() Game {
	return Game{make(map[Vector]Cell)}
}

// MoveCell allows to move a Cell to another position, replacing its old Content with the Content, of the Cell, it moved to
func (g Game) MoveCell(cell Cell, moveTo Vector) (Game, Cell) {
	oldCell := cell
	newCell, exists := g.GameMap[moveTo]
	if !exists {
		return g, cell
	}
	cell.Pos = newCell.Pos
	g.AddCell(oldCell.Pos, newCell.Content)
	g.AddCell(newCell.Pos, oldCell.Content)
	return g, cell
}
