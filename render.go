package gogame

// Render prints the Game‎s grid as an
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
