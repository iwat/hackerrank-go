package search

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridlandMetro1(t *testing.T) {
	grid := NewGridlandMetro()
	grid.AddTrack(2, Track{1, 5})
	grid.AddTrack(2, Track{2, 4})
	grid.AddTrack(2, Track{8, 8})
	assert.Equal(t, big.NewInt(6), grid.SumTrack())
}

func TestGridlandMetro2(t *testing.T) {
	grid := NewGridlandMetro()
	grid.AddTrack(2, Track{2, 3})
	grid.AddTrack(3, Track{1, 4})
	grid.AddTrack(4, Track{4, 4})
	assert.Equal(t, big.NewInt(7), grid.SumTrack())
}
