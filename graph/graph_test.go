package graph

import "testing"

func TestGraph(t *testing.T) {
	keys := []int{
		10, 20, 30, 40, 50, 60, 70, 80, 90, 100,
	}

	edges := [][]int{
		{10, 20}, {20, 30}, {30, 40}, {40, 50}, {50, 60}, {60, 70}, {70, 80}, {80, 90}, {90, 100}, {100, 10},
	}

	graph := NewGraph()

	for _, k := range keys {
		_, err := graph.AddVertex(k)
		if err != nil {
			t.Errorf("error adding vertex: %v", err)
		}
	}

	for _, k := range keys {
		_, err := graph.GetVertex(k)
		if err != nil {
			t.Errorf("error getting vertex: %v", err)
		}
	}

	_, err := graph.GetVertex(200)
	if err == nil {
		t.Errorf("expected to return error, but got no error")
	}

	for _, adjKeys := range edges {
		if err := graph.AddEdge(adjKeys[0], adjKeys[1]); err != nil {
			t.Errorf("error adding edge: %v", err)
		}
	}

	if err := graph.AddEdge(200, 500); err == nil {
		t.Errorf("error adding edge between non-existing verticies [%d,%d]", 200, 500)
	}

	for _, adjKeys := range edges {
		if err := graph.RemoveEdge(adjKeys[0], adjKeys[1]); err != nil {
			t.Errorf("error removing edge: %v", err)
		}
	}

	for _, k := range keys {
		if err := graph.DeleteVertex(k); err != nil {
			t.Errorf("error deleting vertex: %v", err)
		}
	}
}
