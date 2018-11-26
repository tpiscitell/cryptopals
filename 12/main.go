package main

import (
	"github.com/tpiscitell/cryptopals/utils"
)

var secretString = utils.ReadBase64File("12.txt")
var key = utils.RandBytes(16)

func Encrypt(in []byte) []byte {
	var paddedIn []byte

	paddedIn = append(paddedIn, in...)
	paddedIn = append(paddedIn, secretString...)

	return utils.AesEcbEncrypt(paddedIn, key)
}

func getInput(size int, fill byte) []byte {
	input := make([]byte, size)
	for i := range input {
		input[i] = fill
	}

	return input
}
func main() {
	blockSize := discoverBlockSize()
	println("Block size is", blockSize)

	out := Encrypt(getInput(blockSize*6, 'A'))

	repeats := utils.CountRepeats(out, blockSize)
	if repeats > 0 {
		println("ECB")
		oneByteShort := getInput(blockSize-1, 'A')
		obsOut := Encrypt(oneByteShort)

		for i := byte(0); i < 255; i++ {
			fullBlock := append(oneByteShort, i)
			fullOut := Encrypt(fullBlock)
			if sliceEq(fullOut[:blockSize], obsOut[:blockSize]) {
				println("First byte is ", i)
			}
		}
	} else {
		println("CBC detected :(")
	}
}

func sliceEq(a, b []byte) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func discoverBlockSize() int {
	var input []byte
	var lastLen int
	for i := 0; i <= 64; i++ {
		input = append(input, 'A')
		out := Encrypt(input)
		outLen := len(out)

		if lastLen == 0 {
			lastLen = outLen
		} else {
			if outLen > lastLen {
				return outLen - lastLen
			}
		}
	}
	return 0
}
