package main

import (
	"strings"

	"github.com/iwat/algorithms/linkedlist"
)

type StackMin struct {
	head *linkedlist.Node
	min  *linkedlist.Node
}

func NewStackMin() *StackMin {
	return &StackMin{nil, nil}
}

func (s *StackMin) Push(data string) {
	oldHead := s.head
	s.head = &linkedlist.Node{data, nil}
	s.head.Next = oldHead

	min := data
	if s.min != nil {
		if strings.Compare(data, s.min.Value) > 0 {
			min = s.min.Value
		}
	}
	oldMin := s.min
	s.min = &linkedlist.Node{min, nil}
	s.min.Next = oldMin
}

func (s *StackMin) Pop() string {
	if s.head == nil {
		return ""
	}
	value := s.head.Value
	s.head = s.head.Next
	s.min = s.min.Next
	return value
}

func (s *StackMin) Min() string {
	if s.min == nil {
		return ""
	}
	return s.min.Value
}
