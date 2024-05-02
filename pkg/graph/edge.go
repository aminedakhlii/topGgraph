package graph

// Edge represents a directed edge in a graph that connects two nodes.
type Edge struct {
	Start      *Node                  // Pointer to the starting node of the edge
	End        *Node                  // Pointer to the ending node of the edge
	Relation   string                 // Describes the type of relationship or connection
	Properties map[string]interface{} // Holds additional properties or metadata of the edge
}

// NewEdge creates and returns a new Edge connecting two nodes.
func NewEdge(start, end *Node, relation string) *Edge {
	return &Edge{
		Start:      start,
		End:        end,
		Relation:   relation,
		Properties: make(map[string]interface{}),
	}
}

// SetProperty adds or updates a property on an edge.
func (e *Edge) SetProperty(key string, value interface{}) {
	e.Properties[key] = value
}

// GetProperty retrieves the value of a property from an edge.
// Returns the value and a boolean indicating if the property exists.
func (e *Edge) GetProperty(key string) (interface{}, bool) {
	value, found := e.Properties[key]
	return value, found
}

// RemoveProperty removes a property from an edge by key.
func (e *Edge) RemoveProperty(key string) {
	delete(e.Properties, key)
}
