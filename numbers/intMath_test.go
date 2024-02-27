package numbers

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	if Factorial(8) != 40320 {
		t.Errorf("error in TestFactorial. Expected 40320, got %d", Factorial(8))
	}
}
