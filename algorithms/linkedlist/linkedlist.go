package linkedlist

import (
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
}

func Create(value ...interface{}) *Node {
	if len(value) == 0 {
		return nil
	}
	return &Node{
		value[0],
		Create(value[1:]...),
	}
}

func (n *Node) String() string {
	if n.Next == nil {
		return fmt.Sprintf("%v", n.Value)
	} else {
		return fmt.Sprintf("%v %v", n.Value, n.Next.String())
	}
}
