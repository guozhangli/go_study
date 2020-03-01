package TestProject

import (
	"encoding/json"
	"testing"
)

func printGML(graphMList *GraphMList, t *testing.T) {
	str, _ := json.Marshal(graphMList)
	t.Log(string(str))
}

func TestInitGraphMList(t *testing.T) {
	graphMList := InitGraphMList("V0", "V1", "V2", "V3")
	printGML(graphMList, t)
}

func InitGML() *GraphMList {
	return InitGraphMList("V0", "V1", "V2", "V3")
}

func TestAddEdgeInGraphMList(t *testing.T) {
	graphMList := InitGML()
	graphMList.AddEdgeInGraphMList(graphMList.InitEdge(0, 1, 4))
	graphMList.AddEdgeInGraphMList(graphMList.InitEdge(0, 2, 2))
	printGML(graphMList, t)
}
