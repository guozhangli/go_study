package TestProject

import (
	"fmt"
	"math"
)

type MyEdgeType int

const INFINITY = math.MaxInt32

var veriexFlag []bool

//图（邻接矩阵）
type GraphMatrix struct {
	Veriex []interface{}  //顶点集合
	Edge   [][]MyEdgeType //边集合
}

func checkGraphMatrix(graph *GraphMatrix) {
	if graph == nil {
		panic("未创建图")
	}
}

func InitUndirectedGraphMatrix(value ...interface{}) *GraphMatrix {
	verNum := len(value)
	graph := &GraphMatrix{
		Veriex: make([]interface{}, verNum),
		Edge:   make([][]MyEdgeType, verNum),
	}
	for i, v := range value {
		graph.Veriex[i] = v
	}
	for i := 0; i < verNum; i++ {
		graph.Edge[i] = make([]MyEdgeType, verNum)
		for j := 0; j < verNum; j++ {
			graph.Edge[i][j] = 0
		}
	}
	return graph
}

func InitDirectedGraphMatrix(value ...interface{}) *GraphMatrix {
	verNum := len(value)
	graph := &GraphMatrix{
		Veriex: make([]interface{}, verNum),
		Edge:   make([][]MyEdgeType, verNum),
	}
	for i, v := range value {
		graph.Veriex[i] = v
	}
	for i := 0; i < verNum; i++ {
		graph.Edge[i] = make([]MyEdgeType, verNum)
		for j := 0; j < verNum; j++ {
			if i == j {
				graph.Edge[i][j] = 0
			} else {
				graph.Edge[i][j] = INFINITY
			}
		}
	}
	return graph
}

func (graphMatrix *GraphMatrix) AddEdgeInUndirectedGraphMatrix(i int, j int) {
	checkGraphMatrix(graphMatrix)
	if i < len(graphMatrix.Veriex) && i >= 0 && j < len(graphMatrix.Veriex) && j >= 0 {
		graphMatrix.Edge[i][j] = 1
	} else {
		panic("输入坐标不合法")
	}
}

func (graphMatrix *GraphMatrix) AddEdgeInDirectedGraphMatrix(i int, j int, weight int) {
	checkGraphMatrix(graphMatrix)
	if i < len(graphMatrix.Veriex) && i >= 0 && j < len(graphMatrix.Veriex) && j >= 0 {
		graphMatrix.Edge[i][j] = MyEdgeType(weight)
	} else {
		panic("输入坐标不合法")
	}
}

func (graphMatrix *GraphMatrix) DeleteEdgeInUndirectedGraphMatrix(i int, j int) {
	checkGraphMatrix(graphMatrix)
	if i < len(graphMatrix.Veriex) && i >= 0 && j < len(graphMatrix.Veriex) && j >= 0 {
		graphMatrix.Edge[i][j] = 0
	} else {
		panic("输入坐标不合法")
	}
}

func (graphMatrix *GraphMatrix) DeleteEdgeInDirectedGraphMatrix(i int, j int) {
	checkGraphMatrix(graphMatrix)
	if i < len(graphMatrix.Veriex) && i >= 0 && j < len(graphMatrix.Veriex) && j >= 0 {
		if i == j {
			graphMatrix.Edge[i][j] = 0
		} else {
			graphMatrix.Edge[i][j] = INFINITY
		}
	} else {
		panic("输入坐标不合法")
	}
}

