package main

import (
	"bytes"
)

type BinaryNode struct {
	data  string
	left  *BinaryNode
	right *BinaryNode
}

func BuildBinaryTreeInOrder(input ...string) *BinaryNode {
	if len(input)%2 == 0 || input[0] == " " {
		return nil
	}

	return &BinaryNode{
		data:  input[len(input)/2],
		left:  BuildBinaryTreeInOrder(input[0 : len(input)/2]...),
		right: BuildBinaryTreeInOrder(input[len(input)/2+1:]...),
	}
}

func BuildBinaryTreePreOrder(input ...string) *BinaryNode {
	if len(input)%2 == 0 || input[0] == " " {
		return nil
	}

	return &BinaryNode{
		data:  input[0],
		left:  BuildBinaryTreePreOrder(input[1 : len(input)/2+1]...),
		right: BuildBinaryTreePreOrder(input[len(input)/2+1:]...),
	}
}

func BuildBinaryTreePostOrder(input ...string) *BinaryNode {
	if len(input)%2 == 0 || input[0] == " " {
		return nil
	}

	return &BinaryNode{
		data:  input[len(input)-1],
		left:  BuildBinaryTreePostOrder(input[0 : (len(input)-1)/2]...),
		right: BuildBinaryTreePostOrder(input[(len(input)-1)/2 : len(input)-1]...),
	}
}

func (n *BinaryNode) String() string {
	buffer := bytes.Buffer{}
	n.TraverseInOrder(func(s string) {
		buffer.WriteString(s)
	})
	return buffer.String()
}

func (n *BinaryNode) TraverseInOrder(cb func(string)) {
	if n == nil {
		return
	}
	n.left.TraverseInOrder(cb)
	cb(n.data)
	n.right.TraverseInOrder(cb)
}

func (n *BinaryNode) TraversePreOrder(cb func(string)) {
	if n == nil {
		return
	}
	cb(n.data)
	n.left.TraversePreOrder(cb)
	n.right.TraversePreOrder(cb)
}

func (n *BinaryNode) TraversePostOrder(cb func(string)) {
	if n == nil {
		return
	}
	n.left.TraversePostOrder(cb)
	n.right.TraversePostOrder(cb)
	cb(n.data)
}
