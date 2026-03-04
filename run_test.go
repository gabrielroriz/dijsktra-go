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
	e := &Vertice{Label: "E"} // unreachable from A

	a.Edges = []Edge{
		{Weight: 4, To: b},
		{Weight: 1, To: c},
	}
	b.Edges = []Edge{
		{Weight: 1, To: d},
	}
	c.Edges = []Edge{
		{Weight: 2, To: b},
		{Weight: 5, To: d},
	}

	graph := &Graph{
		Vertices: []*Vertice{a, b, c, d, e},
	}

	Run(graph, a)

	if a.Distance != 0 {
		t.Fatalf("distance(A) = %v, want 0", a.Distance)
	}

	if b.Distance != 3 {
		t.Fatalf("distance(B) = %v, want 3", b.Distance)
	}

	if c.Distance != 1 {
		t.Fatalf("distance(C) = %v, want 1", c.Distance)
	}

	if d.Distance != 4 {
		t.Fatalf("distance(D) = %v, want 4", d.Distance)
	}

	if !math.IsInf(e.Distance, 1) {
		t.Fatalf("distance(E) = %v, want +Inf for unreachable node", e.Distance)
	}

	if b.Predecessor != c {
		t.Fatalf("predecessor(B) = %v, want C", b.Predecessor)
	}

	if d.Predecessor != b {
		t.Fatalf("predecessor(D) = %v, want B", d.Predecessor)
	}
}
