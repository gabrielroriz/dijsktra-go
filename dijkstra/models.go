// Package dijkstra finds shortest paths in directed graphs with nonnegative weights.
package dijkstra

// Vertice is a graph vertex and stores the results of the most recent Run.
type Vertice struct {
	Label string
	Edges []Edge

	Distance    float64
	Predecessor *Vertice
}

// Edge is a weighted, directed connection to another vertex.
type Edge struct {
	Weight float64
	To     *Vertice
}

// Graph contains all vertices, including those without outgoing edges.
type Graph struct {
	Vertices []*Vertice
}
