package calc

import "math"

// Abs return the absolute value of x
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Ceil returns the first integer higher than the given float64.
func Ceil(x float64) int {
	return int(math.Ceil(x))
}

// Copysign returns a value with the magnitude of x and the sign of y
func Copysign(x, y int) int {
	if y < 0 {
		return -Abs(x)
	}
	return Abs(x)
}

// Floor returns the first integer lower than the given float64
func Floor(x float64) int {
	return int(math.Floor(x))
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

// Pow returns x^y, the base-x exponential of y, with as condition y >= 0. If
// y < 0, function returns 0.
func Pow(x, y int) (pow int) {
	switch {
	case x == 0 || y < 0:
		pow = 0
	case y == 0:
		pow = 1
	default:
		pow = 1
		for i := 0; i < y; i++ {
			pow *= x
		}
	}
	return
}

// Round returns the correctly rounded int from the given float64 x
func Round(x float64) int {
	return Floor(x + .5)
}
