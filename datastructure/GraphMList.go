package TestProject

//邻接多重表(表示无向图)
type Edge struct {
	IVex   int
	ILink  *Edge
	JVex   int
	JLink  *Edge
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

func (graphMList *GraphMList) InitEdge(iVer int, jVer int, weight int) *Edge {
	if iVer < 0 || iVer >= len(graphMList.Veriexs) || jVer < 0 || jVer >= len(graphMList.Veriexs) {
		panic("顶点索引值不合法")
	}
	edge := &Edge{
		IVex:   iVer,
		ILink:  nil,
		JVex:   jVer,
		JLink:  nil,
		Weight: weight,
	}
	return edge
}

func (graphMList *GraphMList) AddEdgeInGraphMList(edge *Edge) {
	checkGraphMList(graphMList)
	if edge != nil {
		iVer := edge.IVex
		ver := graphMList.Veriexs[iVer]
		if ver.FirstEdge == nil {
			ver.FirstEdge = edge
		} else {
			node := ver.FirstEdge
			for node.ILink != nil {
				node = node.ILink
			}
			node.ILink = edge
		}
	}

}
