package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	input := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	swaps := BubbleSort(input)
	assert.Equal(t, 1, input[0])
	assert.Equal(t, 9, input[8])
	assert.Equal(t, 36, swaps)

	input = []int{1, 2, 3}
	swaps = BubbleSort(input)
	assert.Equal(t, 1, input[0])
	assert.Equal(t, 3, input[2])
	assert.Equal(t, 0, swaps)
}
