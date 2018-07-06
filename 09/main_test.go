package main

import "testing"

func TestPad(t *testing.T) {
	in := []byte("YELLOW SUBMARINE")

	out := Pad(in, 20)

	if len(out) != 20 {
		t.Error("Output should have a length of 20", out)
	}

	for _, b := range out[16:] {
		if b != 0x04 {
			t.Error("Output should be padded with 0x04", out)

		}
	}
}
