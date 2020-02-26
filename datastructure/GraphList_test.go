package TestProject

import (
	"encoding/json"
	"testing"
)

func InitGL() *GraphList {
	graph := InitGraphList(4, "V0", "V1", "V2", "V3")
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
	graph.AddEdgeInGraphList(0, 4, 4)
	t.Log(printGL(graph))
}
