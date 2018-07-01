package main

import (
	"errors"
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

func HammingDist(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("a and b must be the same length")
	}

	xor := utils.XOR(a, b)
	sum := 0

	for _, i := range xor {
		sum += HammingWeight(i)
	}

	return sum, nil
}

func HammingDistStr(a, b string) (int, error) {
	return HammingDist([]byte(a), []byte(b))
}

func main() {
	fmt.Println(HammingDistStr("this is a test", "wokka wokka!!!"))
}
