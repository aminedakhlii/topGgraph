package main

// import (
// 	"fmt"

// 	"github.com/aminedakhlii/TopGGraph/pkg/store"
// )

// func main() {
// 	// Open BadgerDB where the graph is stored
// 	db, err := store.OpenBadgerDB("./tmp/badger")
// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to open BadgerDB: %v", err))
// 	}
// 	defer store.CloseBadgerDB(db)

// 	store.DumpDatabase(db)

// 	// Load the entire graph from the database
// 	loadedGraph, err := store.LoadGraph(db)
// 	if err != nil {
// 		fmt.Println("Failed to load graph:", err)
// 		return
// 	}

// 	// Perform a search on the loaded graph
// 	searchResults, err := store.SearchGraph(loadedGraph, "name", "Alice")
// 	if err != nil {
// 		fmt.Println("Search failed:", err)
// 		return
// 	}

// 	fmt.Println("Search Results:")
// 	for _, node := range searchResults {
// 		fmt.Printf("Found Node: ID=%s, Name=%s\n", node.ID, node.Properties["name"])
// 	}
// }
