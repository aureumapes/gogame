gogame
======
A simple TUI-Game Library for go

Installation
------------
1. Run `go get github.com/AureumApes/gogame`
2. Start using with `ìmport github.com/AureumApes/gogame`

Example
-------
This is a simple game:
```go
package main

import "github.com/AureumApes/gogame"

func main() {
	game := gogame.NewGame()
	player := gogame.NewCell(gogame.Vector{X: 1, Y: 1}, '🔴')
	game.AddRow("⬛⬛⬛⬛⬛", 0)
	game.AddRow("⬛⬛⬛⬛⬛", 1)
	game.AddRow("⬛⬛⬛⬛⬛", 2)
	game.AddRow("⬛⬛⬛⬛⬛", 3)
	game.AddRow("⬛⬛⬛⬛⬛", 4)
	game.AddCell(player.Pos, player.Content)

	for true {
		game.Render()
		input := gogame.ReadInput("Input > ")
		switch input {
		case 'w':
			game, player = game.MoveCell(player, player.GetUpper())
			break
		case 'a':
			game, player = game.MoveCell(player, player.GetLeft())
			break
		case 's':
			game, player = game.MoveCell(player, player.GetDowner())
			break
		case 'd':
			game, player = game.MoveCell(player, player.GetRight())
			break
		default:
			break
		}
		print("\033[H\033[2J") // This only works if your terminal supports Ascii codes
	}
}

```
