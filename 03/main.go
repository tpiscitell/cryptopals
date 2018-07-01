package main

import (
	"fmt"

	"github.com/tpiscitell/cryptopals/utils"
)

func ChallengeThree(in string) (string, float64) {
	return utils.CrackXOR(in)
}

func main() {
	in := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	out, _ := ChallengeThree(in)
	fmt.Println(out)
}
