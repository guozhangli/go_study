package TestProject

//邻接多重表(表示无向图)
type Edge struct {
	IVex   int
	ILink  *Edge
	JVex   int
	JLinke *Edge
	Weight int
}

type VeriexMNode struct {
	value     interface{}
	FirstEdge *Edge
}

type GraphMList struct {
	Veriexs []*VeriexMNode
}

func InitGraphMList(value ...interface{}) *GraphMList {
	graphMList := &GraphMList{
		Veriexs: make([]*VeriexMNode, len(value)),
	}
	for i, v := range value {
		ver := &VeriexMNode{
			value:     v,
			FirstEdge: nil,
		}
		graphMList.Veriexs[i] = ver
	}
	return graphMList
}

func checkGraphMList(graphMList *GraphMList) {
	if graphMList == nil {
		panic("未创建图")
	}
	if graphMList.Veriexs == nil {
		panic("未创建图的顶点")
	}
}

func (graphMList *GraphMList) AddEdgeInGraphMList(iVer int, jVer int, weight int) {
	checkGraphMList(graphMList)
	edge := &Edge{
		IVex:   iVer,
		ILink:  nil,
		JVex:   jVer,
		JLinke: nil,
		Weight: weight,
	}
	ver := graphMList.Veriexs[iVer]

}
