package main

import (
	"encoding/hex"
	"fmt"

	"github.com/tpiscitell/cryptopals/utils"
)

func ChallengeFive(in string) string {
	inBytes := []byte(in)
	key := []byte("ICE")

	return hex.EncodeToString(utils.XOR(inBytes, key))
}

func main() {
	fmt.Println(ChallengeFive("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"))
}
