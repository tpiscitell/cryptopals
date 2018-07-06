package utils

import "crypto/aes"

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
