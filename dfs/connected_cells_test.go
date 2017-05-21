package dfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectedCells(t *testing.T) {
	matrix := NewMatrix(4, 4)
	matrix.FillRow(0, 1, 1, 0, 0)
	matrix.FillRow(1, 0, 1, 1, 0)
	matrix.FillRow(2, 0, 0, 1, 0)
	matrix.FillRow(3, 1, 0, 0, 0)
	assert.Equal(t, 5, matrix.LargestRegionSize())
}

func TestConnectedCells2(t *testing.T) {
	matrix := NewMatrix(5, 4)
	matrix.FillRow(0, 0, 0, 1, 1)
	matrix.FillRow(1, 0, 0, 1, 0)
	matrix.FillRow(2, 0, 1, 1, 0)
	matrix.FillRow(3, 0, 1, 0, 0)
	matrix.FillRow(4, 1, 1, 0, 0)
	assert.Equal(t, 8, matrix.LargestRegionSize())
}

func TestConnectedCells3(t *testing.T) {
	matrix := NewMatrix(7, 5)
	matrix.FillRow(0, 1, 1, 1, 0, 1)
	matrix.FillRow(1, 0, 0, 1, 0, 0)
	matrix.FillRow(2, 1, 1, 0, 1, 0)
	matrix.FillRow(3, 0, 1, 1, 0, 0)
	matrix.FillRow(4, 0, 0, 0, 0, 0)
	matrix.FillRow(5, 0, 1, 0, 0, 0)
	matrix.FillRow(6, 0, 0, 1, 1, 0)
	assert.Equal(t, 9, matrix.LargestRegionSize())
}

func TestConnectedCells4(t *testing.T) {
	matrix := NewMatrix(7, 8)
	matrix.FillRow(0, 1, 0, 0, 1, 0, 1, 0, 0)
	matrix.FillRow(1, 0, 0, 0, 0, 0, 0, 0, 1)
	matrix.FillRow(2, 1, 0, 1, 0, 1, 0, 0, 0)
	matrix.FillRow(3, 0, 0, 0, 0, 0, 0, 1, 0)
	matrix.FillRow(4, 1, 0, 0, 1, 0, 0, 0, 0)
	matrix.FillRow(5, 0, 0, 0, 0, 0, 0, 0, 1)
	matrix.FillRow(6, 0, 1, 0, 0, 0, 1, 0, 0)
	assert.Equal(t, 1, matrix.LargestRegionSize())
}
