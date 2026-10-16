# Algoritmos em Go

Implementações de algoritmos em Go, começando por Dijkstra para encontrar os menores caminhos em um grafo. Sem dependências externas.

Requer **Go 1.27.2** ou superior.

```sh
go run ./cmd/dijkstra
go test ./...
go vet ./...
```

O exemplo imprime as distâncias a partir de A: `A: 0`, `B: 4` e `C: 6`.

## Dijkstra

Importe `github.com/gabrielroriz/dijsktra-go/dijkstra`, monte um `Graph` com seus vértices (`Vertice`) e arestas (`Edge`) e chame `dijkstra.Run(graph, origem)`. Consulte `Distance` e `Predecessor` em cada vértice para obter o resultado.

O grafo é direcionado; para uma conexão nos dois sentidos, adicione uma aresta em cada direção. Os pesos devem ser finitos e não negativos, e a origem e os destinos das arestas devem pertencer ao grafo. Vértices inalcançáveis ficam com distância `+Inf` e predecessor `nil`. Cada execução substitui os resultados anteriores. Complexidade: O(V² + E) em tempo e O(V) em memória adicional.

## Adicionar um algoritmo

Crie um pacote em um diretório próprio na raiz, como `bfs/`, com a implementação e testes `*_test.go`. Exemplos executáveis ficam em `cmd/<algoritmo>/`. Use o mesmo módulo Go e adicione aqui uma breve descrição e o comando do exemplo.

Antes de commitar, execute `gofmt -w .`, `go test ./...` e `go vet ./...`.
