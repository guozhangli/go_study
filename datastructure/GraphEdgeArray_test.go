package TestProject

import (
	"encoding/json"
	"testing"
)

func InitGEA() *GraphEdgeArray {
	graph := InitGraphEdgeArray("V0", "V1", "V2", "V3", "V4", "V5", "V6", "V7", "V8")
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

func TestAddEdgeInGraphEA(t *testing.T) {
	graphEdgeArray := InitGEA()
	graphEdgeArray.AddEdgeInGraphEA(0, 1, 10)
	graphEdgeArray.AddEdgeInGraphEA(0, 5, 11)
	graphEdgeArray.AddEdgeInGraphEA(1, 2, 18)
	graphEdgeArray.AddEdgeInGraphEA(1, 6, 16)
	graphEdgeArray.AddEdgeInGraphEA(1, 8, 12)
	graphEdgeArray.AddEdgeInGraphEA(2, 3, 22)
	graphEdgeArray.AddEdgeInGraphEA(2, 8, 8)
	graphEdgeArray.AddEdgeInGraphEA(3, 4, 20)
	graphEdgeArray.AddEdgeInGraphEA(3, 4, 17)
	graphEdgeArray.AddEdgeInGraphEA(3, 8, 21)
	graphEdgeArray.AddEdgeInGraphEA(3, 7, 16)
	graphEdgeArray.AddEdgeInGraphEA(4, 5, 26)
	graphEdgeArray.AddEdgeInGraphEA(4, 7, 7)
	graphEdgeArray.AddEdgeInGraphEA(5, 6, 17)
	graphEdgeArray.AddEdgeInGraphEA(6, 7, 19)
	t.Log(printGEA(graphEdgeArray))
	graph := graphEdgeArray.MiniSpanTree_Kruskal()
	t.Log(printGEA(graph))
}
