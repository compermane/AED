package Graph

func (dg *DiGraph) BellmanFord(origin *Node, destiny *Node) (float64, bool) {
	distances := make(map[*Node]float64)
	predecessors := make(map[*Node]*Node)
	notVisited := make([]*Node, 0)
	for _, node := range dg.Nodes {
		distances[node] = INF
		predecessors[node] = nil
		notVisited = append(notVisited, node)
	}

	distances[origin] = 0

	i := 0
	for ; i < len(dg.Nodes) - 1; i++ {
		for _, edge := range dg.Edges {
			source := edge.origin
			target := edge.destiny
			peso := edge.weight
			pesoPotencial := peso + distances[source]

			if pesoPotencial < distances[target] {
				distances[target] = pesoPotencial
				predecessors[target] = source
			}
		}
	}

	for _, edge := range dg.Edges {
		source := edge.origin
		target := edge.destiny
		peso := edge.weight
		
		// Contém ciclo com peso negativo
		if distances[source] + peso < distances[target] {
			return -INF, false
		}
	}

	if distances[destiny] == INF {
		return INF, false
	} else {
		return distances[destiny], true
	}
}