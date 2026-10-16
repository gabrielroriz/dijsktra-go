# Algorithms in Go

Algorithm implementations in Go, starting with Dijkstra for finding shortest paths in a graph. No external dependencies.

Requires **Go 1.27.2** or later.

```sh
go run ./cmd/dijkstra
go test ./...
go vet ./...
```

The example prints distances from A: `A: 0`, `B: 4`, and `C: 6`.

## Dijkstra

Import `github.com/gabrielroriz/dijsktra-go/dijkstra`, build a `Graph` with vertices (`Vertice`) and edges (`Edge`), and call `dijkstra.Run(graph, origin)`. Read `Distance` and `Predecessor` on each vertex to get the results.

The graph is directed; for a connection in both directions, add an edge in each direction. Weights must be finite and nonnegative, and the origin and all edge destinations must belong to the graph. Unreachable vertices have a distance of `+Inf` and a `nil` predecessor. Each run replaces the previous results. Complexity: O(V² + E) time and O(V) extra space.

## Adding an algorithm

Create a package in its own directory at the repository root, such as `bfs/`, with the implementation and `*_test.go` tests. Executable examples belong in `cmd/<algorithm>/`. Use the same Go module and add a brief description and the example command here.

Before committing, run `gofmt -w .`, `go test ./...`, and `go vet ./...`.
