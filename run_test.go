package dijkstra

import (
	"math"
	"testing"
)

func TestRunComputesShortestPaths(t *testing.T) {
	a := &Vertice{Label: "A"}
	b := &Vertice{Label: "B"}
	c := &Vertice{Label: "C"}
	d := &Vertice{Label: "D"}
	e := &Vertice{Label: "E"}

	a.Edges = []Edge{
		{Weight: 4, To: b},
		{Weight: 2, To: c},
	}
	b.Edges = []Edge{
		{Weight: 1, To: d},
	}
	c.Edges = []Edge{
		{Weight: 1, To: b},
		{Weight: 5, To: d},
	}

	graph := &Graph{
		Vertices: []*Vertice{a, b, c, d, e},
	}

	Run(graph, a)

	if a.Distance != 0 {
		t.Fatalf("distance from A to A should be 0, got %v", a.Distance)
	}
	if c.Distance != 2 {
		t.Fatalf("distance from A to C should be 2, got %v", c.Distance)
	}
	if b.Distance != 3 {
		t.Fatalf("distance from A to B should be 3, got %v", b.Distance)
	}
	if d.Distance != 4 {
		t.Fatalf("distance from A to D should be 4, got %v", d.Distance)
	}
	if !math.IsInf(e.Distance, 1) {
		t.Fatalf("distance from A to unreachable E should be +Inf, got %v", e.Distance)
	}

	if c.Predecessor != a {
		t.Fatalf("predecessor of C should be A, got %v", c.Predecessor)
	}
	if b.Predecessor != c {
		t.Fatalf("predecessor of B should be C, got %v", b.Predecessor)
	}
	if d.Predecessor != b {
		t.Fatalf("predecessor of D should be B, got %v", d.Predecessor)
	}
	if e.Predecessor != nil {
		t.Fatalf("predecessor of unreachable E should be nil, got %v", e.Predecessor)
	}
}
