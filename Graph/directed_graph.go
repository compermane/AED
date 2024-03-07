package Graph

import "fmt"

/*
	Há várias formas de representar um grafo direcionado em código, como matrizes
	ou listas adjacência. Aqui, usaremos a seguinte notação:

    vertices = {
        v1: {v2: distancia_de_v1_ate_v2, v3: distancia_de_v1_ate_v3},
        v2: {v3: distancia_de_v2_ate_v3},
        ...
    }

	Ou seja, um grafo é um conjunto de nós, cada nó armazenando um conjunto de
	distâncias para vários outros nós.
*/

const INF float64 = float64(^uint(0) >> 1)

type DiGraph struct {
	Nodes 	[]*Node
	Edges 	[]*Edge
}

type Node struct {
	ID		string
}

type Edge struct {
	origin 	*Node
	destiny	*Node
	weight	float64
}

func CreateNode(id string) *Node {
	return &Node {
		ID: id,
	}
}

func CreateEdge(nodeA, nodeB *Node, weight float64) *Edge {
	return &Edge {
		origin: nodeA,
		destiny: nodeB,
		weight: weight,
	}
}

func CreateGraph() *DiGraph {
	nodes := make([]*Node, 0)
	edges := make([]*Edge, 0)

	return &DiGraph {
		Nodes: nodes,
		Edges: edges,
	}
}

func (dg *DiGraph) AddNode(newNode *Node) {
	dg.Nodes = append(dg.Nodes, newNode)
}

func (dg *DiGraph) AddEdge(nodeA, nodeB *Node, weight float64) {
	newEdge := CreateEdge(nodeA, nodeB, weight)
	dg.Edges = append(dg.Edges, newEdge)
}

func CreateGraphByMap(nodesMap map[string]map[string]float64) *DiGraph {
	dg := CreateGraph()

	for id := range nodesMap {
		dg.AddNode(CreateNode(id))
	}

	for origin, neighbors := range nodesMap {
		org := dg.GetNodeFromID(origin)
		for neighborID, distance := range neighbors {
			dty := dg.GetNodeFromID(neighborID)
			dg.AddEdge(org, dty, distance)
		}
	}
	return dg
}

func (dg *DiGraph) GetNodeFromID(nodeID string) *Node {
	for _, node := range dg.Nodes {
		if node.ID == nodeID {
			return node
		}
	}

	return nil
}

func (dg *DiGraph) PrintDiGraph() {
	for _, node := range dg.Nodes {
		fmt.Printf("%v: {", node.ID)
		var currentEdges []*Edge

		for _, edge := range dg.Edges {
			if edge.origin == node {
				currentEdges = append(currentEdges, edge)
			}
		}

		n := len(currentEdges) - 1
		i := 0
		for _, edge := range currentEdges {
			fmt.Printf("%v: %v", edge.destiny.ID, edge.weight)
			if i < n {
				fmt.Printf(", ")
			}
			i++
		}
		fmt.Printf("}\n")
	}
}

func (dg *DiGraph) GetWeightFromEdge(nodeA, nodeB *Node) (float64, bool) {
	for _, edge := range dg.Edges {
		if edge.origin == nodeA && edge.destiny == nodeB {
			return edge.weight, true
		}
	}

	return 0, false
}

func (dg *DiGraph) Neighbors(origin *Node) []*Node {
	neighbors := make([]*Node, 0)

	for _, edge := range dg.Edges {
		if edge.origin == origin {
			neighbors = append(neighbors, edge.destiny)
		}
	}

	return neighbors
}

func popNode(delNode *Node, arr []*Node) []*Node {
	newArr := make([]*Node, 0)

	for _, node := range arr {
		if node.ID != delNode.ID {
			newArr = append(newArr, node)
		}
	}

	return newArr
}

/*
Função para verificar a existência de um node em um slice de nodes. Útil para BFS
:param node: node a ser verificado
:param nodes: slice de nodes
:returns: bool; true se existe e false caso contrário 
*/
func checkForNodes(node *Node, nodes []*Node) bool {
	for _, currentNode := range nodes {
		if currentNode == node {
			return true
		}
	}

	return false
}