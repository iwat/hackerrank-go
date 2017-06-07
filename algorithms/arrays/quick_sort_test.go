package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	data := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	QuickSort(data, 0, len(data))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, data)
}
