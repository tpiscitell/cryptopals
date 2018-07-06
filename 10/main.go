package main

import (
	"fmt"

	"github.com/tpiscitell/cryptopals/utils"
)

func main() {
	key := []byte("YELLOW SUBMARINE")
	iv := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	ciphertext := utils.ReadBase64File("10.txt")

	plaintext := utils.AesCbcDecrypt(ciphertext, key, iv)

	fmt.Println(string(plaintext))
}
