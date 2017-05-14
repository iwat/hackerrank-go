package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuartiles(t *testing.T) {
	input := []int{3, 7, 8, 5, 12, 14, 21, 13, 18}
	q1, q2, q3 := Quartiles(input)
	assert.Equal(t, 6, q1)
	assert.Equal(t, 12, q2)
	assert.Equal(t, 16, q3)

	input = []int{3, 7, 8, 5, 12, 14, 13, 18}
	q1, q2, q3 = Quartiles(input)
	assert.Equal(t, 6, q1)
	assert.Equal(t, 10, q2)
	assert.Equal(t, 13, q3)

	input = []int{3, 7, 8, 5, 12, 14, 21, 15, 18, 14}
	q1, q2, q3 = Quartiles(input)
	assert.Equal(t, 7, q1)
	assert.Equal(t, 13, q2)
	assert.Equal(t, 15, q3)
}
