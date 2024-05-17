package gogame

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/gob"
	"os"
)

// Game is a collection of maps and the current Map
type Game struct {
	Maps       map[string]Map
	CurrentMap string
	Name       string
}

// NewGame creates a new Game-Object
func NewGame(name string, startMap string) Game {
	return Game{make(map[string]Map), startMap, name}
}

// Save saves tge game and its Maps to files
func (g Game) Save(key []byte) {
	for name, gameMap := range g.Maps {
		gameMap.SaveMap("Saves/"+g.Name+"_"+name, key)
	}
	f, _ := os.Create("Saves/master")
	enc := gob.NewEncoder(f)
	enc.Encode(g.CurrentMap + "splitter" + g.Name)
	f.Close()

	plaintext, _ := os.ReadFile("Saves/master")
	block, _ := aes.NewCipher(key)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	os.WriteFile("Saves/master", ciphertext, 666)
}
