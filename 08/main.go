package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/tpiscitell/cryptopals/utils"
)

func ChallengeEight() (string, int) {
	f, _ := os.Open("8.txt")
	scanner := bufio.NewScanner(f)

	var answer string
	var maxRepeats int

	for scanner.Scan() {
		ciphertext := scanner.Text()
		b, _ := hex.DecodeString(ciphertext)

		repeats := utils.CountRepeats(b, 16)

		if repeats > maxRepeats {
			answer = ciphertext
			maxRepeats = repeats
		}
	}

	return answer, maxRepeats
}
func main() {
	fmt.Println(ChallengeEight())
}
