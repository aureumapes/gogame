package gogame

// Render prints the Game‎s grid as an
func (m Map) Render() {
	y := 0
	x := 0
	for i := 0; i <= len(m); i++ {
		curr, exists := m[Vector{x, y}]
		if !exists {
			y++
			x = 0
			print("\n")
			curr, exists = m[Vector{x, y}]
			if !exists {
				break
			}
			print(string(curr.Content))
			x++
		} else {
			curr = m[Vector{x, y}]
			print(string(curr.Content))
			x++
		}
	}
}
