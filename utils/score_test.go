package utils

import (
	"testing"
)

func TestDotVec(t *testing.T) {
	a := []float64{1, 2, 3}
	b := []float64{4, 5, 6}
	expected := 32.0

	out := dotVec(a, b)

	if out != expected {
		t.Error("Expected:", expected, "Got:", out)
	}

}
func TestCosine(t *testing.T) {
	a := []float64{1.0, 2.0, 34.2234, 0.0}
	b := []float64{1.0, 2.0, 34.2234, 0.0}

	score := cosine(a, b)
	if int(score) != 1 {
		t.Error("Expected: 1", "Got:", score)
	}
}
