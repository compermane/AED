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

type DiGraph struct {
	Nodes map[string]*Node
}

type Node struct {
	NodeID 		string
	Distances 	map[string]int
}

func CreateNode(nodeID string, distances map[string]int) *Node {
	return &Node {
		NodeID: nodeID,
		Distances: distances,
	}
}

func CreateDiGraph(nodes map[string]*Node) *DiGraph {
	return &DiGraph {
		Nodes: nodes,
	}
}

func (dg *DiGraph) PrintDiGraph() {
	for _, node := range dg.Nodes {
		fmt.Printf("%v: {", node.NodeID)
		for dest, dist := range node.Distances {
			fmt.Printf("%v: %v, ", dest, dist)
		}
		fmt.Printf("}\n")
	}
}

func main() {
    // Create a DiGraph instance
    graph := DiGraph{
        Nodes: make(map[string]*Node),
    }

    // Add some nodes to the graph
    nodeA := &Node{
        NodeID: "A",
        Distances: map[string]int{
            "B": 10,
            "C": 20,
        },
    }
    graph.Nodes["A"] = nodeA

    nodeB := &Node{
        NodeID: "B",
        Distances: map[string]int{
            "A": 5,
        },
    }
    graph.Nodes["B"] = nodeB

	graph.PrintDiGraph()
}
