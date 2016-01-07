package calc

// Min returns the smalles integer from its arguments.
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
