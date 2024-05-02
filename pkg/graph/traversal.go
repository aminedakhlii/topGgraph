package graph

// DFS (Depth-First Search) recursively explores nodes, going as deep as possible along each branch before backtracking.
func (g *Graph) DFS(startID string, visit func(*Node)) {
	visited := make(map[string]bool)
	g.dfsHelper(startID, visit, visited)
}

// Helper function for DFS to handle recursion.
func (g *Graph) dfsHelper(nodeID string, visit func(*Node), visited map[string]bool) {
	node, exists := g.Nodes[nodeID]
	if !exists {
		return
	}
	if visited[nodeID] {
		return
	}
	visited[nodeID] = true
	visit(node) // Process the current node

	for _, edge := range node.Edges {
		g.dfsHelper(edge.End.ID, visit, visited)
	}
}

// BFS (Breadth-First Search) explores the neighbor nodes at the present depth prior to moving on to nodes at the next depth level.
func (g *Graph) BFS(startID string, visit func(*Node)) {
	queue := []*Node{}
	visited := make(map[string]bool)

	startNode, exists := g.Nodes[startID]
	if !exists {
		return
	}

	queue = append(queue, startNode) // Enqueue start node
	visited[startID] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] // Dequeue
		visit(node)       // Process the node

		for _, edge := range node.Edges {
			if !visited[edge.End.ID] {
				visited[edge.End.ID] = true
				queue = append(queue, edge.End)
			}
		}
	}
}
