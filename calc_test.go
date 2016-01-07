package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	a := Abs(1)
	assert.Equal(t, 1, a, "Abs(1) should be 1")

	a = Abs(-1)
	assert.Equal(t, 1, a, "Abs(-1) should be 1")
}

func TestMin(t *testing.T) {
	m := Min(1, 2, 3)
	assert.Equal(t, 1, m, "Min of [1 2 3] should be 1")

	m = Min(-1, -2, -3)
	assert.Equal(t, -3, m, "Min of [-1 -2 -3] should be -3")
}

func TestMax(t *testing.T) {
	m := Max(1, 2, 3)
	assert.Equal(t, 3, m, "Max of [1 2 3] should be 3")

	m = Max(-1, -2, -3)
	assert.Equal(t, -1, m, "Max of [-1 -2 -3] should be -1")
}

func TestPow(t *testing.T) {
	p := Pow(2, 4)
	assert.Equal(t, 16, p, "2^2 should be 16")

	p = Pow(-2, 3)
	assert.Equal(t, -8, p, "-2^3 should be -8")

	p = Pow(0, 1000)
	assert.Equal(t, 0, p, "0^1000 should be 0")

	p = Pow(1000, 0)
	assert.Equal(t, 1, p, "1000^0 should be 0")
}
