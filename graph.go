package main

type edge struct {
	node   string
	weight int
}

type graph struct {
	nodes map[string][]edge
}

func newGraph() *graph {
	return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) addEdge(origin, destiny string, weight int) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
	g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
}

func (g *graph) getEdges(node string) []edge {
	return g.nodes[node]
}

func (g *graph) getPath(origin, destiny string) (int, []string) {
	h := newHeap()

	for _, e := range g.getEdges(origin) {
		h.push(path{value: e.weight, nodes: []string{origin, e.node}})
	}

	visited := []string{}
	visited = append(visited, origin)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if wasVisited(node, visited) {
			continue
		}

		if node == destiny {
			return p.value, p.nodes
		}

		for _, e := range g.getEdges(node) {
			if !wasVisited(e.node, visited) {
				// We calculate the total spent so far plus the cost and the path of getting here
				h.push(path{value: p.value + e.weight, nodes: append(p.nodes, e.node)})
			}
		}

		visited = append(visited, node)
	}

	return 0, nil
}
