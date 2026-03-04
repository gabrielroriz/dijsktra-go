package dijkstra

import (
	"math"
	"testing"
)

func TestRunFindsShortestPaths(t *testing.T) {
	a := &Vertice{Label: "A"}
	b := &Vertice{Label: "B"}
	c := &Vertice{Label: "C"}
	d := &Vertice{Label: "D"}

	a.Edges = []Edge{
		{Weight: 1, To: b},
		{Weight: 4, To: c},
	}
	b.Edges = []Edge{
		{Weight: 2, To: c},
		{Weight: 6, To: d},
	}
	c.Edges = []Edge{
		{Weight: 3, To: d},
	}

	graph := &Graph{Vertices: []*Vertice{a, b, c, d}}
	Run(graph, a)

	if a.Distance != 0 {
		t.Fatalf("expected A distance 0, got %v", a.Distance)
	}
	if a.Predecessor != nil {
		t.Fatalf("expected A predecessor nil, got %v", a.Predecessor)
	}

	if b.Distance != 1 {
		t.Fatalf("expected B distance 1, got %v", b.Distance)
	}
	if b.Predecessor != a {
		t.Fatalf("expected B predecessor A, got %v", b.Predecessor)
	}

	if c.Distance != 3 {
		t.Fatalf("expected C distance 3, got %v", c.Distance)
	}
	if c.Predecessor != b {
		t.Fatalf("expected C predecessor B, got %v", c.Predecessor)
	}

	if d.Distance != 6 {
		t.Fatalf("expected D distance 6, got %v", d.Distance)
	}
	if d.Predecessor != c {
		t.Fatalf("expected D predecessor C, got %v", d.Predecessor)
	}
}

func TestRunKeepsUnreachableVertexAsInfinity(t *testing.T) {
	a := &Vertice{Label: "A"}
	b := &Vertice{Label: "B"}
	c := &Vertice{Label: "C"}

	a.Edges = []Edge{{Weight: 2, To: b}}

	graph := &Graph{Vertices: []*Vertice{a, b, c}}
	Run(graph, a)

	if b.Distance != 2 {
		t.Fatalf("expected B distance 2, got %v", b.Distance)
	}
	if b.Predecessor != a {
		t.Fatalf("expected B predecessor A, got %v", b.Predecessor)
	}

	if !math.IsInf(c.Distance, 1) {
		t.Fatalf("expected C distance +Inf, got %v", c.Distance)
	}
	if c.Predecessor != nil {
		t.Fatalf("expected C predecessor nil, got %v", c.Predecessor)
	}
}

func TestRunComputesShortestPathsAndPredecessors(t *testing.T) {
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
		t.Fatalf("distance for A = %v, want 0", a.Distance)
	}
	if c.Distance != 1 {
		t.Fatalf("distance for C = %v, want 1", c.Distance)
	}
	if b.Distance != 3 {
		t.Fatalf("distance for B = %v, want 3", b.Distance)
	}
	if d.Distance != 4 {
		t.Fatalf("distance for D = %v, want 4", d.Distance)
	}
	if !math.IsInf(e.Distance, 1) {
		t.Fatalf("distance for E = %v, want +Inf", e.Distance)
	}

	if c.Predecessor != a {
		t.Fatalf("predecessor for C = %v, want A", c.Predecessor)
	}
	if b.Predecessor != c {
		t.Fatalf("predecessor for B = %v, want C", b.Predecessor)
	}
	if d.Predecessor != b {
		t.Fatalf("predecessor for D = %v, want B", d.Predecessor)
	}
	if e.Predecessor != nil {
		t.Fatalf("predecessor for E = %v, want nil", e.Predecessor)
	}
}
