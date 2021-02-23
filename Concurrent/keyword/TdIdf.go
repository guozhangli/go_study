package keyword

import (
	"math"
	"strings"
)

//词频
func td(list []string, word string) float64 {
	var n int
	for _, v := range list {
		if strings.ToLower(v) == word {
			n++
		}
	}
	return float64(n) / float64(len(list))
}

//文档频率
func df(collections [][]string, word string) int {
	if word == "" {
		return 0
	}
	var n int
	for _, v := range collections {
		for _, v1 := range v {
			if strings.ToLower(v1) == word {
				n++
				break
			}
		}
	}
	return n
}

//逆文档频率
func idf(collections [][]string, word string) float64 {
	return math.Log(float64(len(collections)) / float64(df(collections, word)+1))
}

func tdidf(list []string, collections [][]string, word string) float64 {
	return td(list, word) * idf(collections, word)
}
