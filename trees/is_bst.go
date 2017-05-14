package trees

import (
	"strings"
)

func IsBST(n *BinaryNode) bool {
	return isBST(n, "a", "z")
}

func isBST(n *BinaryNode, min, max string) bool {
	if n == nil {
		return true
	}
	if strings.Compare(n.data, min) < 0 {
		return false
	}
	if strings.Compare(n.data, max) > 0 {
		return false
	}
	return isBST(n.left, min, n.data) && isBST(n.right, n.data, max)
}
