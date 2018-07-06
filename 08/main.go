package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("8.txt")
	scanner := bufio.NewScanner(f)

	var answer string
	var maxRepeats int

	for scanner.Scan() {
		counter := make(map[string]int)
		ciphertext := scanner.Text()
		blockSize := 16
		blockCount := len(ciphertext) / blockSize

		for i := 0; i < blockCount; i++ {
			start := i * blockSize
			end := start + blockSize

			slice := ciphertext[start:end]
			counter[slice]++
		}

		var repeats int
		for _, v := range counter {
			if v > 1 {
				repeats += v
			}
		}

		if repeats > maxRepeats {
			answer = ciphertext
			maxRepeats = repeats
		}
	}

	fmt.Println(answer, maxRepeats)
}
