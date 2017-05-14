package stacks

import (
	"github.com/iwat/algorithms/linkedlist"
)

type Stack struct {
	head *linkedlist.Node
}

func NewStack() *Stack {
	return &Stack{nil}
}

func (s *Stack) Peek() interface{} {
	return s.head.Value
}

func (s *Stack) Push(data interface{}) {
	oldHead := s.head
	s.head = &linkedlist.Node{data, nil}
	s.head.Next = oldHead
}

func (s *Stack) Pop() interface{} {
	if s.head == nil {
		return ""
	}
	value := s.head.Value
	s.head = s.head.Next
	return value
}
