package Concurrent

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"time"
)

/**
Levenshtein Distance 算法，又叫 Edit Distance 算法，是指两个字符串之间，由一个转成另一个所需的最少编辑操作次数。
许可的编辑操作包括将一个字符替换成另一个字符，插入一个字符，删除一个字符。一般来说，编辑距离越小，两个串的相似度越大。
计算相似度
*/
func Levenshtein(str1, str2 string) float64 {
	by1 := []byte(str1)
	by2 := []byte(str2)
	len1 := len(by1) + 1
	len2 := len(by2) + 1
	arr := make([][]int, len1)
	for i := 0; i < len1; i++ {
		arr[i] = make([]int, len2)
		arr[i][0] = i
	}
	for j := 0; j < len2; j++ {
		arr[0][j] = j
	}
	fmt.Println(arr)
	for j := 1; j < len2; j++ {
		for i := 1; i < len1; i++ {
			top := arr[i-1][j] + 1
			left := arr[i][j-1] + 1
			left_top := arr[i-1][j-1]
			if by1[i-1] == by2[j-1] {
				left_top = left_top + 0
			} else {
				left_top = left_top + 1
			}
			arr[i][j] = min(left_top, min(top, left))
		}
	}

	fmt.Printf("%f\n", float64(arr[len1-1][len2-1]))
	//计算相似度
	var result float64
	result = float64(arr[len1-1][len2-1]) / float64(max(len1-1, len2-1))
	return float64(1 - result)
}

func min(a, b int) int {
	if a > b {
		return b
	} else if a < b {
		return a
	} else {
		return a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else if a < b {
		return b
	} else {
		return a
	}
}

//计算差距
func LevenshteinDistance(str1, str2 string) int {
	len1 := len(str1) + 1
	len2 := len(str2) + 1
	arr := make([][]int, len1)
	for i := 0; i < len1; i++ {
		arr[i] = make([]int, len2)
		arr[i][0] = i
	}
	for j := 0; j < len2; j++ {
		arr[0][j] = j
	}
	for j := 1; j < len2; j++ {
		for i := 1; i < len1; i++ {
			top := arr[i-1][j] + 1
			left := arr[i][j-1] + 1
			left_top := arr[i-1][j-1]
			if str1[i-1] == str2[j-1] {
				left_top = left_top + 0
			} else {
				left_top = left_top + 1
			}
			arr[i][j] = min(left_top, min(top, left))
		}
	}
	return arr[len1-1][len2-1]
}

type BestMatchingData struct {
	distance int
	words    []string
}

func NewBestMatchingData(distance int, words []string) *BestMatchingData {
	bestMatchingData := new(BestMatchingData)
	bestMatchingData.words = words
	bestMatchingData.distance = distance
	return bestMatchingData
}

func load(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Print("open file is error")
	}
	var dic []string
	bf := bufio.NewReader(file)
	for {
		line, _, err := bf.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		dic = append(dic, string(line))
	}
	return dic
}

func getBestMatchingWords(word string, dictionary []string) *BestMatchingData {
	min := math.MaxInt32
	var result []string
	for _, d := range dictionary {
		distance := LevenshteinDistance(word, d)
		if distance < min {
			min = distance
			result = nil
			result = append(result, d)
		} else if distance == min {
			result = append(result, d)
		}
	}
	bestMatchingData := NewBestMatchingData(min, result)
	return bestMatchingData
}

//最佳匹配算法的串行版本
func MatchingData() {
	dictionary := load("data/UK Advanced Cryptics Dictionary.txt")
	fmt.Println("Dictionary Size: ", len(dictionary))
	startTime := time.Now().UnixNano()
	word := "stitter"
	bestMatchingData := getBestMatchingWords(word, dictionary)
	results := bestMatchingData.words
	endTime := time.Now().UnixNano()
	fmt.Printf("Word: %s\n", word)
	fmt.Printf("Minimun distance: %d\n", bestMatchingData.distance)
	fmt.Printf("List of best matching words: %d\n", len(results))
	for _, r := range results {
		fmt.Println(r)
	}
	fmt.Printf("Execution Time: %dms\n", (endTime-startTime)/1000000)
}

func getBestMatchingWordsParallel(word string, dictionary []string, startIndex, endIndex int) *BestMatchingData {
	min := math.MaxInt32
	var result []string
	for i := startIndex; i < endIndex; i++ {
		distance := LevenshteinDistance(word, dictionary[i])
		if distance < min {
			min = distance
			result = nil
			result = append(result, dictionary[i])
		} else if distance == min {
			result = append(result, dictionary[i])
		}
	}
	bestMatchingData := NewBestMatchingData(min, result)
	return bestMatchingData
}

//最佳匹配算法的并行版本
func MatchingDataParallel() {
	dictionary := load("data/UK Advanced Cryptics Dictionary.txt")
	fmt.Println("Dictionary Size: ", len(dictionary))
	startTime := time.Now().UnixNano()
	word := "stitter"
	rejected := NewRejectedHandler(func() {
		log.Fatal("pool closed,rejected task")

	})
	var poolNum = int32(100)
	pool := NewPoolRejectedHandler(poolNum, rejected)
	step := len(dictionary) / int(poolNum)
	startIndex := 0
	endIndex := step
	ch := make(chan *BestMatchingData)
	for i := 0; i < step; i++ {
		pool.Execute(func() error {
			bestMatchingData := getBestMatchingWordsParallel(word, dictionary, startIndex, endIndex)
			ch <- bestMatchingData
			return nil
		})
		startIndex = endIndex
		if i < step-2 {
			endIndex = endIndex + step
		} else {
			endIndex = len(dictionary)
		}
	}
	var min = math.MaxInt32
	var results []string
	for v := range ch {
		if v.distance < min {
			min = v.distance
			results = nil
			for _, w := range v.words {
				results = append(results, w)
			}

		} else if v.distance == min {
			for _, w := range v.words {
				results = append(results, w)
			}
		}
	}
	endTime := time.Now().UnixNano()
	fmt.Printf("Word: %s\n", word)
	fmt.Printf("Minimun distance: %d\n", min)
	fmt.Printf("List of best matching words: %d\n", len(results))
	for _, r := range results {
		fmt.Println(r)
	}
	fmt.Printf("Execution Time: %dms\n", (endTime-startTime)/1000000)
}
