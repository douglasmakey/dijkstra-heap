package main

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	t.Run("one", func(tt *testing.T) {
		graph := newGraph()
		graph.addEdge("S", "P", 2)
		graph.addEdge("S", "U", 3)
		graph.addEdge("P", "Q", 5)
		graph.addEdge("P", "X", 4)
		graph.addEdge("U", "X", 1)
		graph.addEdge("U", "V", 3)
		graph.addEdge("X", "Q", 7)
		graph.addEdge("X", "Y", 6)
		graph.addEdge("X", "V", 8)
		graph.addEdge("V", "W", 4)
		graph.addEdge("Y", "R", 1)
		graph.addEdge("Y", "W", 3)
		graph.addEdge("Q", "R", 2)
		graph.addEdge("R", "T", 6)
		graph.addEdge("W", "T", 5)
		cost, path := graph.getPath("S", "T")
		if cost != 15 {
			t.Errorf("expected 15 got %d", cost)
		}

		pathExpected := []string{"S", "P", "Q", "R", "T"}
		if !reflect.DeepEqual(path, pathExpected) {
			t.Errorf("expected %v got %v", pathExpected, path)
		}
	})

	t.Run("two", func(tt *testing.T) {
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
		cost, path := graph.getPath("S", "T")
		if cost != 12 {
			t.Errorf("expected 12 got %d", cost)
		}

		pathExpected := []string{"S", "C", "B", "D", "E", "T"}
		if !reflect.DeepEqual(path, pathExpected) {
			t.Errorf("expected %v got %v", pathExpected, path)
		}
	})
}
