package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunningMedian1(t *testing.T) {
	median := NewRunningMedian()
	assert.InDelta(t, 12.0, median.Add(12), 0.01)
	assert.InDelta(t, 8.0, median.Add(4), 0.01)
	assert.InDelta(t, 5.0, median.Add(5), 0.01)
	assert.InDelta(t, 4.5, median.Add(3), 0.01)
	assert.InDelta(t, 5.0, median.Add(8), 0.01)
	assert.InDelta(t, 6.0, median.Add(7), 0.01)
}

func TestRunningMedian2(t *testing.T) {
	median := NewRunningMedian()
	assert.InDelta(t, 1.0, median.Add(1), 0.01)
	assert.InDelta(t, 1.5, median.Add(2), 0.01)
	assert.InDelta(t, 2.0, median.Add(3), 0.01)
	assert.InDelta(t, 2.5, median.Add(4), 0.01)
	assert.InDelta(t, 3.0, median.Add(5), 0.01)
	assert.InDelta(t, 3.5, median.Add(6), 0.01)
	assert.InDelta(t, 4.0, median.Add(7), 0.01)
	assert.InDelta(t, 4.5, median.Add(8), 0.01)
	assert.InDelta(t, 5.0, median.Add(9), 0.01)
	assert.InDelta(t, 5.5, median.Add(10), 0.01)
}

func TestRunningMedian3(t *testing.T) {
	median := NewRunningMedian()
	assert.InDelta(t, 10.0, median.Add(10), 0.01)
	assert.InDelta(t, 9.5, median.Add(9), 0.01)
	assert.InDelta(t, 9.0, median.Add(8), 0.01)
	assert.InDelta(t, 8.5, median.Add(7), 0.01)
	assert.InDelta(t, 8.0, median.Add(6), 0.01)
	assert.InDelta(t, 7.5, median.Add(5), 0.01)
	assert.InDelta(t, 7.0, median.Add(4), 0.01)
	assert.InDelta(t, 6.5, median.Add(3), 0.01)
	assert.InDelta(t, 6.0, median.Add(2), 0.01)
	assert.InDelta(t, 5.5, median.Add(1), 0.01)
}

func TestRunningMedian4(t *testing.T) {
	median := NewRunningMedian()
	assert.InDelta(t, 94455.0, median.Add(94455), 0.01)
	assert.InDelta(t, 57505.0, median.Add(20555), 0.01)
	assert.InDelta(t, 20555.0, median.Add(20535), 0.01)
	assert.InDelta(t, 36840.0, median.Add(53125), 0.01)
	assert.InDelta(t, 53125.0, median.Add(73634), 0.01)
	assert.InDelta(t, 36840.0, median.Add(148), 0.01)
	assert.InDelta(t, 53125.0, median.Add(63772), 0.01)
	assert.InDelta(t, 36840.0, median.Add(17738), 0.01)
	assert.InDelta(t, 53125.0, median.Add(62995), 0.01)
}

func TestHeap(t *testing.T) {
	heap := NewHeap(1)
	heap.Add(6)
	heap.Add(5)
	heap.Add(4)
	heap.Add(3)
	heap.Add(2)
	heap.Add(1)
	heap.Add(6)
	heap.Add(5)
	heap.Add(4)
	heap.Add(3)
	heap.Add(2)
	heap.Add(1)
	assert.Equal(t, 1, heap.Poll())
	assert.Equal(t, 1, heap.Poll())
	assert.Equal(t, 2, heap.Poll())
	assert.Equal(t, 2, heap.Poll())
	assert.Equal(t, 3, heap.Poll())
	assert.Equal(t, 3, heap.Poll())
	assert.Equal(t, 4, heap.Poll())
	assert.Equal(t, 4, heap.Poll())
	assert.Equal(t, 5, heap.Poll())
	assert.Equal(t, 5, heap.Poll())
	assert.Equal(t, 6, heap.Poll())
	assert.Equal(t, 6, heap.Poll())
}
