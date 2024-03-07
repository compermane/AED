package Graph

func (dg *DiGraph) BFS(originID string) []string {
	originNode := dg.GetNodeFromID(originID)

	visited := make([]*Node, 0)
	visitedStr := make([]string, 0)
	fila := make([]*Node, 0)

	visited = append(visited, originNode)
	visitedStr = append(visitedStr, originID)
	fila = append(fila, dg.Neighbors(originNode)...)

	for {
		atual := fila[0]

		for checkForNodes(atual, visited) && len(fila) != 0{
			fila = fila[1:]
			if len(fila) == 0 {
				atual = nil
				break
			}
			atual = fila[0]
		}

		if len(fila) == 0 || atual == nil {
			break
		} else {		
			fila = append(fila, dg.Neighbors(atual)...)
			fila = fila[1:]

			visited = append(visited, atual)
			visitedStr = append(visitedStr, atual.ID)
	}
	}

	return visitedStr
}