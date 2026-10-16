package main

import (
	"fmt"

	"github.com/gabrielroriz/dijsktra-go/dijkstra"
)

func main() {
	a := &dijkstra.Vertice{Label: "A"}
	b := &dijkstra.Vertice{Label: "B"}
	c := &dijkstra.Vertice{Label: "C"}
	a.Edges = []dijkstra.Edge{{To: b, Weight: 4}, {To: c, Weight: 10}}
	b.Edges = []dijkstra.Edge{{To: c, Weight: 2}}
	graph := &dijkstra.Graph{Vertices: []*dijkstra.Vertice{a, b, c}}

	dijkstra.Run(graph, a)
	for _, vertex := range graph.Vertices {
		fmt.Printf("%s: %g\n", vertex.Label, vertex.Distance)
	}
}
