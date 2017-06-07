package bfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestReach1(t *testing.T) {
	graph := NewGraph(4)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	distances := graph.DistancesFrom(1)
	assert.Equal(t, 6, distances[1])
	assert.Equal(t, 6, distances[2])
	assert.Equal(t, 0, distances[3])
}

func TestShortestReach2(t *testing.T) {
	graph := NewGraph(3)
	graph.AddEdge(2, 3)
	distances := graph.DistancesFrom(2)
	assert.Equal(t, 0, distances[0])
	assert.Equal(t, 6, distances[2])
}

func TestShortestReach3(t *testing.T) {
	graph := NewGraph(4)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	distances := graph.DistancesFrom(1)
	assert.Equal(t, 6, distances[1])
	assert.Equal(t, 12, distances[2])
	assert.Equal(t, 18, distances[3])
}

func TestShortestReach4(t *testing.T) {
	graph := NewGraph(4)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 1)
	distances := graph.DistancesFrom(1)
	assert.Equal(t, 6, distances[1])
	assert.Equal(t, 12, distances[2])
	assert.Equal(t, 6, distances[3])
}
