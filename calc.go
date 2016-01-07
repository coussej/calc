package calc

// Abs return the absolute value of x
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the smallest integer from its arguments.
func Min(x ...int) (min int) {
	for i, v := range x {
		switch {
		case i == 0:
			min = v
		case v < min:
			min = v
		}
	}
	return
}

// Max returns the largest integer from its arguments.
func Max(x ...int) (max int) {
	for i, v := range x {
		switch {
		case i == 0:
			max = v
		case v > max:
			max = v
		}
	}
	return
}

// Pow returns x^y, the base-x exponential of y, with as condition y >= 0. If
// y < 0, function returns 0.
func Pow(x, y int) (pow int) {
	switch {
	case x == 0:
		pow = 0
	case y == 0:
		pow = 1
	case y < 0:
		pow = 0
	default:
		pow = 1
		for i := 0; i < y; i++ {
			pow *= x
		}
	}
	return
}
