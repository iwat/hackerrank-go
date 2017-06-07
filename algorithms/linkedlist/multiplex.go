package linkedlist

func Multiplex(head *Node) *Node {
	p1 := head
	p2 := head
	for p1 != nil {
		p1 = p1.Next.Next
		p2 = p2.Next
	}
	p1 = head
	for p2 != nil {
		tmp1 := p1.Next
		tmp2 := p2.Next
		p1.Next = p2
		if p2.Next != nil {
			p2.Next = tmp1
		}
		p1 = tmp1
		p2 = tmp2
	}
	return head
}
