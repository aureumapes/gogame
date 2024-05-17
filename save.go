package gogame

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/gob"
	"os"
)

// SaveMap saves the Map to the specified path
func (m Map) SaveMap(path string, key []byte) {
	f, _ := os.Create(path)
	enc := gob.NewEncoder(f)
	enc.Encode(m)
	f.Close()

	plaintext, _ := os.ReadFile(path)
	block, _ := aes.NewCipher(key)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	os.WriteFile(path, ciphertext, 666)
}

// LoadMap returns a Map object, from a save file
func LoadMap(path string, key []byte) Map {
	gameBytes, _ := os.ReadFile(path)
	block, _ := aes.NewCipher(key)
	iv := gameBytes[:aes.BlockSize]
	gameBytes = gameBytes[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(gameBytes, gameBytes)

	buf := new(bytes.Buffer)
	buf.Write(gameBytes)
	dec := gob.NewDecoder(buf)
	var game Map
	dec.Decode(&game)
	return game
}
