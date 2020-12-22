package kMeansClustering

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

/**
k-means 聚类算法
k-means 聚类算法将预先未分类的项集分组到预定的K个簇。它在数据挖掘和机器学习领域非常流行，并且在这些领域中用于
以无监督方式组织和分类数据。
每一项通常都由一个特征向量定义。所有项都有相同数目的属性。每个簇也由一个含有同样属性数目的向量定义，这些属性
描述了所有可分类到该簇的项。该向量叫做centriod。例如，如果这些项是用数值型向量定义的，那么簇就定义为划分到该簇
的各项的平均值。
该算法基本上可以分为四个步骤：
1、初始化 在第一步中，要创建最初代表K个簇的向量，通常，你可以随机初始化这些向量。
2、指派 然后，你可以将每一项划分到一个簇中。为了选择该簇，可以计算项个每个簇之间的距离。可以使用欧式距离作为
距离度量方式，计算代表项的向量和代表簇的向量之间的距离。之后，你可以将该项分配到与其距离最短的簇中。
3、更新 一旦对所有项进行分类之后，必须重新计算定义每个簇的向量。如前所述，通常要计算划分到该簇所有项的向量的平均值。
4、结束 最后，检查是否有些项改变了为其指派的簇。如果存在变化，需要再次转入指派步骤。否则算法结束，所有项都已分类完毕。
*/

const (
	PATH1 = "../data/movies.words"
	PATH2 = "../data/movies.data"
)

type Word struct {
	index int     //该单词的索引
	tfidf float64 //其在文档中的tf-idf值
}

var K, seed int

type wordList []*Word

func (w wordList) Len() int {
	return len(w)
}
func (w wordList) Less(i, j int) bool {
	return w[i].index < w[j].index
}

func (w wordList) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

type Document struct {
	name    string
	words   wordList
	cluster *DocumentCluster
}

func (d *Document) setCluster(cluster *DocumentCluster) bool {
	if d.cluster == cluster {
		return false
	} else {
		d.cluster = cluster
		return true
	}
}

type DocumentCluster struct {
	centroid  []float64
	documents []*Document
}

func NewDocumentCluster(size int, documents []*Document) *DocumentCluster {
	return &DocumentCluster{
		centroid:  make([]float64, size),
		documents: documents,
	}
}

func (dc *DocumentCluster) addCluster(doc *Document) {
	dc.documents = append(dc.documents, doc)
}

func (dc *DocumentCluster) clearCluster() {
	dc.documents = nil
}

func (dc *DocumentCluster) calculateCentroid() {
	for _, v := range dc.documents {
		for _, w := range v.words {
			dc.centroid[w.index] += w.tfidf
		}
	}
	for i := 0; i < len(dc.centroid); i++ {
		dc.centroid[i] /= float64(len(dc.documents))
	}
}

func (dc *DocumentCluster) initialize(r *rand.Rand) {
	for i := 0; i < len(dc.centroid); i++ {
		dc.centroid[i] = r.Float64()
	}
}

type DocumentClusterList []*DocumentCluster

func (d DocumentClusterList) Len() int {
	return len(d)
}

func (d DocumentClusterList) Less(i, j int) bool {
	return d[i].getDocumentCount() > d[j].getDocumentCount()
}

func (d DocumentClusterList) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (dc *DocumentCluster) getDocumentCount() int {
	return len(dc.documents)
}

func loadfile2Map() map[string]int {
	bytes, err := ioutil.ReadFile(PATH1)
	if err != nil {
		log.Fatal("read file error")
	}
	m := make(map[string]int)
	var index int
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		if line != "" {
			m[line] = index
			index++
		}
	}
	return m
}

func loadfile2Slice(vocIndex map[string]int) []*Document {
	bytes, err := ioutil.ReadFile(PATH2)
	if err != nil {
		log.Fatal("read file error")
	}
	var docs []*Document
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		if line != "" {
			doc := processItem(line, vocIndex)
			docs = append(docs, doc)
		}
	}
	return docs
}

func processItem(line string, vocIndex map[string]int) *Document {
	tokens := strings.Split(line, ",")
	doc := &Document{
		name:  tokens[0],
		words: make(wordList, len(tokens)-1),
	}
	for i, token := range tokens {
		if i > 0 {
			wordInfo := strings.Split(token, ":")
			tfidf, _ := strconv.ParseFloat(wordInfo[1], 64)
			word := &Word{
				index: vocIndex[wordInfo[0]],
				tfidf: tfidf,
			}
			doc.words[i-1] = word
		}
	}
	sort.Sort(doc.words)
	return doc
}

func enclideanDistance(words wordList, centroid []float64) float64 {
	var distance float64
	var wordIndex int
	for i, c := range centroid {
		if wordIndex < len(words) && words[wordIndex].index == i {
			distance += math.Pow(words[wordIndex].tfidf-c, 2)
			wordIndex++
		} else {
			distance += c * c
		}
	}
	return math.Sqrt(distance)
}

func calculate(docments []*Document, clusterCount int, vocSize int, seed int64) []*DocumentCluster {
	clusters := make([]*DocumentCluster, clusterCount)
	src := rand.NewSource(seed)
	r := rand.New(src)
	for i := 0; i < clusterCount; i++ {
		var doc []*Document
		clusters[i] = NewDocumentCluster(vocSize, doc)
		clusters[i].initialize(r)
	}
	change := true
	var step int
	for change {
		change = assignment(clusters, docments)
		update(clusters)
		step++
	}
	fmt.Printf("Number of steps:%d\n", step)
	return clusters
}

func assignment(clusters []*DocumentCluster, docments []*Document) bool {
	for _, c := range clusters {
		c.clearCluster()
	}
	var numChanges int
	for _, doc := range docments {
		distance := math.MaxFloat64
		var selectCluster *DocumentCluster
		for _, c := range clusters {
			curDistance := enclideanDistance(doc.words, c.centroid)
			if distance > curDistance {
				distance = curDistance
				selectCluster = c
			}
		}
		if selectCluster != nil {
			selectCluster.addCluster(doc)
			change := doc.setCluster(selectCluster)
			if change {
				numChanges++
			}
		}
	}
	fmt.Printf("Number of change:%d\n", numChanges)
	return numChanges > 0
}

func update(clusters []*DocumentCluster) {
	for _, c := range clusters {
		c.calculateCentroid()
	}
}

func SerialKMeansClustering(ki, seedi int) {
	var K, SEED int
	if ki == 0 || seedi == 0 {
		log.Fatal("please specify K and seed")
	} else {
		K = ki
		SEED = seedi
	}
	vocIndex := loadfile2Map()
	fmt.Printf("voc size:%d\n", len(vocIndex))
	documents := loadfile2Slice(vocIndex)
	fmt.Printf("document size:%d\n", len(documents))
	start := time.Now().UnixNano()
	clusters := calculate(documents, K, len(vocIndex), int64(SEED))
	end := time.Now().UnixNano()
	fmt.Printf("K:%d,SEED:%d\n", K, SEED)
	fmt.Printf("Execution Time:%d\n", (end-start)/1000000)
	sort.Sort(DocumentClusterList(clusters))
	var msg string
	for _, v := range clusters {
		msg += "Cluster sizes:" + fmt.Sprintf("%d", v.getDocumentCount()) + ","
	}
	fmt.Println(msg)
}
