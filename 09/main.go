package main

import "fmt"

func Pad(in []byte, blockSize int) []byte {
	pad := blockSize - (len(in) % blockSize)

	var out []byte
	out = append(out, in...)

	for i := 0; i < pad; i++ {
		out = append(out, byte(pad))
	}

	return out
}

func main() {
	in := []byte("YELLOW SUBMARINE")
	out := Pad(in, 20)

	fmt.Printf("%q\n", out)
}
