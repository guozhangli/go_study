package keyword

import (
	"fmt"
	"go_study/datastructure"
	"math"
)

//余弦相似度理论

func calCosSim(v1, v2 map[string]float64) float64 {
	var sclar, norm1, norm2, similarity float64
	set1 := TestProject.NewSet(func(o, n interface{}) bool {
		return o == n
	})
	set2 := TestProject.NewSet(func(o, n interface{}) bool {
		return o == n
	})
	for k, _ := range v1 {
		set1.Insert(k)
	}
	for k, _ := range v2 {
		set2.Insert(k)
	}
	set3 := TestProject.NewSet(func(o, n interface{}) bool {
		return o == n
	})
	TestProject.Intersection(set1, set2, set3)
	l := set3.Iterator()
	for l.HasNode() {
		fmt.Printf("%v\n", l.Data())
		l = l.NextNode()
	}

	l2 := set3.Iterator()
	for l2.HasNode() {
		sclar += v1[l2.Data().(string)] * v2[l2.Data().(string)]
		l2 = l2.NextNode()
	}
	for _, v := range v1 {
		norm1 += math.Pow(v, 2)
	}
	for _, v := range v2 {
		norm2 += math.Pow(v, 2)
	}
	similarity = sclar / math.Sqrt(norm1*norm2)
	fmt.Printf("%f\n", sclar)
	fmt.Printf("%f\n", norm1)
	fmt.Printf("%f\n", norm2)
	fmt.Printf("%f\n", similarity)
	return 0.0
}
