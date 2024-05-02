# topGgraph

TopGGraph is an open-source graph management library developed in Go, designed to efficiently handle complex graph data structures. It's built to support various graph operations, making it a perfect toolkit for applications requiring interconnected data handling.

## Features

- **Persistent Graph Storage**: Implements efficient storage mechanisms for graphs.
- **Advanced Graph Traversal**: Supports multiple traversal algorithms including DFS and BFS.
- **High Performance**: Optimized for performance using Go's powerful concurrency features.
- **Versatile Graph Types**: Handles different types of graphs like directed, undirected, weighted, and unweighted.
- **Extensible**: Designed to be flexible and easy to extend with new features and algorithms.

## Installation

```bash
go get github.com/aminedakhlii/TopGgraph
```

## Creating a Graph
To start using TopGGraph, you first need to create a new graph instance:

```
package main

import (
    "github.com/yourusername/TopGGraph/pkg/graph"
)

func main() {
    g := graph.NewGraph()
}

```

## Adding Nodes and Edges
Once you have a graph instance, you can add nodes and edges to it:

```
// Add nodes
g.AddNode("1", map[string]interface{}{"name": "Alice"})
g.AddNode("2", map[string]interface{}{"name": "Bob"})

// Add an edge
g.AddEdge("1", "2", "knows", map[string]interface{}{"since": "2021"})

```

## Performing a BFS
To perform a BFS: 

```
g.BFS("1", func(n *graph.Node) {
  fmt.Println("Visited node:", n.ID)
})
```

## Saving the Graph
To persist your graph data, save it to BadgerDB:

```
import "github.com/yourusername/TopGGraph/pkg/store"

func main() {
    db, err := store.OpenBadgerDB("./tmp/badger")
    if err != nil {
        panic(err)
    }
    defer store.CloseBadgerDB(db)

    // Assume `node` is a *graph.Node you've already created and modified
    nodeData, _ := store.MarshalNode(node)
    if err := store.SaveNode(db, node.ID, nodeData); err != nil {
        panic(err)
    }
}
```

## Loading the Graph
To load an existing graph from BadgerDB:

```
func main() {
    db, err := store.OpenBadgerDB("./tmp/badger")
    if err != nil {
        panic(err)
    }
    defer store.CloseBadgerDB(db)

    loadedGraph, err := store.LoadGraph(db)
    if err != nil {
        panic(err)
    }

    // Use your graph
}
```

## License

TopGGraph is BSD 3-clause licensed, as found in the LICENSE file.

## Support

If you have any questions or issues, please open an issue on the GitHub repository.
