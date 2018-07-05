package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/tpiscitell/cryptopals/utils"
)

func ChallengeFour(scanner *bufio.Scanner) (string, string, float64) {
	var maxScore float64
	var ciphertext, plaintext string

	for scanner.Scan() {
		inStr := scanner.Text()
		in, _ := hex.DecodeString(inStr)
		guess, _, score := utils.CrackXOR(in)

		if score > maxScore {
			maxScore = score
			plaintext = string(guess)
			ciphertext = inStr
		}
	}

	return ciphertext, plaintext, maxScore
}

func main() {
	f, _ := os.Open("4.txt")
	defer f.Close()

	_, out, _ := ChallengeFour(bufio.NewScanner(f))

	fmt.Println(out)
}
