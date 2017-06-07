package search

import (
	"fmt"
)

type KnightL struct {
	size int
}

type Point struct {
	ndx  int
	cost int
	size int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d c %d)", p.ndx%p.size, p.ndx/p.size, p.cost)
}

func (p Point) nextMoves(x, y int) []Point {
	posX := p.ndx % p.size
	posY := p.ndx / p.size
	next := []Point{}
	if posX+x < p.size {
		if posY+y < p.size {
			next = append(next, Point{(posY+y)*p.size + (posX + x), p.cost + 1, p.size})
		}
		if posY-y >= 0 {
			next = append(next, Point{(posY-y)*p.size + (posX + x), p.cost + 1, p.size})
		}
	}
	if posX-x >= 0 {
		if posY+y < p.size {
			next = append(next, Point{(posY+y)*p.size + (posX - x), p.cost + 1, p.size})
		}
		if posY-y >= 0 {
			next = append(next, Point{(posY-y)*p.size + (posX - x), p.cost + 1, p.size})
		}
	}
	if posX+y < p.size {
		if posY+x < p.size {
			next = append(next, Point{(posY+x)*p.size + (posX + y), p.cost + 1, p.size})
		}
		if posY-x >= 0 {
			next = append(next, Point{(posY-x)*p.size + (posX + y), p.cost + 1, p.size})
		}
	}
	if posX-y >= 0 {
		if posY+x < p.size {
			next = append(next, Point{(posY+x)*p.size + (posX - y), p.cost + 1, p.size})
		}
		if posY-x >= 0 {
			next = append(next, Point{(posY-x)*p.size + (posX - y), p.cost + 1, p.size})
		}
	}
	fmt.Println("next", next)
	return next
}

func NewKnightL(size int) KnightL {
	return KnightL{size}
}

func (k KnightL) Count(x, y int) int {
	queue := []Point{{0, 0, k.size}}
	visited := make(map[int]bool)

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if head.ndx == k.size*k.size-1 {
			return head.cost
		}

		if _, ok := visited[head.ndx]; !ok {
			fmt.Println(x, y, head)
			visited[head.ndx] = true
			queue = append(queue, head.nextMoves(x, y)...)
		}
	}

	return -1
}
