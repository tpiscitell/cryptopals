package main

import "testing"

func TestHammingWeight(t *testing.T) {
	wt := HammingWeight(byte(0xff))

	if wt != 8 {
		t.Error("Weight of 0xff should be 8. Got:", wt)
	}

	wt = HammingWeight(byte(0x00))

	if wt != 0 {
		t.Error("Weight of 0x00 should be 0. Got", wt)
	}

	wt = HammingWeight(byte(0x0f))

	if wt != 4 {
		t.Error("Weight of 0x0f should be 4. Got:", wt)
	}
}

func TestHammingDist(t *testing.T) {
	a := "this is a test"
	b := "wokka wokka!!!"

	diff, _ := HammingDistStr(a, b)

	if diff != 37 {
		t.Error("Dist should be 37. Got:", diff)
	}
}
