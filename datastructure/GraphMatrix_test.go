package TestProject

import "testing"

func InitUnDirectedGraph() *GraphMatrix {
	graph := InitUndirectedGraphMatrix(4, "V0", "V1", "V2", "V3")
	return graph
}

func InitDirectedGraph() *GraphMatrix {
	graph := InitDirectedGraphMatrix(4, "V0", "V1", "V2", "V3")
	return graph
}
func TestCreateGraphMatrix(t *testing.T) {
	graph := InitUnDirectedGraph()
	t.Log(graph)
}

func TestCreateGraphMatrixDirected(t *testing.T) {
	graph := InitDirectedGraph()
	t.Log(graph)
}

func TestAddElementInGraphMatrix(t *testing.T) {
	graph := InitUnDirectedGraph()
	graph.AddEdgeInUndirectedGraphMatrix(1, 2)
	t.Log(graph)
}

func TestAddElementInDirectedGraphMatrix(t *testing.T) {
	graph := InitDirectedGraph()
	graph.AddEdgeInDirectedGraphMatrix(1, 2, 8)
	t.Log(graph)
}

func TestDeleteEdgeInUndirectedGraphMatrix(t *testing.T) {
	graph := InitUnDirectedGraph()
	graph.AddEdgeInUndirectedGraphMatrix(1, 2)
	t.Log(graph)
	graph.DeleteEdgeInUndirectedGraphMatrix(1, 2)
	t.Log(graph)
}

func TestDeleteEdgeInDrirectedGraphMatrix(t *testing.T) {
	graph := InitDirectedGraph()
	graph.AddEdgeInDirectedGraphMatrix(1, 2, 1)
	t.Log(graph)
	graph.DeleteEdgeInDirectedGraphMatrix(1, 2)
	t.Log(graph)
}
