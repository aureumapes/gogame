package main

import (
	_ "embed"
	"encoding/hex"
	"github.com/aureumapes/gogame"
	"os"
)

type Level struct {
	gogame.Map
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
	gameMap1 := Level{
		gogame.NewMap(),
		"startlvl",
		01,
		gogame.Vector{1, 1},
	}
	gameMap2 := Level{
		gogame.NewMap(),
		"secondlvl",
		02,
		gogame.Vector{1, 1},
	}
	game := gogame.NewGame("test", gameMap1.Name)
	game.Maps[gameMap1.Name] = gameMap1.Map
	game.Maps[gameMap2.Name] = gameMap2.Map

	player := gogame.NewCell(gogame.Vector{X: 1, Y: 1}, '🔴')

	gameMap1.AddRow("⬛⬛⬛⬛⬛", 0)
	gameMap1.AddRow("⬛⬛⬛⬛⬛", 1)
	gameMap1.AddRow("⬛⬛⬛⬛⬛", 2)
	gameMap1.AddRow("⬛⬛⬛⬛⬛", 3)
	gameMap1.AddRow("⬛⬛⬛⬛⬛", 4)
	gameMap1.AddCell(player.Pos, player.Content)

	gameMap2.AddRow("⬜⬜⬜⬜⬜", 0)
	gameMap2.AddRow("⬜⬜⬜⬜⬜", 1)
	gameMap2.AddRow("⬜⬜⬜⬜⬜", 2)
	gameMap2.AddRow("⬜⬜⬜⬜⬜", 3)
	gameMap2.AddRow("⬜⬜⬜⬜⬜", 4)
	gameMap2.AddCell(player.Pos, '🟢')

	currGameMap := gameMap1

	for true {
		currGameMap.Render()
		input := gogame.ReadInput("Input > ")
		switch input {
		case 'w':
			currGameMap.Map, player = currGameMap.MoveCell(player, player.GetUpper())
			currGameMap.PlayerPos = player.Pos
			break
		case 'a':
			currGameMap.Map, player = currGameMap.MoveCell(player, player.GetLeft())
			currGameMap.PlayerPos = player.Pos
			break
		case 's':
			currGameMap.Map, player = currGameMap.MoveCell(player, player.GetDowner())
			currGameMap.PlayerPos = player.Pos
			break
		case 'd':
			currGameMap.Map, player = currGameMap.MoveCell(player, player.GetRight())
			currGameMap.PlayerPos = player.Pos
			break
		case 'k':
			game.Save(key)
			break
		case 'l':
			currGameMap.Map = gogame.LoadMap("Saves/"+game.Name+"_"+currGameMap.Name, key)
			player.Pos = currGameMap.PlayerPos
			break
		case 'o':
			game.Save(key)
			if game.CurrentMap == gameMap1.Name {
				gameMap1.PlayerPos = player.Pos
				game.CurrentMap = gameMap2.Name
				currGameMap = gameMap2
				player.Content = '🟢'
				player.Pos = gameMap2.PlayerPos
			} else {
				gameMap2.PlayerPos = player.Pos
				game.CurrentMap = gameMap1.Name
				currGameMap = gameMap1
				player.Content = '🔴'
				player.Pos = gameMap1.PlayerPos
			}
			break
		case 'q':
			for gameMapName, _ := range game.Maps {
				os.Remove("Saves/" + game.Name + "_" + gameMapName)
			}
			os.Remove("Saves/master")
			os.Exit(130)
		default:
			break
		}
		print("\033[H\033[2J")
	}
}
