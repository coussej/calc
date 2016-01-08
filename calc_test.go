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

func TestCeil(t *testing.T) {
	c := Ceil(3.25)
	assert.Equal(t, 4, c, "Ceil(3.25) should be 4")

	c = Ceil(-3.25)
	assert.Equal(t, -3, c, "Ceil(-3.25) should be -3")
}

func TestCopysign(t *testing.T) {
	c := Copysign(12, -2)
	assert.Equal(t, -12, c, "Copysign(12,-2) should be -12")

	c = Copysign(-12, 2)
	assert.Equal(t, 12, c, "Copysign(-12,2) should be 12")
}

func TestFloor(t *testing.T) {
	f := Floor(3.25)
	assert.Equal(t, 3, f, "Ceil(3.25) should be 4")

	f = Floor(-3.25)
	assert.Equal(t, -4, f, "Ceil(-3.25) should be -3")
}

func TestMax(t *testing.T) {
	m := Max(1, 2, 3)
	assert.Equal(t, 3, m, "Max of [1 2 3] should be 3")

	m = Max(-1, -2, -3)
	assert.Equal(t, -1, m, "Max of [-1 -2 -3] should be -1")
}

func TestMin(t *testing.T) {
	m := Min(1, 2, 3)
	assert.Equal(t, 1, m, "Min of [1 2 3] should be 1")

	m = Min(-1, -2, -3)
	assert.Equal(t, -3, m, "Min of [-1 -2 -3] should be -3")
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

func TestRound(t *testing.T) {
	r := Round(0.5)
	assert.Equal(t, 1, r, "Round(0.5) should be 1")

	r = Round(0.49999999999999)
	assert.Equal(t, 0, r, "Round(0.49999999999999) should be 0")

	r = Round(-0.5)
	assert.Equal(t, 0, r, "Round(-0.5) should be 0")

	r = Round(-0.49999999999999)
	assert.Equal(t, 0, r, "Round(-0.49999999999999) should be 0")

	r = Round(-0.50000000000001)
	assert.Equal(t, -1, r, "Round(-0.50000000000001) should be -1")
}
