package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStackMin(t *testing.T) {
	stack := NewStackMin()

	stack.Push("b")
	assert.Equal(t, "b", stack.Min())

	stack.Push("a")
	assert.Equal(t, "a", stack.Min())

	stack.Push("c")
	assert.Equal(t, "a", stack.Min())

	assert.Equal(t, "c", stack.Pop())
	assert.Equal(t, "a", stack.Min())

	assert.Equal(t, "a", stack.Pop())
	assert.Equal(t, "b", stack.Min())

	assert.Equal(t, "b", stack.Pop())
	assert.Equal(t, "", stack.Min())
	assert.Equal(t, "", stack.Pop())
}
