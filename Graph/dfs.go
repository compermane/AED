package Graph

func (dg *DiGraph) DFS(originID string) []string {
	originNode := dg.GetNodeFromID(originID)

	if originNode == nil {
		return nil
	}

	visited := make([]*Node, 0)
	visited = append(visited, originNode)
	visitedStr := make([]string, 0)
	visitedStr = append(visitedStr, originID)

	stack := make([]*Node, 0)
	stack = append(stack, dg.Neighbors(originNode)...)
	atual := originNode

	for  {
		atual = stack[len(stack) - 1]

		for checkForNodes(atual, visited) && len(stack) > 0 {
			stack = stack[:len(stack) - 1]

			if len(stack) == 0 {
				atual = nil
			} else {
				atual = stack[len(stack) - 1]
			}
		}

		if atual == nil {
			break
		}

		visited = append(visited, atual)
		visitedStr = append(visitedStr, atual.ID)
		stack = append(stack, dg.Neighbors(atual)...)

	}

	return visitedStr
}