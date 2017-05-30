package search

import (
	"fmt"
	"math/big"
)

type GridlandMetro struct {
	rows map[int][]Track
}

func NewGridlandMetro() *GridlandMetro {
	return &GridlandMetro{make(map[int][]Track)}
}

func (g *GridlandMetro) AddTrack(row int, newTrack Track) {
	if _, ok := g.rows[row]; !ok {
		g.rows[row] = []Track{newTrack}
		return
	}
	for i, existingTrack := range g.rows[row] {
		if existingTrack.Contains(newTrack.from) || existingTrack.Contains(newTrack.to) {
			existingTrack.from = min(existingTrack.from, newTrack.from)
			existingTrack.to = max(existingTrack.to, newTrack.to)
			g.rows[row][i] = existingTrack
			return
		}
	}
	g.rows[row] = append(g.rows[row], newTrack)
}

func (g *GridlandMetro) SumTrack() *big.Int {
	sum := big.NewInt(0)
	for _, tracks := range g.rows {
		for _, track := range tracks {
			sum = sum.Add(sum, big.NewInt(int64(track.to-track.from+1)))
		}
	}
	return sum
}

type Track struct {
	from int
	to   int
}

func (t Track) Contains(point int) bool {
	return t.from <= point && point <= t.to
}

func (t Track) String() string {
	return fmt.Sprintf("[%d..%d]", t.from, t.to)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
