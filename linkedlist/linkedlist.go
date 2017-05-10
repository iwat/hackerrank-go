package linkedlist

type Node struct {
	Value string
	Next  *Node
}

func Create(value ...string) *Node {
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
		return n.Value
	} else {
		return n.Value + " " + n.Next.String()
	}
}
