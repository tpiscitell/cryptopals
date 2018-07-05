package utils

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestXOR(t *testing.T) {
	a := []byte{0x13, 0xed}
	b := []byte{0x00}

	out := XOR(a, b)

	if !bytes.Equal(a, out) {
		t.Error("Expected:", a, "Got:", out)
	}

	out = XOR(a, []byte{0xff})
	expected := []byte{0xec, 0x12}

	if !bytes.Equal(expected, out) {
		t.Error("Expected:", expected, "Got:", out)
	}

}

func TestSingleXOR(t *testing.T) {
	a := []byte{0x13, 0xed}

	out := singleXOR(a, byte(0x00))

	if !bytes.Equal(a, out) {
		t.Error("Expected:", a, "Got:", out)
	}

	out = singleXOR(a, byte(0xff))
	expected := []byte{0xec, 0x12}

	if !bytes.Equal(expected, out) {
		t.Error("Expected:", expected, "Got:", out)
	}
}

func TestCrackXOR(t *testing.T) {
	in, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	expected := "Cooking MC's like a pound of bacon"

	out, _, _ := CrackXOR(in)

	if string(out) != expected {
		t.Error("Expected:", expected, "Got:", string(out))
	}
}
