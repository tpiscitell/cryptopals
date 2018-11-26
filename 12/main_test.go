package main

import (
	"testing"
)

func TestDiscoverBlockSize(t *testing.T) {
	blockSize := discoverBlockSize()
	if blockSize != 16 {
		t.Errorf("Block size is %d. Got: %d", 16, blockSize)
	}
}

func TestSliceEq(t *testing.T) {
	if !sliceEq([]byte{0x11, 0x22, 0x33}, []byte{0x11, 0x22, 0x33}) {
		t.Errorf("slices should be equal but are not")
	}
}
