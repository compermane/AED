package Graph

import "testing"

type GraphMap map[string]map[string]float64
var NodeIDs []string
var GraphMaps []GraphMap

func init() {
	NodeIDs = []string{"A", "B", "C", "D", "E", "F"}
	GraphMaps = make([]GraphMap, 0)

	Graph1 := make(GraphMap)
	Graph1["A"] = map[string]float64{"B": 10, "D": 5}
	Graph1["B"] = map[string]float64{"C": 1, "D": 2}
	Graph1["C"] = map[string]float64{"E": 4}
	Graph1["D"] = map[string]float64{"B": 3, "C": 9, "E": 2}
	Graph1["E"] = map[string]float64{"A": 7, "C": 6}


	Graph2 := make(GraphMap)
	Graph2["A"] = map[string]float64{"B": 1, "C": 4}
	Graph2["B"] = map[string]float64{"C": 3, "D": 2, "E": 2}
	Graph2["C"] = map[string]float64{}
	Graph2["D"] = map[string]float64{"B": 1, "C": 5}
	Graph2["E"] = map[string]float64{"D": -3}

	GraphMaps = append(GraphMaps, Graph1, Graph2)
}

func assertNodeID(t *testing.T, id string, node *Node) {
	if node.ID != id {
		t.Errorf("Node.ID error: expected: %v got %v", id, node.ID)
	}
}

func assertEmptyGraph(t *testing.T, dg *DiGraph) {
	if len(dg.Nodes) != 0 && len(dg.Edges) != 0 {
		t.Errorf("Graph error: expected lenghts 0 and 0, got %v and %v", len(dg.Nodes), len(dg.Edges))
	}
}

func assertSameEdgeWeight(t *testing.T, dg *DiGraph, nodeA, nodeB *Node, weight float64) {
	w, ok := dg.GetWeightFromEdge(nodeA, nodeB)

	if w != weight {
		t.Errorf("Weight error: weight in graph is not equal to the maps weight; expected %v got %v from nodes %v and %v", weight, w, nodeA.ID, nodeB.ID)
	}
	if !ok {
		t.Errorf("Edge error: edge does not exist in graph from nodes %v and %v", nodeA.ID, nodeB.ID)
	}
}
func TestCreateNode(t *testing.T) {
	for _, id := range NodeIDs {
		node := CreateNode(id)
		assertNodeID(t, id, node)
	}
}

func TestCreateGraph(t *testing.T) {
	dg := CreateGraph()
	assertEmptyGraph(t, dg)
}

func TestCreateGraphFromMap(t *testing.T) {
	for _, graphMap := range GraphMaps {
		dg := CreateGraphByMap(graphMap)
		for origin, destinies := range graphMap {
			for destiny, weight := range destinies {
				org := dg.GetNodeFromID(origin)
				dty := dg.GetNodeFromID(destiny)
				assertSameEdgeWeight(t, dg, org, dty, weight)
			}
		}
	}
}