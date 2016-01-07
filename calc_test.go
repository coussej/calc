package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	m := Min(1, 2, 3)
	assert.True(t, m == 1, "Min of [1 2 3] should be 1")

	m = Min(-1, -2, -3)
	assert.True(t, m == -3, "Min of [-1 -2 -3] should be -3")
}

func TestMax(t *testing.T) {
	m := Max(1, 2, 3)
	assert.True(t, m == 3, "Max of [1 2 3] should be 3")

	m = Max(-1, -2, -3)
	assert.True(t, m == -1, "Max of [-1 -2 -3] should be -1")
}
