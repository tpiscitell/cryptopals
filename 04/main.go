package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tpiscitell/cryptopals/utils"
)

func ChallengeFour(scanner *bufio.Scanner) (string, string, float64) {
	minScore := 10000.0
	var ciphertext, plaintext string

	for scanner.Scan() {
		inStr := scanner.Text()
		guess, score := utils.CrackXOR(inStr)

		if score > 1 && score < minScore {
			minScore = score
			plaintext = guess
			ciphertext = inStr
		}
	}

	return ciphertext, plaintext, minScore
}

func main() {
	f, _ := os.Open("4.txt")
	defer f.Close()

	_, out, _ := ChallengeFour(bufio.NewScanner(f))

	fmt.Println(out)
}
