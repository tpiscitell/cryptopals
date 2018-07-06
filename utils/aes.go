package utils

import (
	"crypto/aes"
	"encoding/hex"
)

func AesEcbEncrypt(plaintext, key []byte) []byte {
	blockSize := len(key)
	blockCount := len(plaintext) / blockSize

	buf := make([]byte, blockSize)
	block, _ := aes.NewCipher(key)

	var ciphertext []byte

	for i := 0; i < blockCount; i++ {
		start := i * blockSize
		end := start + blockSize

		block.Encrypt(buf, plaintext[start:end])
		ciphertext = append(ciphertext, buf...)
	}

	return ciphertext
}

func AesEcbDecrypt(ciphertext, key []byte) []byte {
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

	return plaintext
}

func AesCbcEncrypt(plaintext, key, iv []byte) []byte {
	blockSize := len(key)
	blockCount := len(plaintext) / blockSize

	buf := make([]byte, blockSize)
	block, _ := aes.NewCipher(key)

	var ciphertext []byte
	prevBlock := iv

	for i := 0; i < blockCount; i++ {
		start := i * blockSize
		end := start + blockSize

		block.Encrypt(buf, XOR(plaintext[start:end], prevBlock))
		ciphertext = append(ciphertext, buf...)
		prevBlock = buf
	}

	return ciphertext
}

func AesCbcDecrypt(ciphertext, key, iv []byte) []byte {
	blockSize := len(key)
	blockCount := len(ciphertext) / blockSize

	buf := make([]byte, blockSize)
	block, _ := aes.NewCipher(key)

	var plaintext []byte
	prevBlock := iv

	for i := 0; i < blockCount; i++ {
		start := i * blockSize
		end := start + blockSize

		block.Decrypt(buf, ciphertext[start:end])
		xorBuf := XOR(buf, prevBlock)
		plaintext = append(plaintext, xorBuf...)
		prevBlock = ciphertext[start:end]
	}

	return plaintext
}

func CountRepeats(in []byte, blockSize int) int {
	counter := make(map[string]int)
	blockCount := len(in) / blockSize

	for i := 0; i < blockCount; i++ {
		start := i * blockSize
		end := start + blockSize

		slice := hex.EncodeToString(in[start:end])
		counter[slice]++
	}

	var repeats int
	for _, v := range counter {
		if v > 1 {
			repeats += v
		}
	}

	return repeats
}
