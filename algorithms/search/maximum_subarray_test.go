package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumSubarraySum(t *testing.T) {
	assert.Equal(t, 6, MaximumSubarraySum(7, 3, 3, 9, 9, 5))
}
