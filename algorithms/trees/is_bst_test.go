package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBST(t *testing.T) {
	root := BuildBinaryTreeInOrder("a", "b", " ", "c", "d", "e", "f")
	assert.True(t, IsBST(root))
	root = BuildBinaryTreeInOrder("a", "b", " ", "d", "c", "e", "f")
	assert.False(t, IsBST(root))
	root = BuildBinaryTreeInOrder("a", "b", " ", "d", "c", "e", "e")
	assert.False(t, IsBST(root))
}
