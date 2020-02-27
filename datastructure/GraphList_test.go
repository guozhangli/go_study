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

func TestAddEdgeInGraphList(t *testing.T) {
	graph := InitGL()
	graph.AddEdgeInGraphList(0, 1, 3)
	graph.AddEdgeInGraphList(0, 4, 4)
	graph.AddEdgeInGraphList(1, 2, 4)
	graph.AddEdgeInGraphList(2, 0, 10)
	graph.AddEdgeInGraphList(2, 1, 5)
	graph.AddEdgeInGraphList(2, 3, 7)
	graph.AddEdgeInGraphList(3, 1, 9)
	graph.AddEdgeInGraphList(4, 3, 6)
	t.Log(printGL(graph))
}

func TestDeleteEdgeInGraphList(t *testing.T) {
	graph := InitGL()
	graph.DeleteEdgeInGraphList(0, 1)
	t.Log(printGL(graph))
}
