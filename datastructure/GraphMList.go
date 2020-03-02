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
	Value     interface{}
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
			Value:     v,
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
	//head add
	if edge != nil {
		iVer := edge.IVex
		iEdge := graphMList.Veriexs[iVer]
		edge.ILink = iEdge.FirstEdge
		iEdge.FirstEdge = edge
		jVer := edge.JVex
		jEdge := graphMList.Veriexs[jVer]
		edge.JLink = jEdge.FirstEdge
		jEdge.FirstEdge = edge
	}
}

func (graphMList *GraphMList) GetEdgeInGraphMList(iVer int, jVer int) (*Edge, *Edge) {
	checkGraphMList(graphMList)
	if iVer < 0 || iVer >= len(graphMList.Veriexs) || jVer < 0 || jVer >= len(graphMList.Veriexs) {
		panic("顶点索引值不合法")
	}
	ver := graphMList.Veriexs[iVer]
	node := ver.FirstEdge
	flag := false
	var pre *Edge
	for node != nil {
		if node.IVex == iVer && node.JVex == jVer {
			flag = true
			return pre, node
		} else {
			pre = node
			if node.IVex == iVer {
				node = node.ILink
			} else if node.JVex == iVer {
				node = node.JLink
			}
		}
	}
	if !flag {
		panic("未查询出删除的边")
	}
	return nil, nil
}

func (graphMList *GraphMList) DeleteEdgeInGraphMList(iVer int, jVer int) {
	checkGraphMList(graphMList)
	if iVer < 0 || iVer >= len(graphMList.Veriexs) || jVer < 0 || jVer >= len(graphMList.Veriexs) {
		panic("顶点索引值不合法")
	}
	iver := graphMList.Veriexs[iVer]
	node := iver.FirstEdge
	var pre *Edge
	for node != nil && !(node.IVex == iVer && node.JVex == jVer) {
		pre = node
		if node.IVex == iVer {
			node = node.ILink
		} else if node.JVex == iVer {
			node = node.JLink
		}
	}
	if node == nil {
		panic("未查询出删除的边")
	}
	if pre != nil {
		if pre.IVex == iVer {
			if node.IVex == iVer {
				pre.ILink = node.ILink
			} else if node.JVex == iVer {
				pre.ILink = node.JLink
			}
		} else if pre.JVex == iVer {
			if node.IVex == iVer {
				pre.JLink = node.ILink
			} else if node.JVex == iVer {
				pre.JLink = node.JLink
			}
		}
	} else {
		if node.IVex == iVer {
			pre = node.ILink
		} else if node.JVex == iVer {
			pre = node.JLink
		}
	}
	iver.FirstEdge = pre
	jver := graphMList.Veriexs[jVer]
	node1 := jver.FirstEdge
	var pre1 *Edge
	for node1 != nil && !(node1.IVex == node.IVex && node1.JVex == node.JVex) {
		pre1 = node1
		if node1.IVex == jVer {
			node1 = node1.ILink
		} else if node1.JVex == jVer {
			node1 = node1.JLink
		}
	}
	if pre1 != nil {
		if pre1.IVex == jVer {
			if node1.IVex == jVer {
				pre1.ILink = node1.ILink
			} else if node1.JVex == jVer {
				pre1.ILink = node1.JLink
			}
		} else if pre1.JVex == jVer {
			if node1.IVex == jVer {
				pre1.JLink = node1.ILink
			} else if node1.JVex == jVer {
				pre1.JLink = node1.JLink
			}
		}
	} else {
		if node1.IVex == jVer {
			pre1 = node1.ILink
		} else if node1.JVex == jVer {
			pre1 = node1.JLink
		}
	}
	jver.FirstEdge = pre1
}
