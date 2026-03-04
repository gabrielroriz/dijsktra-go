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
		{Weight: 10, To: b},
		{Weight: 1, To: c},
	}
	b.Edges = []Edge{
		{Weight: 1, To: d},
	}
	c.Edges = []Edge{
		{Weight: 2, To: b},
		{Weight: 10, To: d},
	}

	g := &Graph{Vertices: []*Vertice{a, b, c, d, e}}

	Run(g, a)

	if a.Distance != 0 {
		t.Fatalf("distance from A to A: got %v, want 0", a.Distance)
	}

	if c.Distance != 1 {
		t.Fatalf("distance from A to C: got %v, want 1", c.Distance)
	}

	if b.Distance != 3 {
		t.Fatalf("distance from A to B: got %v, want 3", b.Distance)
	}

	if d.Distance != 4 {
		t.Fatalf("distance from A to D: got %v, want 4", d.Distance)
	}

	if !math.IsInf(e.Distance, 1) {
		t.Fatalf("distance from A to E: got %v, want +Inf", e.Distance)
	}

	if c.Predecessor != a {
		t.Fatalf("predecessor of C: got %v, want %v", c.Predecessor, a)
	}

	if b.Predecessor != c {
		t.Fatalf("predecessor of B: got %v, want %v", b.Predecessor, c)
	}

	if d.Predecessor != b {
		t.Fatalf("predecessor of D: got %v, want %v", d.Predecessor, b)
	}

	if e.Predecessor != nil {
		t.Fatalf("predecessor of E: got %v, want nil", e.Predecessor)
	}
}
