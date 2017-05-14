package statistics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeanMedianMode(t *testing.T) {
	input := []int{64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060}
	mean, median, mode := MeanMedianMode(input)
	assert.InDelta(t, 43900.6, mean, 0.01)
	assert.InDelta(t, 44627.5, median, 0.01)
	assert.Equal(t, 4978, mode)
}
