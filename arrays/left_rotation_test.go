package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateLeft(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	RotateLeft(data, 1)
	expected := []int{2, 3, 4, 5, 6, 1}
	for i := 0; i < len(data); i++ {
		assert.Equal(t, expected[i], data[i])
	}

	data = []int{1, 2, 3, 4, 5, 6}
	RotateLeft(data, 5)
	expected = []int{6, 1, 2, 3, 4, 5}
	for i := 0; i < len(data); i++ {
		assert.Equal(t, expected[i], data[i])
	}

	data = []int{1, 2, 3, 4, 5, 6}
	RotateLeft(data, 3)
	expected = []int{4, 5, 6, 1, 2, 3}
	for i := 0; i < len(data); i++ {
		assert.Equal(t, expected[i], data[i])
	}

	data = []int{1, 2, 3, 4, 5, 6}
	RotateLeft(data, 2)
	expected = []int{3, 4, 5, 6, 1, 2}
	for i := 0; i < len(data); i++ {
		assert.Equal(t, expected[i], data[i])
	}

	data = []int{1, 2, 3, 4, 5, 6, 7}
	RotateLeft(data, 3)
	expected = []int{4, 5, 6, 7, 1, 2, 3}
	for i := 0; i < len(data); i++ {
		assert.Equal(t, expected[i], data[i])
	}
}
