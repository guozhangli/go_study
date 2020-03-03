package TestProject

import "fmt"

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

func InitGraphList(value ...interface{}) *GraphList {
	verNum := len(value)
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
			for node != nil {
				if node.Index != j {
					pre = node
					node = node.Next
				} else {
					if pre != nil {
						pre.Next = node.Next
					} else {
						pre = node.Next
					}
					break
				}
			}
			v.FirstEdge = pre
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

func (graphList *GraphList) GetOutDegreeGraphList(value interface{}) int {
	checkGraphList(graphList)
	flag := false
	for _, v := range graphList.Veriexs {
		if v.Veriex == value {
			flag = true
			node := v.FirstEdge
			count := 0
			for node != nil {
				node = node.Next
				count++
			}
			return count
		}
	}
	if !flag {
		panic("veriex is not exist")
	}
	return 0
}

var veriexFlag_1 []bool

//深度优先遍历-邻接表（递归方式）
func (graphList *GraphList) DfsTraverseList() {
	checkGraphList(graphList)
	veriexFlag_1 = make([]bool, len(graphList.Veriexs))
	for i, _ := range graphList.Veriexs {
		veriexFlag_1[i] = false
	}
	for i, _ := range graphList.Veriexs {
		if !veriexFlag_1[i] {
			graphList.DfsGraphList(i)
		}
	}
}

func (graphList *GraphList) DfsGraphList(i int) {
	veriexFlag_1[i] = true
	fmt.Println(graphList.Veriexs[i].Veriex)
	p := graphList.Veriexs[i].FirstEdge
	for p != nil {
		if !veriexFlag_1[p.Index] {
			graphList.DfsGraphList(p.Index)
		}
		p = p.Next
	}
}

//深度优先遍历-邻接表（运用栈）
func (graphList *GraphList) DfsTraverseListStack() {
	checkGraphList(graphList)
	veriexFlag_1 = make([]bool, len(graphList.Veriexs))
	veriexFlag_1[0] = true
	fmt.Println(graphList.Veriexs[0].Veriex)
	stack := NewStackLinked()
	stack.Push(0)
	for !stack.IsEmpty() {
		i := graphList.GetUnvisitedVeriex(stack.GetTop().(int))
		if i == -1 {
			stack.Pop()
		} else {
			veriexFlag_1[i] = true
			fmt.Println(graphList.Veriexs[i].Veriex)
			stack.Push(i)
		}
	}
	for i, _ := range graphList.Veriexs {
		veriexFlag_1[i] = false
	}
}

func (graphList *GraphList) GetUnvisitedVeriex(v int) int {
	p := graphList.Veriexs[v].FirstEdge
	for p != nil {
		if !veriexFlag_1[p.Index] {
			return p.Index
		}
		p = p.Next
	}
	return -1
}

//广度优先遍历-邻接表（运用队列）
func (graphList *GraphList) BfsTraverseListQueue() {
	checkGraphList(graphList)
	veriexFlag_1 = make([]bool, len(graphList.Veriexs))
	for i, _ := range graphList.Veriexs {
		veriexFlag_1[i] = false
	}
	queue := NewQueueLinked()
	for i, v := range graphList.Veriexs {
		if !veriexFlag_1[i] {
			veriexFlag_1[i] = true
			fmt.Println(v.Veriex)
			queue.EnQueue(i)
			for !queue.IsEmpty() {
				i = queue.DeQueue().(*Node).Data.(int)
				p := graphList.Veriexs[i].FirstEdge
				for p != nil {
					if !veriexFlag_1[p.Index] {
						veriexFlag_1[p.Index] = true
						fmt.Println(graphList.Veriexs[p.Index].Veriex)
						queue.EnQueue(p.Index)
					}
					p = p.Next
				}
			}
		}
	}
}
