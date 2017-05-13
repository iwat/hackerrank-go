package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	stack := NewStack()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Push("d")
	stack.Push("e")
	stack.Push("f")
	stack.Push("g")

	assert.Equal(t, "g", stack.Pop())
	assert.Equal(t, "f", stack.Pop())
	assert.Equal(t, "e", stack.Pop())
	assert.Equal(t, "d", stack.Pop())
	assert.Equal(t, "c", stack.Pop())
	assert.Equal(t, "b", stack.Pop())
	assert.Equal(t, "a", stack.Pop())
	assert.Nil(t, stack.Pop())
}
