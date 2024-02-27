package numbers

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func AbsInt(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	var f int = 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	return f
}
