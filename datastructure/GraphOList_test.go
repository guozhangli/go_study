package TestProject

import (
	"encoding/json"
	"testing"
)

func TestInitGraphOList(t *testing.T) {
	graphOList := InitGraphOList("V0", "V1", "V2", "V3")
	printGOL(graphOList, t)
}

func InitGOL() *GraphOList {
	graphOList := InitGraphOList("V0", "V1", "V2", "V3")
	return graphOList
}

func printGOL(graph *GraphOList, t *testing.T) {
	str, _ := json.Marshal(graph)
	t.Log(string(str))
}

func createGraphOListOutAct(graphOList *GraphOList) {
	graphOList.AddOutArcInGraphOList(0, 3, 7)
	graphOList.AddOutArcInGraphOList(1, 0, 3)
	graphOList.AddOutArcInGraphOList(1, 2, 4)
	graphOList.AddOutArcInGraphOList(2, 0, 6)
	graphOList.AddOutArcInGraphOList(2, 1, 5)
}

func createGraphOListInAct(graphOList *GraphOList) {
	graphOList.AddInArcInGraphOList(0, 1, 3)
	graphOList.AddInArcInGraphOList(0, 2, 6)
	graphOList.AddInArcInGraphOList(1, 2, 5)
	graphOList.AddInArcInGraphOList(2, 1, 4)
	graphOList.AddInArcInGraphOList(3, 0, 7)
}

func TestAddOutArcInGraphOList(t *testing.T) {
	graphOList := InitGOL()
	createGraphOListOutAct(graphOList)
	printGOL(graphOList, t)
}

func TestAddInArcInGraphOList(t *testing.T) {
	graphOList := InitGOL()
	createGraphOListOutAct(graphOList)
	createGraphOListInAct(graphOList)
	printGOL(graphOList, t)
}

func TestGetOutDegreeInGraphOList(t *testing.T) {
	graphOList := InitGOL()
	createGraphOListOutAct(graphOList)
	createGraphOListInAct(graphOList)
	printGOL(graphOList, t)
	id := graphOList.GetOutDegreeInGraphOList("V3")
	t.Log(id)
}

func TestGetInDegreeInGraphOList(t *testing.T) {
	graphOList := InitGOL()
	createGraphOListOutAct(graphOList)
	createGraphOListInAct(graphOList)
	printGOL(graphOList, t)
	id := graphOList.GetInDegreeInGraphOList("V3")
	t.Log(id)
}
