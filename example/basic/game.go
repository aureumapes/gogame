package main

import (
	_ "embed"
	"encoding/hex"
	"github.com/aureumapes/gogame"
	"os"
)

type Level struct {
	gogame.Game
	Name      string
	Number    int64
	PlayerPos gogame.Vector
}

var keyStr string = "816760757b8f92036e7d739e7065eb496015fe9808616e90cd5c346aa613bfaa"

func main() {
	key, _ := hex.DecodeString(keyStr)
	println(keyStr)
	f, _ := os.Create("key")
	_, err := f.WriteString(keyStr)
	if err != nil {
		panic(err)
		return
	}
	f.Close()
	game := Level{
		gogame.NewGame(),
		"",
		01,
		gogame.Vector{1, 1},
	}
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
			game.Game, player = game.MoveCell(player, player.GetUpper())
			break
		case 'a':
			game.Game, player = game.MoveCell(player, player.GetLeft())
			break
		case 's':
			game.Game, player = game.MoveCell(player, player.GetDowner())
			break
		case 'd':
			game.Game, player = game.MoveCell(player, player.GetRight())
			break
		case 'k':
			game.Game.SaveGame("game.save", key)
			break
		case 'l':
			game.Game = gogame.LoadGame("game.save", key)
			player.Pos = game.FindCell(player.Content)
			break
		default:
			break
		}
		print("\033[H\033[2J")
	}
}
