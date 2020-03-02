package TestProject

import (
	"encoding/json"
	"testing"
)

func InitGML() *GraphMList {
	return InitGraphMList("V0", "V1", "V2", "V3")
}

func printGML(graphMList *GraphMList) string {
	str, _ := json.Marshal(graphMList)
	return string(str)
}

func TestInitGraphMList(t *testing.T) {
	graphMList := InitGML()
	t.Log(printGML(graphMList))
}

func AddEdgeInGML(graphMList *GraphMList) {
	graphMList.AddEdgeInGraphMList(graphMList.InitEdge(3, 0, 5))
	graphMList.AddEdgeInGraphMList(graphMList.InitEdge(0, 2, 4))
	graphMList.AddEdgeInGraphMList(graphMList.InitEdge(0, 1, 2))
	graphMList.AddEdgeInGraphMList(graphMList.InitEdge(1, 2, 3))
	graphMList.AddEdgeInGraphMList(graphMList.InitEdge(2, 3, 6))
}

func TestAddEdgeInGraphMList(t *testing.T) {
	graphMList := InitGML()
	AddEdgeInGML(graphMList)
	t.Log(printGML(graphMList))
}

func TestGetEdgeInGraphMList(t *testing.T) {
	graphMList := InitGML()
	AddEdgeInGML(graphMList)
	t.Log(printGML(graphMList))
	pre, edge := graphMList.GetEdgeInGraphMList(0, 1)
	str, _ := json.Marshal(edge)
	t.Log(string(str))
	str1, _ := json.Marshal(pre)
	t.Log(string(str1))
}

func TestDeleteEdgeInGraphMList(t *testing.T) {
	graphMList := InitGML()
	AddEdgeInGML(graphMList)
	t.Log(printGML(graphMList))
	graphMList.DeleteEdgeInGraphMList(0, 2)
	t.Log(printGML(graphMList))
}
