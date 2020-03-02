package TestProject

//边集数组表示图
type EdgeType struct {
	Begin  int
	End    int
	Weight int
}
type GraphEdgeArray struct {
	Veriexs []interface{}
	Edges   []*Edge
}

func InitGraphEdgeArray(value ...interface{}) *GraphEdgeArray {
	vNum := len(value)
	graphEdgeArray := &GraphEdgeArray{
		Veriexs: make([]interface{}, vNum),
		Edges:   make([]*Edge, vNum*(vNum-1)/2),
	}
	return graphEdgeArray
}
