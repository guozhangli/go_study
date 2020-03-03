package TestProject

import "testing"

func InitUnDirectedGraph() *GraphMatrix {
	graph := InitUndirectedGraphMatrix("V0", "V1", "V2", "V3", "V4")
	return graph
}

func InitDirectedGraph() *GraphMatrix {
	graph := InitDirectedGraphMatrix("V0", "V1", "V2", "V3", "V4")
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

func AddEdgeInUnDirectedGraphMatrix(graph *GraphMatrix) {
	graph.AddEdgeInUndirectedGraphMatrix(0, 1)
	graph.AddEdgeInUndirectedGraphMatrix(0, 2)
	graph.AddEdgeInUndirectedGraphMatrix(0, 3)
	graph.AddEdgeInUndirectedGraphMatrix(1, 0)
	graph.AddEdgeInUndirectedGraphMatrix(1, 2)
	graph.AddEdgeInUndirectedGraphMatrix(2, 0)
	graph.AddEdgeInUndirectedGraphMatrix(2, 1)
	graph.AddEdgeInUndirectedGraphMatrix(2, 3)
	graph.AddEdgeInUndirectedGraphMatrix(2, 4)
	graph.AddEdgeInUndirectedGraphMatrix(3, 0)
	graph.AddEdgeInUndirectedGraphMatrix(3, 2)
}

//无向图添加边
func TestAddElementInGraphMatrix(t *testing.T) {
	graph := InitUnDirectedGraph()
	AddEdgeInUnDirectedGraphMatrix(graph)
	t.Log(graph)
}

func AddEdgeInUndriectedGraphMatrix(graph *GraphMatrix) {
	graph.AddEdgeInDirectedGraphMatrix(0, 1, 3)
	graph.AddEdgeInDirectedGraphMatrix(0, 4, 4)
	graph.AddEdgeInDirectedGraphMatrix(1, 2, 4)
	graph.AddEdgeInDirectedGraphMatrix(2, 0, 10)
	graph.AddEdgeInDirectedGraphMatrix(2, 1, 5)
	graph.AddEdgeInDirectedGraphMatrix(2, 3, 7)
	graph.AddEdgeInDirectedGraphMatrix(3, 1, 9)
	graph.AddEdgeInDirectedGraphMatrix(4, 3, 6)
}

//有相图添加边
func TestAddElementInDirectedGraphMatrix(t *testing.T) {
	graph := InitDirectedGraph()
	AddEdgeInUndriectedGraphMatrix(graph)
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

func TestIsEdgeInUndirectedGraphMatrix(t *testing.T) {
	graph := InitUnDirectedGraph()
	graph.AddEdgeInUndirectedGraphMatrix(1, 2)
	t.Log(graph)
	b := graph.IsEdgeInUndirectedGraphMatrix(1, 3)
	t.Log(b)
}

func TestIsEdgeInDirectedGraphMatrix(t *testing.T) {
	graph := InitDirectedGraph()
	graph.AddEdgeInDirectedGraphMatrix(1, 2, 4)
	t.Log(graph)
	b := graph.IsEdgeInDirectedGraphMatrix(1, 3)
	t.Log(b)
}

func TestGraphMatrix_DfsTraverse(t *testing.T) {
	graph := InitUnDirectedGraph()
	AddEdgeInUnDirectedGraphMatrix(graph)
	t.Log(graph)
	graph.DfsTraverseMatirx()
}

func TestGraphMatrix_DfsTraverseStack(t *testing.T) {
	graph := InitUnDirectedGraph()
	AddEdgeInUnDirectedGraphMatrix(graph)
	t.Log(graph)
	graph.DfsTraverseMatirxStack()
}
