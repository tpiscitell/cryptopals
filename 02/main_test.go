package main

import "testing"

func TestChallengeTwo(t *testing.T) {
	in := "1c0111001f010100061a024b53535009181c"
	key := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	out := ChallengeTwo(in, key)

	if out != expected {
		t.Error("Expected:", expected, "Got:", out)
	}
}
