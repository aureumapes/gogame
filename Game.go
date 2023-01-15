package gogame

type Game struct {
	GameMap map[Vector]Cell
}

func NewGame() Game {
	return Game{make(map[Vector]Cell)}
}

func (g Game) AddCell(pos Vector, content rune) {
	g.GameMap[pos] = Cell{pos, content}
}

func (g Game) AddRow(row string, y int) {
	for x, cell := range []rune(row) {
		cell1 := Cell{Vector{x, y}, cell}
		g.GameMap[cell1.Pos] = cell1
	}
}

func (g Game) Render() {
	y := 0
	x := 0
	for i := 0; i <= len(g.GameMap); i++ {
		curr, exists := g.GameMap[Vector{x, y}]
		if !exists {
			y++
			x = 0
			print("\n")
			curr, exists = g.GameMap[Vector{x, y}]
			if !exists {
				break
			}
			print(string(curr.Content))
			x++
		} else {
			curr = g.GameMap[Vector{x, y}]
			print(string(curr.Content))
			x++
		}
	}
}

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
