package graph

import (
	"fmt"
)

// Node represents a graph node with an ID and a set of properties.
type Node struct {
	ID         string                 `json:"id"`
	Properties map[string]interface{} `json:"properties"`
	Edges      []*Edge                `json:"edges"`
}

type SerializableNode struct {
	ID         string                 `json:"id"`
	Properties map[string]interface{} `json:"properties"`
	Edges      []string               `json:"edges"` // Store only IDs of connected nodes
}

// NewNode creates and returns a new Node with the given ID.
func NewNode(id string) *Node {
	return &Node{
		ID:         id,
		Properties: make(map[string]interface{}),
		Edges:      []*Edge{},
	}
}

// AddProperty adds or updates a property to the node.
func (n *Node) AddProperty(key string, value interface{}) {
	n.Properties[key] = value
}

// GetProperty retrieves a property by key from the node.
// Returns the property value and a boolean indicating if the key was found.
func (n *Node) GetProperty(key string) (interface{}, bool) {
	value, found := n.Properties[key]
	return value, found
}

// RemoveProperty removes a property from the node by key.
func (n *Node) RemoveProperty(key string) {
	delete(n.Properties, key)
}

// AddEdge adds an edge to this node's list of edges.
func (n *Node) AddEdge(edge *Edge) {
	n.Edges = append(n.Edges, edge)
}

// RemoveEdge attempts to remove an edge from the node and returns a boolean indicating success.
func (n *Node) RemoveEdge(toNodeID string) bool {
	for i, edge := range n.Edges {
		if edge.End.ID == toNodeID {
			// Remove the edge by slicing it out of the array
			n.Edges = append(n.Edges[:i], n.Edges[i+1:]...)
			return true
		}
	}
	return false
}

// ListEdges prints all edges connected to the node.
func (n *Node) ListEdges() {
	fmt.Println("Edges connected to Node:", n.ID)
	for _, edge := range n.Edges {
		fmt.Printf("  - %s to %s (%s)\n", edge.Start.ID, edge.End.ID, edge.Relation)
	}
}

func (n *Node) ToSerializable() SerializableNode {
	edges := make([]string, len(n.Edges))
	for i, edge := range n.Edges {
		edges[i] = edge.End.ID // Assuming End is the node it points to
	}
	return SerializableNode{
		ID:         n.ID,
		Properties: n.Properties,
		Edges:      edges,
	}
}
