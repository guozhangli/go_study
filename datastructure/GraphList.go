package TestProject

import (
	"encoding/json"
	"fmt"
)

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
	if graphList.Veriexs == nil {
		panic("未创建图的顶点")
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
	if i < 0 || i >= len(graphList.Veriexs) || j >= len(graphList.Veriexs) || j < 0 {
		panic("顶点索引不合法")
	}
	edge := &EdgeNode{
		Index:  j,
		Weight: weight,
		Next:   nil}
	for i1, v := range graphList.Veriexs {
		if i1 == i {
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
}

func (graphList *GraphList) DeleteEdgeInGraphList(i int, j int) {
	checkGraphList(graphList)
	if i < 0 || i >= len(graphList.Veriexs) || j >= len(graphList.Veriexs) || j < 0 {
		panic("顶点索引不合法")
	}
	for i1, v := range graphList.Veriexs {
		if i1 == i {
			var pre *EdgeNode
			node := v.FirstEdge
			for {
				if node != nil && node.Index != j {
					pre = node
					node = node.Next
				}
				if node.Index == j {
					if pre != nil {
						pre.Next = node.Next
					} else {
						pre = node.Next
					}
					break
				}
			}
			v.FirstEdge = pre
			str, _ := json.Marshal(pre)
			fmt.Println(string(str))
		}
	}

}

func (graphList *GraphList) IsEdgeInGraphList(i int, j int) bool {
	checkGraphList(graphList)
	if i < 0 || i >= len(graphList.Veriexs) || j >= len(graphList.Veriexs) || j < 0 {
		panic("顶点索引不合法")
	}
	flag := false
	for i1, v := range graphList.Veriexs {
		if i1 == i {
			node := v.FirstEdge
			for node != nil {
				if node.Index != j {
					node = node.Next
				} else {
					return !flag
				}
			}
		}
	}
	return flag
}
