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

func AddEdgeInDirectedGraphMatrix(graph *GraphMatrix) {
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
	AddEdgeInDirectedGraphMatrix(graph)
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

func TestGraphMatrix_BfsTraverseQueue(t *testing.T) {
	graph := InitUnDirectedGraph()
	AddEdgeInUnDirectedGraphMatrix(graph)
	t.Log(graph)
	graph.BfsTraverseMatirxQueue()
}
func InitDirectedGraph_Prim() *GraphMatrix {
	graph := InitDirectedGraphMatrix("V0", "V1", "V2", "V3", "V4", "V5", "V6", "V7", "V8")
	return graph
}
func AddEdgeInDirectedGraphMatrix_Prim(graph *GraphMatrix) {
	graph.AddEdgeInDirectedGraphMatrix(0, 1, 10)
	graph.AddEdgeInDirectedGraphMatrix(0, 5, 11)
	graph.AddEdgeInDirectedGraphMatrix(1, 0, 10)
	graph.AddEdgeInDirectedGraphMatrix(1, 2, 18)
	graph.AddEdgeInDirectedGraphMatrix(1, 6, 16)
	graph.AddEdgeInDirectedGraphMatrix(1, 8, 12)
	graph.AddEdgeInDirectedGraphMatrix(2, 1, 18)
	graph.AddEdgeInDirectedGraphMatrix(2, 3, 22)
	graph.AddEdgeInDirectedGraphMatrix(2, 8, 8)
	graph.AddEdgeInDirectedGraphMatrix(3, 2, 22)
	graph.AddEdgeInDirectedGraphMatrix(3, 8, 8)
	graph.AddEdgeInDirectedGraphMatrix(4, 3, 20)
	graph.AddEdgeInDirectedGraphMatrix(4, 5, 26)
	graph.AddEdgeInDirectedGraphMatrix(4, 7, 7)
	graph.AddEdgeInDirectedGraphMatrix(5, 0, 11)
	graph.AddEdgeInDirectedGraphMatrix(5, 4, 26)
	graph.AddEdgeInDirectedGraphMatrix(5, 6, 17)
	graph.AddEdgeInDirectedGraphMatrix(6, 1, 16)
	graph.AddEdgeInDirectedGraphMatrix(6, 5, 17)
	graph.AddEdgeInDirectedGraphMatrix(6, 7, 19)
	graph.AddEdgeInDirectedGraphMatrix(7, 3, 16)
	graph.AddEdgeInDirectedGraphMatrix(7, 4, 7)
	graph.AddEdgeInDirectedGraphMatrix(7, 6, 19)
	graph.AddEdgeInDirectedGraphMatrix(8, 1, 12)
	graph.AddEdgeInDirectedGraphMatrix(8, 2, 8)
	graph.AddEdgeInDirectedGraphMatrix(8, 3, 21)
}
func TestMSTByPrim(t *testing.T) {
	graph := InitDirectedGraph_Prim()
	AddEdgeInDirectedGraphMatrix_Prim(graph)
	t.Log(graph)
	MST, sum := graph.MiniSpanTree_Prim()
	t.Log(MST)
	t.Log(sum)
}
