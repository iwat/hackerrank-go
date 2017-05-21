package bfs

type Graph struct {
	nodes []*Node
}

func NewGraph(nodes int) *Graph {
	graph := &Graph{make([]*Node, nodes)}
	for i := 0; i < nodes; i++ {
		graph.nodes[i] = NewNode(i + 1)
	}
	return graph
}

func (g *Graph) AddEdge(from, to int) {
	g.nodes[from-1].edges[g.nodes[to-1]] = 6
	g.nodes[to-1].edges[g.nodes[from-1]] = 6
}

func (g *Graph) DistancesFrom(start int) []int {
	var queue []*Node
	queue = append(queue, g.nodes[start-1])

	visited := make([]bool, len(g.nodes))
	distances := make([]int, len(g.nodes))
	visited[start-1] = true

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		for next, _ := range head.edges {
			if !visited[next.val-1] {
				visited[next.val-1] = true
				distances[next.val-1] = distances[head.val-1] + head.edges[next]
				queue = append(queue, next)
			}
		}
	}

	return distances
}

type Node struct {
	val   int
	edges map[*Node]int
}

func NewNode(val int) *Node {
	return &Node{val, make(map[*Node]int)}
}
