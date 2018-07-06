package main

import (
	"fmt"

	"github.com/tpiscitell/cryptopals/utils"
)

func HammingWeight(a byte) int {
	m1 := byte(0x55) // 01010101
	m2 := byte(0x33) // 00110011
	m4 := byte(0x0f) // 00001111

	a = (a & m1) + ((a >> 1) & m1)
	a = (a & m2) + ((a >> 2) & m2)
	a = (a & m4) + ((a >> 4) & m4)

	return int(a)
}

func HammingDist(a, b []byte) int {
	xor := utils.XOR(a, b)
	sum := 0

	for _, i := range xor {
		sum += HammingWeight(i)
	}

	return sum
}

func HammingDistStr(a, b string) int {
	return HammingDist([]byte(a), []byte(b))
}

func FindKeySize(in []byte) int {
	var keysize int
	minDist := 10000.0 //seed with a really high number
	maxSize := 40
	minSize := 2
	blocks := len(in) / maxSize

	for keyLen := minSize; keyLen < maxSize; keyLen++ {
		dist := 0.0
		for i := 0; i < blocks; i++ {
			a := i * keyLen
			b := (i + 1) * keyLen
			c := (i + 2) * keyLen

			dist += float64(HammingDist(in[a:b], in[b:c])) / float64(keyLen)
		}

		dist /= float64(blocks)
		fmt.Println("Len:", keyLen, "Dist:", dist)
		if dist < minDist {
			keysize = keyLen
			minDist = dist
		}
	}

	return keysize
}

func Chop(in []byte, blkSize int) [][]byte {
	blocks := [][]byte{}

	cnt := len(in) / blkSize

	for i := 0; i < cnt; i++ {
		start := i * blkSize
		end := start + blkSize
		fmt.Println(start, end)
		blocks = append(blocks, in[start:end])
	}

	return blocks
}

func Transpose(in [][]byte) [][]byte {
	transposed := [][]byte{}
	blkCount := len(in)
	blkSize := len(in[0])

	for i := 0; i < blkSize; i++ {
		transposed = append(transposed, make([]byte, blkCount))
		for j := 0; j < blkCount; j++ {
			transposed[i][j] = in[j][i]
		}
	}

	return transposed
}

func main() {
	ciphertext := utils.ReadBase64File("6.txt")

	keysize := FindKeySize(ciphertext)

	blocks := Chop(ciphertext, keysize)

	transposed := Transpose(blocks)

	var key []byte

	for _, blk := range transposed {
		_, k, _ := utils.CrackXOR(blk)
		key = append(key, k)
	}

	fmt.Println("Key:", string(key))
	plaintext := utils.XOR(ciphertext, key)
	fmt.Println(string(plaintext))

}
