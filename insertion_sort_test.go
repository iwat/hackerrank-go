package algorithms

import (
	"math/rand"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	samples := generateSamples(20)
	InsertionSort(samples)
	verifySorted(samples, t)
}

func verifySorted(samples []int, t *testing.T) {
	for i := 1; i < len(samples); i++ {
		if samples[i] < samples[i-1] {
			t.Log(samples)
			t.Fatal("Bad result")
		}
	}
}

func generateSamples(len int) []int {
	samples := []int{}
	for i := 0; i < len; i++ {
		samples = append(samples, int(rand.Int31n(int32(len))))
	}
	return samples
}
