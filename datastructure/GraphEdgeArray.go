package TestProject

import "fmt"

//边集数组表示图
type EdgeType struct {
	Begin  int
	End    int
	Weight int
}
type GraphEdgeArray struct {
	Veriexs []interface{}
	Edges   []*EdgeType
}

func InitGraphEdgeArray(value ...interface{}) *GraphEdgeArray {
	vNum := len(value)
	graphEdgeArray := &GraphEdgeArray{
		Veriexs: make([]interface{}, vNum),
		Edges:   make([]*EdgeType, 0),
	}
	for i, v := range value {
		graphEdgeArray.Veriexs[i] = v
	}
	return graphEdgeArray
}
func checkGraphEdgeArray(graphEdgeArray *GraphEdgeArray) {
	if graphEdgeArray == nil {
		panic("未创建图")
	}
}
func (graphEdgeArray *GraphEdgeArray) AddEdgeInGraphEA(begin int, end int, weight int) {
	checkGraphEdgeArray(graphEdgeArray)
	edge := &EdgeType{
		Begin:  begin,
		End:    end,
		Weight: weight,
	}
	graphEdgeArray.Edges = append(graphEdgeArray.Edges, edge)
}

//最小生成树-克鲁斯卡尔算法
func (graphEdgeArray *GraphEdgeArray) MiniSpanTree_Kruskal() *GraphEdgeArray {
	checkGraphEdgeArray(graphEdgeArray)
	edges := graphEdgeArray.Edges
	for i := 0; i < len(edges); i++ {
		for j := i; j < len(edges); j++ {
			if edges[i].Weight > edges[j].Weight {
				edges[i], edges[j] = edges[j], edges[i]
			}
		}
	}
	verNum := len(graphEdgeArray.Veriexs)
	parent := make([]int, verNum)
	for i := 0; i < verNum; i++ {
		parent[i] = 0
	}
	for i := 0; i < len(edges); i++ {
		m := Find(parent, edges[i].Begin)
		n := Find(parent, edges[i].End)
		if m != n {
			parent[m] = n
			fmt.Println(edges[i])
		}
	}
	return graphEdgeArray
}

func Find(parent []int, f int) int {
	for parent[f] > 0 {
		f = parent[f]
	}
	return f
}
