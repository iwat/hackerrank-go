package linkedlist

import (
	"fmt"
	"strings"
)

func Partition(node *Node, pivot string) *Node {
	lHead := (*Node)(nil)
	lTail := (*Node)(nil)
	eHead := (*Node)(nil)
	eTail := (*Node)(nil)
	gHead := (*Node)(nil)

	for node != nil {
		next := node.Next
		cmp := strings.Compare(fmt.Sprintf("%v", node.Value), pivot)
		if cmp < 0 {
			if lTail == nil {
				lTail = node
			}
			node.Next = lHead
			lHead = node
		} else if cmp > 0 {
			node.Next = gHead
			gHead = node
		} else {
			if eTail == nil {
				eTail = node
			}
			node.Next = eHead
			eHead = node
		}
		node = next
	}

	if lTail != nil {
		if eHead != nil {
			lTail.Next = eHead
		} else if gHead != nil {
			lTail.Next = gHead
		}
	}

	if eTail != nil {
		if gHead != nil {
			eTail.Next = gHead
		}
	}

	if lHead != nil {
		return lHead
	} else if eHead != nil {
		return eHead
	} else {
		return gHead
	}
}
