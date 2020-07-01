package geneticAlgorithm

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 遗传算法是基于自然选择原理的一种自适应启发式搜索算法。用于为最优化问题和搜索问题生成优质解决方案。
/**
遗传算法为一个问题提供可能的解决方案，而该问题被称为个体或者表现型。每个个体都由一组称作染色体的属性描述。
通常，个体都由一个位序列表示，不过也可以选择更加适合具体问题的描述方法。
还需要一个适应度函数，用来确定某个方案的优劣。遗传算法的主要目标是查找一个能够使该函数最大化或者最小化的解决方案。
一旦有了初始种群，可以启动一个含有三个阶段的迭代过程。该迭代过程的每一步称作一代。每一代有如下三个阶段：
1、选择：可以在种群中选择更好的个体，这些个体在适应度函数中具有较好的值。
2、交叉：对前一步选定的个体进行交叉，以生成构成新一代的新个体。这种操作需要两个个体参与并且生成两个新的个体。实现这种操作
依赖于要解决的问题，以及所选择的个体的描述情况。
3、突变：可以应用突变运算符更改某个体的值。通常，只可以对极少量的个体执行该操作。
满足结束标准前，可以重复以上操作。
*/

var PATH = [...]string{"../data/kn57_dist.txt", "../data/lau15_dist.txt"}

const (
	generations = 10000
	individuals = 1000
)

func load(path string) [][]int {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("read file error")
	}
	lines := strings.Split(string(bytes), "\r\n")
	var datas [][]int
	for _, line := range lines {
		if line != "" {
			cells := strings.Split(line, ",")
			var celli []int
			for _, cell := range cells {
				i, _ := strconv.Atoi(cell)
				celli = append(celli, i)
			}
			datas = append(datas, celli)
		}
	}
	return datas
}

type Individual struct {
	chromosomes []int
	value       int
}

func NewIndividual(size int) *Individual {
	return &Individual{
		chromosomes: make([]int, size),
		value:       0,
	}
}

type IndividualList []*Individual

func (il IndividualList) Len() int {
	return len(il)
}

func (il IndividualList) Less(i, j int) bool {
	return il[i].value < il[j].value
}

func (il IndividualList) Swap(i, j int) {
	il[i], il[j] = il[j], il[i]
}

func (il *IndividualList) initIndividual(size int) {
	for i := 0; i < individuals; i++ {
		individual := NewIndividual(size)
		for s := 0; s < size; s++ {
			individual.chromosomes[s] = s
		}
		ch, _ := ShuffleIntegers(individual.chromosomes)
		individual.chromosomes = ch
		*il = append(*il, individual)
	}
}

func (il *IndividualList) evaluate(datas [][]int, size int) {
	for _, v := range *il {
		evaluate(v, datas)
	}
	sort.Sort(il)
}

func evaluate(individual *Individual, datas [][]int) {
	chromosomes := individual.chromosomes
	size := len(chromosomes)
	var total int
	for i := 0; i < size-1; i++ {
		start := chromosomes[i]
		end := chromosomes[i+1]
		total += datas[start][end]
	}
	start := chromosomes[size-1]
	end := chromosomes[0]
	total += datas[start][end]
	individual.value = total
}

func (il *IndividualList) selection() *IndividualList {
	var selection = new(IndividualList)
	len := len(*il) / 2
	for i := 0; i < len; i++ {
		*selection = append(*selection, (*il)[i])
	}
	return selection
}

func (il *IndividualList) crossover(selected *IndividualList, size int) []*Individual {
	var population = make([]*Individual, individuals)
	rand.Seed(time.Now().UnixNano())
	len := len(*selected)
	for i := 0; i < individuals/2; i++ {
		population[2*i] = NewIndividual(size)
		population[2*i+1] = NewIndividual(size)
		index1 := rand.Intn(len)
		index2 := rand.Intn(len)
		for index1 == index2 {
			index2 = rand.Intn(len)
		}
		individual := (*selected)[index1]
		individua2 := (*selected)[index2]
		crossover(individual, individua2, population[2*i], population[2*i+1], size)
	}
	return population
}

func crossover(parent1, parent2, p1, p2 *Individual, size int) {
	rand.Seed(time.Now().UnixNano())
	number1 := rand.Intn(size - 1)
	number2 := rand.Intn(size - 1)
	for number1 == number2 {
		number2 = rand.Intn(size - 1)
	}
	start := Min(number1, number2)
	end := Max(number1, number2)
	ch1 := parent1.chromosomes[start:end]
	ch2 := parent2.chromosomes[start:end]
	for i := 0; i < size; i++ {
		v := (end + i) % size
		v1 := parent1.chromosomes[v]
		v2 := parent2.chromosomes[v]
		addValue(ch1, v2)
		addValue(ch2, v1)
	}
	ch1 = Rotate(ch1, start)
	ch2 = Rotate(ch2, start)
	p1.chromosomes = ch1
	p2.chromosomes = ch2
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func addValue(ch []int, v int) {
	var flag = true
	for _, c := range ch {
		if c == v {
			flag = false
		}
	}
	if flag {
		ch = append(ch, v)
	}
}

func calculate(datas [][]int) *Individual {
	var population = new(IndividualList)
	size := len(datas)
	population.initIndividual(size)
	population.evaluate(datas, size)
	best := (*population)[0]
	for i := 0; i < generations; i++ {
		sl := population.selection()
		*population = population.crossover(sl, size)
		population.evaluate(datas, size)
		if (*population)[0].value < best.value {
			best = (*population)[0]
		}
	}
	return best
}

func SerialGeneticAlgorithm() {
	for _, path := range PATH {
		start := time.Now().UnixNano()
		datas := load(path)
		result := calculate(datas)
		end := time.Now().UnixNano()
		fmt.Println("=======================================")
		fmt.Printf("Example:%s\n", path)
		fmt.Printf("Generations: %d\n", generations)
		fmt.Printf("Population:  %d\n", individuals)
		fmt.Printf("Execution Time:  %d\n", (end-start)/1000000)
		fmt.Printf("Best Individual: %v\n", result)
		fmt.Printf("Total Distance: %d\n", result.value)
		fmt.Println("=======================================")
	}
}
