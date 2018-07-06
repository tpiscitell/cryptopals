package main

import (
	"crypto/aes"
	"fmt"

	"github.com/tpiscitell/cryptopals/utils"
)

func main() {
	key := []byte("YELLOW SUBMARINE")
	ciphertext := utils.ReadBase64File("7.txt")

	blockSize := len(key)
	blockCount := len(ciphertext) / blockSize

	buf := make([]byte, blockSize)
	block, _ := aes.NewCipher(key)

	var plaintext []byte

	for i := 0; i < blockCount; i++ {
		start := i * blockSize
		end := start + blockSize

		block.Decrypt(buf, ciphertext[start:end])
		plaintext = append(plaintext, buf...)
	}

	fmt.Println(string(plaintext))
}
