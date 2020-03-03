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

func InitUndirectedGraphMatrix(verNum int, value ...interface{}) *GraphMatrix {
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

func InitDirectedGraphMatrix(verNum int, value ...interface{}) *GraphMatrix {
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

func (graphMatrix *GraphMatrix) DfsGraphMatirx(i int) {
	veriexFlag[i] = true
	fmt.Println(graphMatrix.Veriex[i])
	for j, _ := range graphMatrix.Veriex {
		if graphMatrix.Edge[i][j] == 1 && !veriexFlag[j] {
			graphMatrix.DfsGraphMatirx(j)
		}
	}
}

func (graphMatrix *GraphMatrix) DfsTraverse() {
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
