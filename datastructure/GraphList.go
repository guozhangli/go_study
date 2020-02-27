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

/*
	add edge in graph list
	param i Veriexs index
      	  j Edge index
		  weight Weight
*/
func (graphList *GraphList) AddEdgeInGraphList(i int, j int, weight int) {
	checkGraphList(graphList)
	edge := &EdgeNode{
		Index:  j,
		Weight: weight,
		Next:   nil}
	var flag = true
	for i1, v := range graphList.Veriexs {
		if i1 == i {
			flag = false
			if v.FirstEdge != nil {
				var node = v.FirstEdge
				for node.Next != nil {
					node = node.Next
				}
				node.Next = edge
			} else {
				v.FirstEdge = edge
			}
		}
	}
	if flag {
		panic("输入参数不正确")
	}
}

func (graphList *GraphList) DeleteEdgeInGraphList(i int, j int) {
	checkGraphList(graphList)
	var flag = true
	for i1, v := range graphList.Veriexs {
		if i1 == i {

			if v.FirstEdge != nil {
				node := v.FirstEdge
				for node != nil {
					if node.Index == j {
						flag = false
					}
					node = node.Next
				}

			}
		}
	}
	if flag {
		panic("输入参数不正确")
	}
}
