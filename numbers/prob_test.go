package numbers

import (
	"math"
	"testing"
)

func TestNCX(t *testing.T) {
	if nCx(52, 5) != 2598960 {
		t.Errorf("Error in nCx. Expecting 2598960, got %f\n", nCx(52, 5))
	}
}

func TestBinomialProbability(t *testing.T) {
	if math.Abs(BinomialProbability(10, 5, 0.5)-0.24609) > 1e6 {
		t.Errorf("Error in BinomialProbability. Expecting 0.246094, got %f\n", BinomialProbability(10, 5, 0.5))
	}
}
