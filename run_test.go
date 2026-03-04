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
		t.Fatalf("distance from A to A = %v, want 0", a.Distance)
	}

	if c.Distance != 1 {
		t.Fatalf("distance from A to C = %v, want 1", c.Distance)
	}

	if b.Distance != 3 {
		t.Fatalf("distance from A to B = %v, want 3", b.Distance)
	}

	if d.Distance != 4 {
		t.Fatalf("distance from A to D = %v, want 4", d.Distance)
	}

	if !math.IsInf(e.Distance, 1) {
		t.Fatalf("distance from A to E = %v, want +Inf", e.Distance)
	}

	if c.Predecessor != a {
		t.Fatalf("predecessor of C = %v, want A", predecessorLabel(c.Predecessor))
	}

	if b.Predecessor != c {
		t.Fatalf("predecessor of B = %v, want C", predecessorLabel(b.Predecessor))
	}

	if d.Predecessor != b {
		t.Fatalf("predecessor of D = %v, want B", predecessorLabel(d.Predecessor))
	}

	if e.Predecessor != nil {
		t.Fatalf("predecessor of E = %v, want nil", predecessorLabel(e.Predecessor))
	}
}

func predecessorLabel(v *Vertice) string {
	if v == nil {
		return "<nil>"
	}
	return v.Label
}
