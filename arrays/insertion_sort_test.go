package arrays

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	samples := generateSamples(20)
	InsertionSort(samples)
	assertSorted(t, samples)
}

func assertSorted(t *testing.T, samples []int) {
	for i := 1; i < len(samples); i++ {
		assert.True(t, samples[i] >= samples[i-1])
	}
}

func generateSamples(len int) []int {
	samples := []int{}
	for i := 0; i < len; i++ {
		samples = append(samples, int(rand.Int31n(int32(len))))
	}
	return samples
}
