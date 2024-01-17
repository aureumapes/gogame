package gogame

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/gob"
	"os"
)

// SaveGame saves the Game tp the specified path
func (g Game) SaveGame(path string, key []byte) {
	f, _ := os.Create(path)
	enc := gob.NewEncoder(f)
	enc.Encode(g)
	f.Close()

	plaintext, _ := os.ReadFile(path)
	block, _ := aes.NewCipher(key)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	os.WriteFile(path, ciphertext, 666)
}

// LoadGame returns a Game object, from a save file
func LoadGame(path string, key []byte) Game {
	gameBytes, _ := os.ReadFile(path)
	block, _ := aes.NewCipher(key)
	iv := gameBytes[:aes.BlockSize]
	gameBytes = gameBytes[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(gameBytes, gameBytes)

	buf := new(bytes.Buffer)
	buf.Write(gameBytes)
	dec := gob.NewDecoder(buf)
	var game Game
	dec.Decode(&game)
	return game
}
