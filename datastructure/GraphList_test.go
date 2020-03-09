package TestProject

import (
	"encoding/json"
	"testing"
)

func InitGL() *GraphList {
	graph := InitGraphList("V0", "V1", "V2", "V3", "V4")
	return graph
}

func printGL(graph *GraphList) string {
	str, _ := json.Marshal(graph)
	return string(str)
}
func TestCreateGraphList(t *testing.T) {
	graph := InitGL()
	t.Log(printGL(graph))
}

func AddEdgeInGL(graph *GraphList) {
	graph.AddEdgeInGraphList(0, 1, 3)
	graph.AddEdgeInGraphList(0, 4, 4)
	graph.AddEdgeInGraphList(1, 2, 4)
	graph.AddEdgeInGraphList(2, 0, 10)
	graph.AddEdgeInGraphList(2, 1, 5)
	graph.AddEdgeInGraphList(2, 3, 7)
	graph.AddEdgeInGraphList(3, 1, 9)
	graph.AddEdgeInGraphList(4, 3, 6)
}

func TestAddEdgeInGraphList(t *testing.T) {
	graph := InitGL()
	AddEdgeInGL(graph)
	t.Log(printGL(graph))
}

func TestDeleteEdgeInGraphList(t *testing.T) {
	graph := InitGL()
	AddEdgeInGL(graph)
	t.Log(printGL(graph))
	graph.DeleteEdgeInGraphList(0, 4)
	graph.DeleteEdgeInGraphList(1, 2)
	graph.DeleteEdgeInGraphList(2, 1)
	t.Log(printGL(graph))
}

func TestIsEdgeInGraphList(t *testing.T) {
	graph := InitGL()
	AddEdgeInGL(graph)
	t.Log(printGL(graph))
	b := graph.IsEdgeInGraphList(0, 4)
	t.Log(b)
}

func TestGetOutDegreeGraphList(t *testing.T) {
	graph := InitGL()
	AddEdgeInGL(graph)
	t.Log(printGL(graph))
	degree := graph.GetOutDegreeGraphList("V4")
	t.Log(degree)
}

func TestGetInDegreeGraphList(t *testing.T) {
	graph := InitGL()
	AddEdgeInGL(graph)
	t.Log(printGL(graph))
	degree := graph.GetInDegreeGraphList("V3")
	t.Log(degree)
}

func TestGraphList_DfsTraverse(t *testing.T) {
	graph := InitGL()
	AddEdgeInGL(graph)
	t.Log(graph)
	graph.DfsTraverseList()
}

func TestGraphList_DfsTraverseStack(t *testing.T) {
	graph := InitGL()
	AddEdgeInGL(graph)
	t.Log(graph)
	graph.DfsTraverseListStack()
}

func TestGraphList_BfsTraverseStack(t *testing.T) {
	graph := InitGL()
	AddEdgeInGL(graph)
	t.Log(graph)
	graph.BfsTraverseListQueue()
}

func InitGL_TopologicalSort() *GraphList {
	graph := InitGraphList("V0", "V1", "V2", "V3", "V4", "V5", "V6", "V7", "V8", "V9", "V10", "V11", "V12", "V13")
	return graph
}
func AddEdgeInGL_TopologicalSort(graph *GraphList) {
	graph.AddEdgeInGraphList(0, 4, 3)
	graph.AddEdgeInGraphList(0, 5, 4)
	graph.AddEdgeInGraphList(0, 11, 4)
	graph.AddEdgeInGraphList(1, 8, 10)
	graph.AddEdgeInGraphList(1, 4, 5)
	graph.AddEdgeInGraphList(1, 2, 7)
	graph.AddEdgeInGraphList(2, 9, 9)
	graph.AddEdgeInGraphList(2, 6, 6)
	graph.AddEdgeInGraphList(2, 5, 8)
	graph.AddEdgeInGraphList(3, 13, 2)
	graph.AddEdgeInGraphList(3, 2, 1)
	graph.AddEdgeInGraphList(4, 7, 4)
	graph.AddEdgeInGraphList(5, 12, 5)
	graph.AddEdgeInGraphList(5, 8, 6)
	graph.AddEdgeInGraphList(6, 5, 3)
	graph.AddEdgeInGraphList(8, 7, 7)
	graph.AddEdgeInGraphList(9, 11, 7)
	graph.AddEdgeInGraphList(9, 10, 8)
	graph.AddEdgeInGraphList(10, 13, 9)
	graph.AddEdgeInGraphList(12, 9, 1)
}

func TestTopologicalSort(t *testing.T) {
	graph := InitGL_TopologicalSort()
	AddEdgeInGL_TopologicalSort(graph)
	t.Log(printGL(graph))
	graph.TopologicalSort()
}
func printStack(stack Stack) string {
	str, _ := json.Marshal(stack)
	return string(str)
}
func InitGL_KeyPath() *GraphList {
	graph := InitGraphList("V0", "V1", "V2", "V3", "V4", "V5", "V6", "V7", "V8", "V9")
	return graph
}
func AddEdgeInGL_KeyPath(graph *GraphList) {
	graph.AddEdgeInGraphList(0, 2, 4)
	graph.AddEdgeInGraphList(0, 1, 3)
	graph.AddEdgeInGraphList(1, 4, 6)
	graph.AddEdgeInGraphList(1, 3, 5)
	graph.AddEdgeInGraphList(2, 5, 7)
	graph.AddEdgeInGraphList(2, 3, 8)
	graph.AddEdgeInGraphList(3, 4, 3)
	graph.AddEdgeInGraphList(4, 7, 4)
	graph.AddEdgeInGraphList(4, 6, 9)
	graph.AddEdgeInGraphList(5, 7, 6)
	graph.AddEdgeInGraphList(6, 9, 2)
	graph.AddEdgeInGraphList(7, 8, 5)
	graph.AddEdgeInGraphList(8, 9, 3)
}
func TestKeyPath(t *testing.T) {
	graph := InitGL_KeyPath()
	AddEdgeInGL_KeyPath(graph)
	t.Log(printGL(graph))
	graph.KeyPath()
}
