package TestProject

type EdgeNode struct {
	Index  int
	Weight int
	Next   *EdgeNode
}

type VeriexNode struct {
	Veriex    interface{}
	FirstEdge *EdgeNode
}

type GraphList struct {
	Veriexs []*VeriexNode
}

func InitGraphList(verNum int, value ...interface{}) *GraphList {
	graph := &GraphList{
		Veriexs: make([]*VeriexNode, verNum),
	}
	for i, v := range value {
		veriexNode := &VeriexNode{
			Veriex:    v,
			FirstEdge: nil,
		}
		graph.Veriexs[i] = veriexNode
	}
	return graph
}

func checkGraphList(graphList *GraphList) {
	if graphList == nil {
		panic("未创建图")
	}
}

func (graphList *GraphList) AddEdgeInGraphList(i int, j int, weight int) {
	checkGraphList(graphList)
}
