package main

import (
	"fmt"

	"github.com/aminedakhlii/TopGgraph/pkg/graph"
	"github.com/aminedakhlii/TopGgraph/pkg/store"
)

func main() {
	// Assuming you have a function to open and close your database.
	db, err := store.OpenBadgerDB("./tmp/badger")
	if err != nil {
		panic(fmt.Sprintf("Failed to open BadgerDB: %v", err))
	}
	defer store.CloseBadgerDB(db)

	// Create a new graph instance
	g := graph.NewGraph()

	// Add nodes to the graph
	g.AddNode("1", map[string]interface{}{"name": "Alice"})
	g.AddNode("2", map[string]interface{}{"name": "Bob"})
	g.AddNode("3", map[string]interface{}{"name": "Charlie"})

	// Add edges between nodes
	g.AddEdge("1", "2", "knows", map[string]interface{}{"since": "2022"})
	g.AddEdge("2", "3", "knows", map[string]interface{}{"since": "2023"})
	g.AddEdge("3", "1", "knows", map[string]interface{}{"since": "2024"})

	// Save nodes to BadgerDB
	for _, node := range g.Nodes {
		nodeData, _ := store.MarshalNode(*node)
		if err := store.SaveNode(db, node.ID, nodeData); err != nil {
			fmt.Println("Error saving node:", err)
		}
	}

	// log the data from the database
	store.DumpDatabase(db)

	// Retrieve and print a node
	nodeData, err := store.GetNode(db, "2")
	if err != nil {
		fmt.Println("Error retrieving node:", err)
	} else {
		if node, err := store.UnmarshalNode(nodeData); err != nil {
			fmt.Println("Error unmarshaling node:", err)
		} else {
			fmt.Printf("Retrieved Node: %+v\n", node)
		}
	}

	// Perform a BFS on the graph starting from node 1
	fmt.Println("BFS starting from node 1:")
	g.BFS("1", func(n *graph.Node) {
		fmt.Println("Visited node:", n.ID)
	})
}
