package Graph

import (
	"testing"
)

type GraphMap map[string]map[string]float64
var NodeIDs []string
var BFSExpected, DFSExpected [][]string
var GraphMaps, BFSMaps []GraphMap

func init() {
	NodeIDs = []string{"A", "B", "C", "D", "E", "F"}
	GraphMaps = make([]GraphMap, 0)
	BFSMaps = make([]GraphMap, 0)
	BFSExpected = make([][]string, 0)
	DFSExpected = make([][]string, 0)

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

	// TODO: ADICIONAR TESTES PARA TRANSVERSAL QUE FUNCIONEM!!
	Graph3 := make(GraphMap)
	Graph3["0"] = map[string]float64{"1": 1, "2": 1, "3": 1}
	Graph3["1"] = map[string]float64{"0": 1, "2": 1}
	Graph3["2"] = map[string]float64{"0": 1, "4": 1}
	Graph3["3"] = map[string]float64{"0": 1}
	Graph3["4"] = map[string]float64{"2": 1}

	Graph4 := make(GraphMap)
	Graph4["0"] = map[string]float64{"1": 1, "3": 1}
	Graph4["1"] = map[string]float64{"0": 1, "2": 1, "3": 1, "5": 1, "6": 1}
	Graph4["2"] = map[string]float64{"1": 1, "3": 1, "4": 1, "5": 1}
	Graph4["3"] = map[string]float64{"0": 1, "1": 1, "2": 1, "4": 1}
	Graph4["4"] = map[string]float64{"2": 1, "3": 1, "6": 1}
	Graph4["5"] = map[string]float64{"1": 1, "2": 1}
	Graph4["6"] = map[string]float64{"1": 1, "4": 1}

	GraphMaps = append(GraphMaps, Graph1, Graph2)
	BFSMaps = append(BFSMaps, Graph3)
	BFSExpected = append(BFSExpected, []string{"0", "1", "2", "3", "4"}, []string{"0", "1", "2", "3", "4", "5", "6"})
	DFSExpected = append(DFSExpected, []string{"0", "1", "2", "4", "3"})
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

func assertTransversal(t *testing.T, got, expected []string, typeT string) {
	var mode string
	if typeT == "dfs" {
		mode = "DFS"
	} else if typeT == "bfs" {
		mode = "BFS"
	} else {
		mode = "UNKNOWN"
	}

	if len(got) != len(expected) {
		t.Errorf(" (%v) Search fail: got transversal of lenght %v expected %v\n", mode, len(got), len(expected))
	}

	for i := range got {
		if got[i] != expected[i]  {
			t.Errorf(" (%v) Value error: in position %v got %v expected %v\n", mode, i, got[i], expected[i])
		}
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

func TestBFS(t *testing.T) {
	for i, graph := range BFSMaps {
		dg := CreateGraphByMap(graph)
		visited := dg.BFS("0")
		expected := BFSExpected[i]

		assertTransversal(t, visited, expected, "bfs")
	}

}

func TestDFS(t *testing.T) {
	for i, graph := range BFSMaps {
		dg := CreateGraphByMap(graph)
		visited := dg.DFS("0")
		expected := DFSExpected[i]

		assertTransversal(t, visited, expected, "dfs")
	}
	
}