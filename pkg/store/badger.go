package store

import (
	"encoding/json"
	"fmt"

	"github.com/aminedakhlii/TopGgraph/pkg/graph"
	"github.com/dgraph-io/badger/v3"
)

// OpenBadgerDB initializes and opens a connection to BadgerDB at a specified directory.
func OpenBadgerDB(dir string) (*badger.DB, error) {
	opts := badger.DefaultOptions(dir)
	opts.Logger = nil // Disable logging for cleaner output, set to a logger if needed
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CloseBadgerDB cleanly closes the BadgerDB connection.
func CloseBadgerDB(db *badger.DB) error {
	return db.Close()
}

func DumpDatabase(db *badger.DB) error {
	return db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			err := item.Value(func(val []byte) error {
				fmt.Printf("Key=%s, Value=%s\n", item.Key(), val)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// SaveNode stores a node in BadgerDB.
func SaveNode(db *badger.DB, nodeID string, nodeData []byte) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(nodeID), nodeData)
	})
}

// GetNode retrieves a node from BadgerDB.
func GetNode(db *badger.DB, nodeID string) ([]byte, error) {
	var nodeData []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(nodeID))
		if err != nil {
			return err
		}
		nodeData, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		return nil, err
	}
	return nodeData, nil
}

// MarshalNode converts a node struct to a byte slice.
func MarshalNode(node graph.Node) ([]byte, error) {
	fmt.Println(node)
	sn := node.ToSerializable()
	return json.Marshal(sn)
}

// UnmarshalNode converts a byte slice back to a node serializable struct.
func UnmarshalNode(data []byte) (*graph.SerializableNode, error) {
	var sn graph.SerializableNode
	err := json.Unmarshal(data, &sn)
	if err != nil {
		return nil, err
	}
	return &sn, nil
}

// LoadGraph loads all nodes from the database and reconstructs the graph.
func LoadGraph(db *badger.DB) (*graph.Graph, error) {
	g := &graph.Graph{Nodes: make(map[string]*graph.Node)}
	tempSerializableNodes := make(map[string]*graph.SerializableNode)

	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			nodeData, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}

			sNode, err := UnmarshalNode(nodeData)
			if err != nil {
				return err
			}

			// Create a node without edges
			node := &graph.Node{
				ID:         sNode.ID,
				Properties: sNode.Properties,
				Edges:      []*graph.Edge{},
			}

			g.Nodes[node.ID] = node
			tempSerializableNodes[node.ID] = sNode
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Now reconstruct the edges
	graph.ReconstructGraph(g.Nodes, tempSerializableNodes)

	return g, nil
}

// SearchGraph searches for nodes in the graph where the specified property matches the given value.
func SearchGraph(g *graph.Graph, propertyName, value string) ([]*graph.Node, error) {
	var results []*graph.Node
	for _, node := range g.Nodes {
		if propValue, ok := node.Properties[propertyName]; ok && propValue == value {
			results = append(results, node)
		}
	}
	return results, nil
}
