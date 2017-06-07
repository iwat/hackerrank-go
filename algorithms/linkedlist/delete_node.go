package linkedlist

func DeleteNode(head *Node, node *Node) {
	prev := head
	p := head.Next
	for p != node {
		prev = p
		p = p.Next
		if p == nil {
			return
		}
	}

	prev.Next = p.Next
	p.Next = nil // Not necessary
}
