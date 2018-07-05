package main

import "testing"

func TestChallengeThree(t *testing.T) {
	in := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expected := "Cooking MC's like a pound of bacon"

	out, _, _ := ChallengeThree(in)

	if string(out) != expected {
		t.Error("Expected:", expected, "Got:", string(out))
	}
}
