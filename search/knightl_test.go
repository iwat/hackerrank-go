package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnightL(t *testing.T) {
	board := NewKnightL(5)
	assert.Equal(t, 4, board.Count(1, 1), "(1, 1)")
	assert.Equal(t, 4, board.Count(1, 2), "(1, 2)")
	assert.Equal(t, 2, board.Count(1, 3), "(1, 3)")
	assert.Equal(t, 8, board.Count(1, 4), "(1, 4)")
	assert.Equal(t, 4, board.Count(2, 1), "(2, 1)")
	assert.Equal(t, 2, board.Count(2, 2), "(2, 2)")
	assert.Equal(t, 4, board.Count(2, 3), "(2, 3)")
	assert.Equal(t, 4, board.Count(2, 4), "(2, 4)")
	assert.Equal(t, 2, board.Count(3, 1), "(3, 1)")
	assert.Equal(t, 4, board.Count(3, 2), "(3, 2)")
	assert.Equal(t, -1, board.Count(3, 3), "(3, 3)")
	assert.Equal(t, -1, board.Count(3, 4), "(3, 4)")
	assert.Equal(t, 8, board.Count(4, 1), "(4, 1)")
	assert.Equal(t, 4, board.Count(4, 2), "(4, 2)")
	assert.Equal(t, -1, board.Count(4, 3), "(4, 3)")
	assert.Equal(t, 1, board.Count(4, 4), "(4, 4)")
}
