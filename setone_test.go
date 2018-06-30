package main

import "testing"

func TestChallengeOne(t *testing.T) {
	in := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	out := ChallengeOne(in)

	if out != expected {
		t.Error("Expected:", expected, "Got:", out)
	}
}

func TestChallengeTwo(t *testing.T) {
	in := "1c0111001f010100061a024b53535009181c"
	key := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	out := ChallengeTwo(in, key)

	if out != expected {
		t.Error("Expected:", expected, "Got:", out)
	}
}
