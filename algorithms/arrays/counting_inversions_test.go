package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingInversionss(t *testing.T) {
	data := []int{2, 4, 1}
	assert.Equal(t, 2, CountingInversions(data))
	assert.Equal(t, 1, data[0])
	assert.Equal(t, 2, data[1])
	assert.Equal(t, 4, data[2])

	data = []int{1, 1, 1, 2, 2}
	assert.Equal(t, 0, CountingInversions(data))
	assert.Equal(t, 1, data[0])
	assert.Equal(t, 1, data[1])
	assert.Equal(t, 1, data[2])
	assert.Equal(t, 2, data[3])
	assert.Equal(t, 2, data[4])

	data = []int{2, 1, 3, 1, 2}
	assert.Equal(t, 4, CountingInversions(data))
	assert.Equal(t, 1, data[0])
	assert.Equal(t, 1, data[1])
	assert.Equal(t, 2, data[2])
	assert.Equal(t, 2, data[3])
	assert.Equal(t, 3, data[4])

	data = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	assert.Equal(t, 45, CountingInversions(data))
	assert.Equal(t, 1, data[0])
	assert.Equal(t, 2, data[1])
	assert.Equal(t, 3, data[2])
	assert.Equal(t, 4, data[3])
	assert.Equal(t, 5, data[4])
	assert.Equal(t, 6, data[5])
	assert.Equal(t, 7, data[6])
	assert.Equal(t, 8, data[7])
	assert.Equal(t, 9, data[8])
	assert.Equal(t, 10, data[9])

	data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, 0, CountingInversions(data))
	assert.Equal(t, 1, data[0])
	assert.Equal(t, 2, data[1])
	assert.Equal(t, 3, data[2])
	assert.Equal(t, 4, data[3])
	assert.Equal(t, 5, data[4])
	assert.Equal(t, 6, data[5])
	assert.Equal(t, 7, data[6])
	assert.Equal(t, 8, data[7])
	assert.Equal(t, 9, data[8])
	assert.Equal(t, 10, data[9])
}
