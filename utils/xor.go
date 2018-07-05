package utils

func CrackXOR(inBytes []byte) ([]byte, byte, float64) {
	var maxScore float64
	var answer []byte
	var key byte

	for i := 0; i < 256; i++ {
		k := byte(i)
		outBytes := singleXOR(inBytes, k)
		score := scoreText(outBytes)

		if score > maxScore {
			key = k
			maxScore = score
			answer = outBytes
		}
	}

	return answer, key, maxScore
}

func XOR(in, key []byte) []byte {
	n := len(in)
	kn := len(key)
	out := make([]byte, n)

	for i := 0; i < n; i++ {
		out[i] = in[i] ^ key[i%kn]
	}

	return out
}

func singleXOR(in []byte, key byte) []byte {
	n := len(in)
	var out []byte

	for i := 0; i < n; i++ {
		out = append(out, in[i]^key)
	}

	return out
}
