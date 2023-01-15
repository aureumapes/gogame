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
	game := gogame.NewGame() // Creating a new Game
	player := gogame.NewCell(gogame.Vector{X: 1, Y: 1}, '🔴') // Createing a Cell for the player
	game.AddRow("⬛⬛⬛⬛⬛", 0) // Filling the map
	game.AddRow("⬛⬛⬛⬛⬛", 1)
	game.AddRow("⬛⬛⬛⬛⬛", 2)
	game.AddRow("⬛⬛⬛⬛⬛", 3)
	game.AddRow("⬛⬛⬛⬛⬛", 4)
	game.AddCell(player.Pos, player.Content) // Place the player
	
	// Be carefull with Infinity loops
	for true {
		game.Render() // Print the current map to the console
		
		// Following is basic Movement
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
		print("\033[H\033[2J") // "Clears" the console, This only works if your terminal supports Ascii codes
	}
}

```
