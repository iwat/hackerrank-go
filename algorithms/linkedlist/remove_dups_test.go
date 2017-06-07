package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDups(t *testing.T) {
	head := Create("a", "b", "a", "c", "c", "d", "d")
	RemoveDups(head)

	assert.Equal(t, "a", head.Value)
	head = head.Next
	assert.Equal(t, "b", head.Value)
	head = head.Next
	assert.Equal(t, "c", head.Value)
	head = head.Next
	assert.Equal(t, "d", head.Value)
	head = head.Next
	assert.Nil(t, head)
}
