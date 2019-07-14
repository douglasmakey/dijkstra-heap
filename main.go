package main

import (
	"fmt"
)

func main() {
	fmt.Println("Dijkstra")
	// Example
	graph := newGraph()
	graph.addEdge("S", "B", 4)
	graph.addEdge("S", "C", 2)
	graph.addEdge("B", "C", 1)
	graph.addEdge("B", "D", 5)
	graph.addEdge("C", "D", 8)
	graph.addEdge("C", "E", 10)
	graph.addEdge("D", "E", 2)
	graph.addEdge("D", "T", 6)
	graph.addEdge("E", "T", 2)
	fmt.Println(graph.getPath("S", "T"))
}