func (graphMatrix *GraphMatrix) IsEdgeInUndirectedGraphMatrix(i int, j int) bool {
	checkGraphMatrix(graphMatrix)
	if i < len(graphMatrix.Veriex) && i >= 0 && j < len(graphMatrix.Veriex) && j >= 0 {
		if graphMatrix.Edge[i][j] != 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (graphMatrix *GraphMatrix) IsEdgeInDirectedGraphMatrix(i int, j int) interface{} {
	checkGraphMatrix(graphMatrix)
	if i < len(graphMatrix.Veriex) && i >= 0 && j < len(graphMatrix.Veriex) && j >= 0 {
		if graphMatrix.Edge[i][j] != 0 && graphMatrix.Edge[i][j] != INFINITY {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

//深度优先遍历-邻接矩阵（递归遍历）
func (graphMatrix *GraphMatrix) DfsTraverseMatirx() {
	checkGraphMatrix(graphMatrix)
	veriexFlag = make([]bool, len(graphMatrix.Veriex))
	for i, _ := range graphMatrix.Veriex {
		veriexFlag[i] = false
	}
	for i, _ := range graphMatrix.Veriex {
		if !veriexFlag[i] {
			graphMatrix.DfsGraphMatirx(i)
		}
	}
}

func (graphMatrix *GraphMatrix) DfsGraphMatirx(i int) {
	veriexFlag[i] = true
	fmt.Println(graphMatrix.Veriex[i])
	for j, _ := range graphMatrix.Veriex {
		if graphMatrix.Edge[i][j] == 1 && !veriexFlag[j] {
			graphMatrix.DfsGraphMatirx(j)
		}
	}
}

//深度优先遍历-邻接矩阵（运用栈）
func (graphMatrix *GraphMatrix) DfsTraverseMatirxStack() {
	checkGraphMatrix(graphMatrix)
	veriexFlag = make([]bool, len(graphMatrix.Veriex))
	veriexFlag[0] = true
	fmt.Println(graphMatrix.Veriex[0])
	stack := NewStackLinked()
	stack.Push(0)
	for !stack.IsEmpty() {
		i := graphMatrix.GetUnvisitedVeriex(stack.GetTop().(int))
		if i == -1 {
			stack.Pop()
		} else {
			veriexFlag[i] = true
			fmt.Println(graphMatrix.Veriex[i])
			stack.Push(i)
		}
	}
	for i, _ := range graphMatrix.Veriex {
		veriexFlag[i] = false
	}
}

func (graphMatrix *GraphMatrix) GetUnvisitedVeriex(v int) int {
	for j, _ := range graphMatrix.Veriex {
		if graphMatrix.Edge[v][j] == 1 && !veriexFlag[j] {
			return j
		}
	}
	return -1
}

//广度优先遍历-邻接矩阵（运用队列）
func (graphMatrix *GraphMatrix) BfsTraverseMatirxQueue() {
	checkGraphMatrix(graphMatrix)
	veriexFlag = make([]bool, len(graphMatrix.Veriex))
	for i, _ := range graphMatrix.Veriex {
		veriexFlag[i] = false
	}
	queue := NewQueueLinked()
	for i, _ := range graphMatrix.Veriex {
		if !veriexFlag[i] {
			veriexFlag[i] = true
			fmt.Println(graphMatrix.Veriex[i])
			queue.EnQueue(i)
			for !queue.IsEmpty() {
				i = queue.DeQueue().(*Node).Data.(int)
				for j, _ := range graphMatrix.Veriex {
					if graphMatrix.Edge[i][j] == 1 && !veriexFlag[j] {
						veriexFlag[j] = true
						fmt.Println(graphMatrix.Veriex[j])
						queue.EnQueue(j)
					}
				}
			}
		}
	}
}

//最小生成树-普鲁姆算法
func (graphMatrix *GraphMatrix) MiniSpanTree_Prim() *GraphMatrix {
	checkGraphMatrix(graphMatrix)
	vNum := len(graphMatrix.Veriex)
	lowcost := make([]int, vNum) //以i为终点的最小权值
	adjvex := make([]int, vNum)  //对应lowcost[i]的起点
	lowcost[0] = 0
	adjvex[0] = 0
	list := NewArray(vNum)
	for i := 1; i < vNum; i++ {
		lowcost[i] = int(graphMatrix.Edge[0][i])
		adjvex[i] = 0
	}
	var min, index int
	for i := 1; i < vNum; i++ {
		min = INFINITY
		var j = 1
		for j < vNum {
			if lowcost[j] != 0 && lowcost[j] < min {
				min = lowcost[j]
				index = j
			}
			j++
		}
		fmt.Println(adjvex[index], index, min)
		arr := [3]int{adjvex[index], index, min}
		list.Add(arr)
		lowcost[index] = 0
		for j = 1; j < vNum; j++ {
			if lowcost[j] != 0 && int(graphMatrix.Edge[index][j]) < lowcost[j] {
				lowcost[j] = int(graphMatrix.Edge[index][j])
				adjvex[j] = index
			}
		}
	}
	return CreateMiniSpanTreeByPrim(list, graphMatrix.Veriex...)
}

func CreateMiniSpanTreeByPrim(list *ArrayList, ver ...interface{}) *GraphMatrix {
	graph := InitDirectedGraphMatrix(ver...)
	for i := 0; i < list.length; i++ {
		arr := list.GetValue(i)
		graph.AddEdgeInDirectedGraphMatrix(arr.([3]int)[0], arr.([3]int)[1], arr.([3]int)[2])
	}
	return graph
}
