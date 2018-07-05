package main

import (
	"bufio"
	"os"
	"testing"
)

func TestChallengeFour(t *testing.T) {
	f, _ := os.Open("4.txt")
	defer f.Close()

	expected := "Now that the party is jumping\n"

	cipher, out, score := ChallengeFour(bufio.NewScanner(f))

	if out != expected {
		t.Error("Expected: ", expected, "Got:", out, "Cipher:", cipher, "Score:", score)
	}
}
