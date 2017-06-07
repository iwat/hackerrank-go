package linkedlist

func RemoveDups(head *Node) {
	p1 := head
	for p1 != nil {
		p2 := p1.Next
		p2Parent := p1
		for p2 != nil {
			if p2.Value == p1.Value {
				p2Parent.Next = p2.Next
			} else {
				p2Parent = p2
			}
			p2 = p2.Next
		}
		p1 = p1.Next
	}
}
