package TestProject

import (
	"encoding/json"
	"testing"
)

func InitGEA() *GraphEdgeArray {
	graph := InitGraphEdgeArray("V0", "V1", "V2", "V3")
	return graph
}

func printGEA(graph *GraphEdgeArray) string {
	str, _ := json.Marshal(graph)
	return string(str)
}
func TestInitGraphEdgeArray(t *testing.T) {
	graphEdgeArray := InitGEA()
	t.Log(printGEA(graphEdgeArray))
}
