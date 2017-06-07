package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteNode(t *testing.T) {
	head := Create("a", "b", "c", "d", "e", "f", "g", "h")
	node := head.Next.Next.Next.Next
	DeleteNode(head, node)
	assert.Equal(t, "a", head.Value)
	head = head.Next
	assert.Equal(t, "b", head.Value)
	head = head.Next
	assert.Equal(t, "c", head.Value)
	head = head.Next
	assert.Equal(t, "d", head.Value)
	head = head.Next
	assert.Equal(t, "f", head.Value)
	head = head.Next
	assert.Equal(t, "g", head.Value)
	head = head.Next
	assert.Equal(t, "h", head.Value)
	head = head.Next
	assert.Nil(t, head)
}
