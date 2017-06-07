package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaleOfTwoStacks(t *testing.T) {
	queue := NewTaleOfTwoStacks()
	queue.Enqueue(42)
	queue.Dequeue()
	queue.Enqueue(14)
	assert.Equal(t, 14, queue.Peek())
	queue.Enqueue(28)
	assert.Equal(t, 14, queue.Peek())
	queue.Enqueue(60)
	queue.Enqueue(78)
	queue.Dequeue()
	queue.Dequeue()
	for i := 0; i < 1000; i++ {
		queue.Enqueue(1)
	}
}
