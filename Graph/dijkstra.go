package Graph

func (dg *DiGraph) Dijkstra(origin, destiny *Node) float64 {
	const INF float64 = float64(^uint(0) >> 1)

	distances := make(map[*Node]float64)
	for _, node := range dg.Nodes {
		distances[node] = INF
	}
	distances[origin] = 0

	predecessors := make(map[*Node]*Node) 
	for _, node := range dg.Nodes {
		predecessors[node] = nil
	}

	notVisited := make([]*Node, 0)
	for _, node := range dg.Nodes {
		notVisited = append(notVisited, node)
	}

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