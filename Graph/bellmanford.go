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

	for i := 0; i < len(dg.Nodes) - 1; i++ {
		for _, edge := range dg.Edges {
			source := edge.origin
			target := edge.destiny
			peso, _ := dg.GetWeightFromEdge(source, target)
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
		peso, _ := dg.GetWeightFromEdge(source, target)
		
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