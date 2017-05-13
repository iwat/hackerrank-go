package linkedlist

func KthLast(head *Node, pos int) *Node {
	p1 := head
	p2 := head
	for i := 0; i < pos; i++ {
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
	}
	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
