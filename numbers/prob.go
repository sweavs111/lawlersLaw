package numbers

import (
	"fmt"
	"log"
	"math"
	"overUnderModel/numbers/logspace"
)

// BinomialDistLog returns log(BinomialDist), where log is the natural logarithm.
// This is ideal for very small probabilities to avoid underflow.
func BinomialDistLog(n int, k int, p float64) float64 {
	coefficient := BinomCoefficientLog(n, k)
	expression := BinomialExpressionLog(n, k, p)
	return logspace.Multiply(coefficient, expression)
}

// BinomialExpressionLog returns p^n * (1 - p)^n-k, which is also referred to as the binomial expression. The answer is provided in logSpace (.
func BinomialExpressionLog(n int, k int, p float64) float64 {
	s := logspace.Pow(math.Log(p), float64(k))
	f := logspace.Pow(math.Log(1.0-p), float64(n-k))
	return logspace.Multiply(s, f)
}

// BinomCoefficientLog returns log(n choose k), where log is the natural logarithm.
// Ideal for large numbers as this raises the overflow ceiling considerably.
func BinomCoefficientLog(n int, k int) float64 {
	if n < 0 || k < 0 || k > n {
		log.Fatalf("The binomial coefficient call could not be handled: n=%d and k=%d\n", n, k)
	}
	if n-k > k {
		k = n - k
	}
	// this special case is handled here so that we don't ask for negative memory for denom
	if k == n {
		return 0.0
	}
	var x, y int
	var numer, denom float64 = 0.0, 0.0
	for x = k + 1; x < n+1; x++ {
		numer = logspace.Multiply(numer, math.Log(float64(x)))
	}
	for y = 2; y < n-k+1; y++ {
		denom = logspace.Multiply(denom, math.Log(float64(y)))
	}
	return logspace.Divide(numer, denom)
}

func BinomialProbability(n, x int, p float64) float64 {
	c := nCx(n, x)
	b := math.Pow(p, float64(x)) * math.Pow(1-p, float64(n-x))
	fmt.Println("30 chose 15: ", c)
	fmt.Println("b: ", b)
	return c * b
}

func BinomialDistribution(n int, p float64) []float64 {
	var slc []float64
	for x := 0; x <= n; x++ {
		slc = append(slc, BinomialProbability(n, x, p))
	}
	return slc
}

func nCx(n, x int) float64 {
	var numerator int = 1

	for i := n; i > MaxInt(n-x, x); i-- {
		numerator *= i
	}
	fmt.Println("numerator: ", numerator)
	fmt.Println("denominator: ", Factorial(MinInt(n-x, x)))
	return float64(numerator) / float64(Factorial(MinInt(n-x, x)))
}
