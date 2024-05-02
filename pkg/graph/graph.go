package graph

import "fmt"

// Graph represents a set of nodes and edges forming a graph.
type Graph struct {
	Nodes map[string]*Node
}

// NewGraph initializes and returns a new instance of a Graph.
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

// ReconstructGraph builds a graph from a set of serializable Nodes
func ReconstructGraph(nodes map[string]*Node, serializableNodes map[string]*SerializableNode) {
	for id, sNode := range serializableNodes {
		node := nodes[id]
		for _, edgeID := range sNode.Edges {
			if targetNode, exists := nodes[edgeID]; exists {
				node.Edges = append(node.Edges, &Edge{Start: node, End: targetNode})
			}
		}
	}
}

// AddNode adds a new node to the graph with the given properties.
func (g *Graph) AddNode(id string, properties map[string]interface{}) {
	if _, exists := g.Nodes[id]; !exists {
		g.Nodes[id] = &Node{
			ID:         id,
			Properties: properties,
			Edges:      []*Edge{},
		}
	}
}

// AddEdge creates an edge between two nodes with the specified relation and properties.
func (g *Graph) AddEdge(startID, endID, relation string, properties map[string]interface{}) error {
	startNode, ok := g.Nodes[startID]
	if !ok {
		return fmt.Errorf("start node %s does not exist", startID)
	}

	endNode, ok := g.Nodes[endID]
	if !ok {
		return fmt.Errorf("end node %s does not exist", endID)
	}

	edge := &Edge{
		Start:      startNode,
		End:        endNode,
		Relation:   relation,
		Properties: properties,
	}

	startNode.Edges = append(startNode.Edges, edge)
	return nil
}

// FindNode returns a node by its ID if it exists in the graph.
func (g *Graph) FindNode(id string) (*Node, error) {
	if node, exists := g.Nodes[id]; exists {
		return node, nil
	}
	return nil, fmt.Errorf("node %s not found", id)
}

// RemoveNode deletes a node and its corresponding edges from the graph.
func (g *Graph) RemoveNode(id string) error {
	if node, exists := g.Nodes[id]; exists {
		// Remove all edges to and from the node
		for _, edge := range node.Edges {
			g.removeEdge(edge.Start.ID, edge.End.ID)
		}
		// Remove node from graph
		delete(g.Nodes, id)
		return nil
	}
	return fmt.Errorf("node %s not found", id)
}

// removeEdge is an unexported helper function to remove edges from a node.
func (g *Graph) removeEdge(startID, endID string) error {
	startNode, ok := g.Nodes[startID]
	if !ok {
		return fmt.Errorf("start node %s does not exist", startID)
	}

	newEdges := []*Edge{}
	for _, edge := range startNode.Edges {
		if edge.End.ID != endID {
			newEdges = append(newEdges, edge)
		}
	}
	startNode.Edges = newEdges
	return nil
}
