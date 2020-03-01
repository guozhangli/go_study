package TestProject

//十字链表（表示有向图--包含邻接表和逆邻接表）

type LinkNode struct {
	Link   *Arc
	Weight int
}

type Arc struct {
	TailIndex int
	HeadIndex int
	HeadLink  *LinkNode
	TailLink  *LinkNode
}

type VeriexONode struct {
	Value    interface{}
	FirstIn  *Arc
	FirstOut *Arc
}

type GraphOList struct {
	Veriexs []*VeriexONode
}

func InitGraphOList(value ...interface{}) *GraphOList {
	graphOL := &GraphOList{
		Veriexs: make([]*VeriexONode, len(value)),
	}
	for i, v := range value {
		veriexONode := &VeriexONode{
			Value:    v,
			FirstIn:  nil,
			FirstOut: nil,
		}
		graphOL.Veriexs[i] = veriexONode
	}
	return graphOL
}

func CheckGraphOList(graphOList *GraphOList) {
	if graphOList == nil {
		panic("未创建图")
	}
	if graphOList.Veriexs == nil {
		panic("未创建图的顶点")
	}
}

func (graphOList *GraphOList) AddOutArcInGraphOList(ti int, hi int, tw int) {
	CheckGraphOList(graphOList)
	if ti < 0 || ti >= len(graphOList.Veriexs) || hi < 0 || hi >= len(graphOList.Veriexs) {
		panic("顶点索引值不合法")
	}
	arc := &Arc{
		TailIndex: ti,
		HeadIndex: hi,
		HeadLink: &LinkNode{
			Link:   nil,
			Weight: INFINITY,
		},
		TailLink: &LinkNode{
			Link:   nil,
			Weight: tw,
		},
	}
	ver := graphOList.Veriexs[ti]
	if ver.FirstOut != nil {
		node := ver.FirstOut
		for node.TailLink.Link != nil {
			node = node.TailLink.Link
		}
		node.TailLink.Link = arc
	} else {
		ver.FirstOut = arc
	}
}

func (graphOList *GraphOList) AddInArcInGraphOList(ti int, hi int, hw int) {
	CheckGraphOList(graphOList)
	if ti < 0 || ti >= len(graphOList.Veriexs) || hi < 0 || hi >= len(graphOList.Veriexs) {
		panic("顶点索引值不合法")
	}
	ver := graphOList.Veriexs[ti]
	hVer := graphOList.Veriexs[hi]
	node := hVer.FirstOut
	for node != nil && node.HeadIndex != ti {
		node = node.TailLink.Link
	}

	if ver.FirstIn != nil {
		tempNode := ver.FirstIn
		for tempNode.HeadLink.Link != nil {
			tempNode = tempNode.HeadLink.Link
		}
		tempNode.HeadLink.Link = node
	} else {
		ver.FirstIn = node
	}
}

func (graphOList *GraphOList) GetVeriex(value interface{}) *VeriexONode {
	CheckGraphOList(graphOList)
	var veriex *VeriexONode
	for _, v := range graphOList.Veriexs {
		if v.Value == value {
			veriex = v
			break
		}
	}
	return veriex
}

func (graphOList *GraphOList) GetOutDegreeInGraphOList(value interface{}) int {
	CheckGraphOList(graphOList)
	veriex := graphOList.GetVeriex(value)
	if veriex == nil {
		panic("顶点不存在")
	}
	arc := veriex.FirstOut
	var count int
	for arc != nil {
		arc = arc.TailLink.Link
		count++
	}
	return count
}

func (graphOList *GraphOList) GetInDegreeInGraphOList(value interface{}) int {
	CheckGraphOList(graphOList)
	veriex := graphOList.GetVeriex(value)
	if veriex == nil {
		panic("顶点不存在")
	}
	arc := veriex.FirstIn
	var count int
	for arc != nil {
		arc = arc.HeadLink.Link
		count++
	}
	return count
}
