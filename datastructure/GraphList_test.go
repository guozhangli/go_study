package TestProject

import (
	"encoding/json"
	"testing"
)

func InitGL() *GraphList {
	graph := InitGraphList(5, "V0", "V1", "V2", "V3", "V4")
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
	degree := graph.GetOutDegreeGraphList("V5")
	t.Log(degree)
}
