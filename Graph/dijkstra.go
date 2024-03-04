package Graph

func (dg *DiGraph) Dijkstra(origin, destiny *Node) float64 {
	distances := make(map[*Node]float64)
	predecessors := make(map[*Node]*Node) 
	notVisited := make([]*Node, 0)
	for _, node := range dg.Nodes {
		distances[node] = INF
		predecessors[node] = nil
		notVisited = append(notVisited, node)
	}
	distances[origin] = 0

	for len(notVisited) != 0 {
		minDistance := INF
		var minNode *Node
		for _, node := range notVisited {
			dist := distances[node]
			if dist < minDistance {
				minDistance = dist
				minNode = node
			}
		}

		atual := minNode
		neighbors := dg.Neighbors(atual)
		notVisited = popNode(atual, notVisited)

		for _, neighbor := range neighbors {
			peso, _ := dg.GetWeightFromEdge(atual, neighbor)
			pesoPotencial := distances[atual] + peso

			if pesoPotencial < distances[neighbor] {
				distances[neighbor] = pesoPotencial
				predecessors[neighbor] = atual
			}
		}
	}

	return distances[destiny]
}