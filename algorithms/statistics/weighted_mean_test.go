package statistics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeightedMean(t *testing.T) {
	input := []int{10, 40, 30, 50, 20}
	weight := []int{1, 2, 3, 4, 5}
	mean := WeightedMean(input, weight)
	assert.InDelta(t, 32.0, mean, 0.01)
}
