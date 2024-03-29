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

//计算图顶点的出度
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

//计算图顶点的入度
func (graphList *GraphList) GetInDegreeGraphList(value interface{}) int {
	checkGraphList(graphList)
	index := -1
	for i, v := range graphList.Veriexs {
		if v.Veriex == value {
			index = i
			break
		}
	}
	if index == -1 {
		panic("未查找到顶点")
	}
	var count int
	for _, v := range graphList.Veriexs {
		node := v.FirstEdge
		for node != nil {
			if node.Index == index {
				count++
			}
			node = node.Next
		}
	}
	return count
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

//拓扑排序
func (graphList *GraphList) TopologicalSort() {
	checkGraphList(graphList)
	vNum := len(graphList.Veriexs)
	inDegree := make([]int, vNum) //入度
	stack := NewStackLinked()
	for i, v := range graphList.Veriexs {
		inDegree[i] = graphList.GetInDegreeGraphList(v.Veriex)
		if inDegree[i] == 0 {
			stack.Push(i)
		}
	}
	var e *EdgeNode
	var count int
	for !stack.IsEmpty() {
		gettop := stack.Pop().(*Node).Data.(int)
		fmt.Println(graphList.Veriexs[gettop].Veriex)
		count++
		e = graphList.Veriexs[gettop].FirstEdge
		for e != nil {
			k := e.Index
			if inDegree[k]--; inDegree[k] == 0 {
				stack.Push(k)
			}
			e = e.Next
		}
	}
	if count < vNum {
		panic("存在环")
	}
}

func (graphList *GraphList) TopologicalSort_1() (Stack, []int) {
	checkGraphList(graphList)
	vNum := len(graphList.Veriexs)
	inDegree := make([]int, vNum) //入度
	stack := NewStackLinked()     //用于存储入度为0的栈
	etv := make([]int, vNum)      //事件最早发生时间
	stack2 := NewStackLinked()    //用于存储拓扑序列的栈
	for i, v := range graphList.Veriexs {
		etv[i] = 0 //初始化为0
		inDegree[i] = graphList.GetInDegreeGraphList(v.Veriex)
		if inDegree[i] == 0 {
			stack.Push(i)
		}
	}
	var e *EdgeNode
	var count int
	for !stack.IsEmpty() {
		gettop := stack.Pop().(*Node).Data.(int)
		stack2.Push(gettop)
		count++
		e = graphList.Veriexs[gettop].FirstEdge
		for e != nil {
			k := e.Index
			if inDegree[k]--; inDegree[k] == 0 {
				stack.Push(k)
			}
			if etv[gettop]+e.Weight > etv[k] {
				etv[k] = etv[gettop] + e.Weight
			}
			e = e.Next
		}
	}
	if count < vNum {
		panic("存在环")
	}
	return stack2, etv
}

//关键路径算法
func (graphList *GraphList) KeyPath() {
	sk, etv := graphList.TopologicalSort_1()
	vNum := len(graphList.Veriexs)
	ltv := make([]int, vNum) //事件最迟发生时间
	for i := 0; i < vNum; i++ {
		ltv[i] = etv[vNum-1]
	}
	var ete, lte int //活动最早发生时间和最迟发生时间
	var e *EdgeNode
	stack := sk.(*StackLinked)
	for !stack.IsEmpty() {
		gettop := stack.Pop().(*Node).Data.(int)
		e = graphList.Veriexs[gettop].FirstEdge
		for e != nil {
			k := e.Index
			if ltv[k]-e.Weight < ltv[gettop] {
				ltv[gettop] = ltv[k] - e.Weight
			}
			e = e.Next
		}
	}
	for j := 0; j < vNum; j++ {
		for e := graphList.Veriexs[j].FirstEdge; e != nil; e = e.Next {
			k := e.Index
			ete = etv[j]            //活动最早发生时间
			lte = ltv[k] - e.Weight //活动最迟发生时间
			if ete == lte {
				fmt.Printf("<%v,%v> length: %d,", graphList.Veriexs[j].Veriex, graphList.Veriexs[k].Veriex, e.Weight)
			}
		}
	}
}
