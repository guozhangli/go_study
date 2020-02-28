package TestProject

import (
	"encoding/json"
	"testing"
)

func TestInitGraphOList(t *testing.T) {
	graphOList := InitGraphOList("V0", "V1", "V2", "V3")
	printGOL(graphOList, t)
}

func printGOL(graph *GraphOList, t *testing.T) {
	str, _ := json.Marshal(graph)
	t.Log(string(str))
}
