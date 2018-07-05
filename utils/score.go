package utils

import (
	"math"
)

var idealFreqs = []float64{
	.0817, .0149, .0278, .0425, .1270, .0223, .0202, .0609, .0697, .0015, .0077, .0402, .0241,
	.0675, .0751, .0193, .0009, .0599, .0633, .0906, .0276, .0098, .0236, .0015, .0197, .0007}

var idealFreqSpace = []float64{
	0.0651738, 0.0124248, 0.0217339, 0.0349835, //'A', 'B', 'C', 'D',...
	0.1041442, 0.0197881, 0.0158610, 0.0492888,
	0.0558094, 0.0009033, 0.0050529, 0.0331490,
	0.0202124, 0.0564513, 0.0596302, 0.0137645,
	0.0008606, 0.0497563, 0.0515760, 0.0729357,
	0.0225134, 0.0082903, 0.0171272, 0.0013692,
	0.0145984, 0.0007836, 0.1918182,
}

func dotVec(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}

func cosine(a, b []float64) float64 {
	return dotVec(a, b) / (math.Sqrt(dotVec(a, a)) * math.Sqrt(dotVec(b, b)))
}

func scoreText(a []byte) float64 {
	cts := make([]int, 27)
	for _, ch := range a {
		if 'A' <= ch && ch <= 'Z' {
			cts[int(ch)-'A']++
		} else if 'a' <= ch && ch <= 'z' {
			cts[int(ch)-'a']++
		} else if ch == ' ' {
			cts[26]++
		} else if 9 <= ch && ch <= 11 {
			//whitespace is ok
		} else if ch < 32 || ch > 126 {
			// not ascii
			return 0.0
		}
	}
	amount := float64(len(a))
	freqs := make([]float64, 27)
	for i, c := range cts {
		freqs[i] = float64(c) / amount
	}

	return cosine(freqs, idealFreqSpace)
}
