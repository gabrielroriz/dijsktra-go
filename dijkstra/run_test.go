package dijkstra_test

import (
	"math"
	"testing"

	"github.com/gabrielroriz/dijsktra-go/dijkstra"
)

func TestRunShortestPaths(t *testing.T) {
	a := &dijkstra.Vertice{Label: "A"}
	b := &dijkstra.Vertice{Label: "B"}
	c := &dijkstra.Vertice{Label: "C"}
	d := &dijkstra.Vertice{Label: "D"}
	unreachable := &dijkstra.Vertice{Label: "unreachable"}
	a.Edges = []dijkstra.Edge{{To: b, Weight: 10}, {To: c, Weight: 3}}
	c.Edges = []dijkstra.Edge{{To: b, Weight: 2}, {To: d, Weight: 8}}
	b.Edges = []dijkstra.Edge{{To: d, Weight: 0}}
	d.Edges = []dijkstra.Edge{{To: b, Weight: 0}, {To: a, Weight: 7}}
	graph := &dijkstra.Graph{Vertices: []*dijkstra.Vertice{unreachable, d, b, a, c}}

	dijkstra.Run(graph, a)

	for _, tt := range []struct {
		vertex      *dijkstra.Vertice
		distance    float64
		predecessor *dijkstra.Vertice
	}{
		{a, 0, nil},
		{b, 5, c},
		{c, 3, a},
		{d, 5, b},
		{unreachable, math.Inf(1), nil},
	} {
		t.Run(tt.vertex.Label, func(t *testing.T) {
			if tt.vertex.Distance != tt.distance {
				t.Errorf("distance = %g, want %g", tt.vertex.Distance, tt.distance)
			}
			if tt.vertex.Predecessor != tt.predecessor {
				t.Errorf("predecessor = %p, want %p", tt.vertex.Predecessor, tt.predecessor)
			}
		})
	}

	// A second origin must clear both distances and predecessor chains.
	dijkstra.Run(graph, unreachable)
	for _, vertex := range graph.Vertices {
		want := math.Inf(1)
		if vertex == unreachable {
			want = 0
		}
		if vertex.Distance != want || vertex.Predecessor != nil {
			t.Errorf("%s after rerun: distance = %g, predecessor = %p; want %g, nil",
				vertex.Label, vertex.Distance, vertex.Predecessor, want)
		}
	}
}

func TestRunEmptyGraph(t *testing.T) {
	dijkstra.Run(&dijkstra.Graph{}, nil)
}

func TestRunSingleVertex(t *testing.T) {
	vertex := &dijkstra.Vertice{Label: "A"}
	vertex.Edges = []dijkstra.Edge{{To: vertex, Weight: 0}}
	dijkstra.Run(&dijkstra.Graph{Vertices: []*dijkstra.Vertice{vertex}}, vertex)
	if vertex.Distance != 0 || vertex.Predecessor != nil {
		t.Fatalf("distance = %g, predecessor = %p; want 0, nil", vertex.Distance, vertex.Predecessor)
	}
}
