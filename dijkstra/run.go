package dijkstra

import "math"

func initialize(g *Graph, origin *Vertice) {
	for _, vertex := range g.Vertices {
		vertex.Distance = math.Inf(1)
		vertex.Predecessor = nil
		if vertex == origin {
			vertex.Distance = 0
		}
	}
}

func relax(from *Vertice, edge Edge) {
	if distance := from.Distance + edge.Weight; distance < edge.To.Distance {
		edge.To.Distance = distance
		edge.To.Predecessor = from
	}
}

func findNext(g *Graph, processed map[*Vertice]bool) *Vertice {
	var lowest *Vertice
	for _, vertex := range g.Vertices {
		if processed[vertex] {
			continue
		}
		if lowest == nil || vertex.Distance < lowest.Distance {
			lowest = vertex
		}
	}
	return lowest
}

// Run computes shortest paths from origin, updating each vertex in g.
// Unreachable vertices retain a distance of +Inf and a nil predecessor.
// The graph must contain origin and all edge endpoints, with no nil vertices.
// Edge weights must be finite and nonnegative. An empty graph may use a nil origin.
// Each call resets the previous results. Run takes O(V² + E) time and O(V) extra space.
func Run(g *Graph, origin *Vertice) {
	initialize(g, origin)
	processed := make(map[*Vertice]bool, len(g.Vertices))

	for node := origin; node != nil; node = findNext(g, processed) {
		if math.IsInf(node.Distance, 1) {
			break
		}
		for _, edge := range node.Edges {
			relax(node, edge)
		}
		processed[node] = true
	}
}
