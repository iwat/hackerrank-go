package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLast(t *testing.T) {
	head := Create("a", "b", "c", "d", "e")
	assert.Equal(t, "e", KthLast(head, 1).Value)
	assert.Equal(t, "d", KthLast(head, 2).Value)
	assert.Equal(t, "c", KthLast(head, 3).Value)
	assert.Equal(t, "b", KthLast(head, 4).Value)
	assert.Equal(t, "a", KthLast(head, 5).Value)
	assert.Nil(t, KthLast(head, 6))
}
