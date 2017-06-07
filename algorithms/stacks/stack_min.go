package stacks

import (
	"fmt"
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
		value := fmt.Sprintf("%v", s.min.Value)
		if strings.Compare(data, value) > 0 {
			min = value
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
	return fmt.Sprintf("%v", value)
}

func (s *StackMin) Min() string {
	if s.min == nil {
		return ""
	}
	return fmt.Sprintf("%v", s.min.Value)
}
