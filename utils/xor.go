package utils

import (
	"encoding/hex"
)

func CrackXOR(in string) (string, float64) {
	inBytes, _ := hex.DecodeString(in)

	minScore := 100000.00
	var answer string

	for i := 0; i < 256; i++ {
		k := byte(i)
		outBytes := singleXOR(inBytes, k)
		score := ScoreString(outBytes)

		if score > 1 && score < minScore {
			minScore = score
			answer = string(outBytes)
		}
	}

	return answer, minScore
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
	out := make([]byte, n)

	for i := 0; i < n; i++ {
		out[i] = in[i] ^ key
	}

	return out
}

func ScoreString(in []byte) float64 {
	expectedFreqs := []float64{
		0.0651738, 0.0124248, 0.0217339, 0.0349835, //'A', 'B', 'C', 'D',...
		0.1041442, 0.0197881, 0.0158610, 0.0492888,
		0.0558094, 0.0009033, 0.0050529, 0.0331490,
		0.0202124, 0.0564513, 0.0596302, 0.0137645,
		0.0008606, 0.0497563, 0.0515760, 0.0729357,
		0.0225134, 0.0082903, 0.0171272, 0.0013692,
		0.0145984, 0.0007836, 0.1918182, //'Y', 'Z', ' '
	}

	charCounts := make([]int, 27)
	charFreqs := make([]float64, 27)
	totalCount := 0

	for _, c := range in {
		//to lower all the ascii capital letters
		if 'A' <= c && c <= 'Z' {
			c -= 32
		}

		if 'a' <= c && c <= 'z' {
			charCounts[int(c)-'a']++
			totalCount++
		} else if c == ' ' {
			charCounts[26]++
			totalCount++
		} else if 10 <= c && c <= 126 {

		} else {
			return 666 // we got some garbage
		}
	}

	if totalCount == 0 {
		return 0
	}

	var chi2 float64
	for i := 0; i < 27; i++ {
		charFreqs[i] = float64(charCounts[i]) / float64(totalCount)
		chi2 += (charFreqs[i] - expectedFreqs[i]) * (charFreqs[i] - expectedFreqs[i]) / expectedFreqs[i]
	}

	return chi2
}
