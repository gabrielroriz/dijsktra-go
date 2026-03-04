package dijkstra

import (
	"math"
	"testing"
)

func TestRun_CalculatesShortestPathsAndPredecessors(t *testing.T) {
	a := &Vertice{Label: "A"}
	b := &Vertice{Label: "B"}
	c := &Vertice{Label: "C"}
	d := &Vertice{Label: "D"}
	e := &Vertice{Label: "E"}
	f := &Vertice{Label: "F"} // unreachable from A

	a.Edges = []Edge{
		{To: b, Weight: 4},
		{To: c, Weight: 2},
	}
	b.Edges = []Edge{
		{To: c, Weight: 5},
		{To: d, Weight: 10},
	}
	c.Edges = []Edge{
		{To: e, Weight: 3},
	}
	e.Edges = []Edge{
		{To: d, Weight: 4},
	}

	g := &Graph{
		Vertices: []*Vertice{a, b, c, d, e, f},
	}

	Run(g, a)

	if a.Distance != 0 {
		t.Fatalf("A distance = %v, want 0", a.Distance)
	}
	if b.Distance != 4 {
		t.Fatalf("B distance = %v, want 4", b.Distance)
	}
	if c.Distance != 2 {
		t.Fatalf("C distance = %v, want 2", c.Distance)
	}
	if e.Distance != 5 {
		t.Fatalf("E distance = %v, want 5", e.Distance)
	}
	if d.Distance != 9 {
		t.Fatalf("D distance = %v, want 9", d.Distance)
	}
	if !math.IsInf(f.Distance, 1) {
		t.Fatalf("F distance = %v, want +Inf", f.Distance)
	}

	if a.Predecessor != nil {
		t.Fatalf("A predecessor = %v, want nil", a.Predecessor)
	}
	if b.Predecessor != a {
		t.Fatalf("B predecessor = %v, want A", b.Predecessor)
	}
	if c.Predecessor != a {
		t.Fatalf("C predecessor = %v, want A", c.Predecessor)
	}
	if e.Predecessor != c {
		t.Fatalf("E predecessor = %v, want C", e.Predecessor)
	}
	if d.Predecessor != e {
		t.Fatalf("D predecessor = %v, want E", d.Predecessor)
	}
	if f.Predecessor != nil {
		t.Fatalf("F predecessor = %v, want nil", f.Predecessor)
	}
}
