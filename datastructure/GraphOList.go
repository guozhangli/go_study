package TestProject

type LinkNode struct {
	Link   *Arc
	Weight int
}

type Arc struct {
	TailIndex int
	HeadIndex int
	HeadLink  LinkNode
	TailLink  LinkNode
}

type VeriexONode struct {
	Value    interface{}
	FirstIn  *Arc
	FirstOut *Arc
}

//十字链表
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
