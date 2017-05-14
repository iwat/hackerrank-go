package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildBinaryTreeInOrder(t *testing.T) {
	root := BuildBinaryTreeInOrder("a", "b", " ", "c", "d", "e", "f")

	buffer := bytes.Buffer{}
	root.TraverseInOrder(func(s string) {
		buffer.WriteString(s)
	})
	assert.Equal(t, "abcdef", buffer.String())

	buffer.Reset()
	root.TraversePreOrder(func(s string) {
		buffer.WriteString(s)
	})
	assert.Equal(t, "cbaedf", buffer.String())

	buffer.Reset()
	root.TraversePostOrder(func(s string) {
		buffer.WriteString(s)
	})
	assert.Equal(t, "abdfec", buffer.String())
}

func TestBuildBinaryTreePreOrder(t *testing.T) {
	root := BuildBinaryTreePreOrder("c", "b", "a", " ", "e", "d", "f")

	buffer := bytes.Buffer{}
	root.TraverseInOrder(func(s string) {
		buffer.WriteString(s)
	})
	assert.Equal(t, "abcdef", buffer.String())

	buffer.Reset()
	root.TraversePreOrder(func(s string) {
		buffer.WriteString(s)
	})
	assert.Equal(t, "cbaedf", buffer.String())

	buffer.Reset()
	root.TraversePostOrder(func(s string) {
		buffer.WriteString(s)
	})
	assert.Equal(t, "abdfec", buffer.String())
}

func TestBuildBinaryTreePostOrder(t *testing.T) {
	root := BuildBinaryTreePostOrder("a", " ", "b", "d", "f", "e", "c")

	buffer := bytes.Buffer{}
	root.TraverseInOrder(func(s string) {
		buffer.WriteString(s)
	})
	assert.Equal(t, "abcdef", buffer.String())

	buffer.Reset()
	root.TraversePreOrder(func(s string) {
		buffer.WriteString(s)
	})
	assert.Equal(t, "cbaedf", buffer.String())

	buffer.Reset()
	root.TraversePostOrder(func(s string) {
		buffer.WriteString(s)
	})
	assert.Equal(t, "abdfec", buffer.String())
}

func TestTree(t *testing.T) {
	fNode := &BinaryNode{"f", nil, nil}
	dNode := &BinaryNode{"d", nil, nil}

	aNode := &BinaryNode{"a", nil, nil}

	bNode := &BinaryNode{"b", aNode, nil}
	eNode := &BinaryNode{"e", dNode, fNode}

	root := &BinaryNode{"c", bNode, eNode}

	buffer := bytes.Buffer{}
	root.TraverseInOrder(func(s string) {
		buffer.WriteString(s)
	})
	assert.Equal(t, "abcdef", buffer.String())

	buffer.Reset()
	root.TraversePreOrder(func(s string) {
		buffer.WriteString(s)
	})
	assert.Equal(t, "cbaedf", buffer.String())

	buffer.Reset()
	root.TraversePostOrder(func(s string) {
		buffer.WriteString(s)
	})
	assert.Equal(t, "abdfec", buffer.String())
}
