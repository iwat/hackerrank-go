package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()

	stack.Push("b")
	assert.Equal(t, "b", stack.Peek())
	stack.Push("a")
	assert.Equal(t, "a", stack.Peek())
	stack.Push("c")
	assert.Equal(t, "c", stack.Pop())
	assert.Equal(t, "a", stack.Pop())
	assert.Equal(t, "b", stack.Pop())
	assert.Equal(t, "", stack.Pop())
}
