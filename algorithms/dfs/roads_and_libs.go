package dfs

import (
	"fmt"
)

type City struct {
	name string
	adj  []*City
}

func (c *City) String() string {
	return c.name
}

type RoadsAndLibs struct {
	cities      []*City
	libraryCost int
	roadCost    int
}

func NewRoadsAndLibs(nCities, nRoads, libraryCost, roadCost int) *RoadsAndLibs {
	cities := make([]*City, nCities)
	for i := 0; i < nCities; i++ {
		cities[i] = &City{fmt.Sprintf("%d", i+1), nil}
	}
	return &RoadsAndLibs{cities, libraryCost, roadCost}
}

func (r *RoadsAndLibs) Link(from, to int) {
	r.cities[from-1].adj = append(r.cities[from-1].adj, r.cities[to-1])
	r.cities[to-1].adj = append(r.cities[to-1].adj, r.cities[from-1])
}

func (r *RoadsAndLibs) Cost() int {
	if r.libraryCost <= r.roadCost {
		return len(r.cities) * r.libraryCost
	}

	visited := make(map[*City]bool, len(r.cities))
	for i := 0; i < len(r.cities); i++ {
		visited[r.cities[i]] = false
	}

	nLibraries := 0
	nRoads := 0

	for i := 0; i < len(r.cities); i++ {
		if visited[r.cities[i]] {
			continue
		}

		visited[r.cities[i]] = true
		nLibraries++

		stack := []*City{r.cities[i]}

		nVisits := 0
		for len(stack) > 0 {
			head := stack[0]
			stack = stack[1:]

			for _, c := range head.adj {
				if !visited[c] {
					stack = append(stack, c)
					visited[c] = true
				}
			}

			nVisits++
		}

		nRoads += nVisits - 1
	}

	return nLibraries*r.libraryCost + nRoads*r.roadCost
}
