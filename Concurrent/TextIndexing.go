package Concurrent

import (
	"bufio"
	"fmt"
	TestProject "go_study/datastructure"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Document struct {
	fileName string
	wc       map[string]int
}

const PATH = "data/text"

func TextIndexingSerial() {
	files, err := ioutil.ReadDir(PATH)
	if err != nil {
		log.Fatal("open dir error,", err)
	}
	start := time.Now().UnixNano()
	invertedIndex := make(map[string]string)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".txt") {
			wc := parse(filepath.Join(PATH, f.Name()))
			updateInvertedIndex(wc, invertedIndex, f.Name())
		}
	}
	end := time.Now().UnixNano()
	fmt.Printf("Execution Time: %d\n", (end-start)/1000000)
	fmt.Printf("invertedIndex: %d\n", len(invertedIndex))
}

func parse(filePath string) map[string]int {
	var wc = make(map[string]int)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("open file error")
	}
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		strs := regex(string(line))
		for _, s := range strs {
			if s != "" {
				if _, ok := wc[s]; !ok {
					wc[s] = 1
				} else {
					wc[s]++
				}
			}
		}
	}
	return wc
}

func regex(line string) []string {
	reg := regexp.MustCompile("\\P{L}+")
	flag := reg.Split(line, -1)
	return flag
}

func updateInvertedIndex(wc map[string]int, ss map[string]string, fileName string) {
	for k, _ := range wc {
		if len(k) > 3 {
			if _, ok := ss[k]; ok {
				ss[k] = fmt.Sprintf("%s%s%s", ss[k], fileName, ";")
			} else {
				ss[k] = fmt.Sprintf("%s%s", fileName, ";")
			}
		}
	}
}

func updateInvertedIndexParallel(wc map[string]int, ss sync.Map, fileName string) {
	for k, _ := range wc {
		if len(k) > 3 {
			if v, ok := ss.Load(k); ok {
				ss.Store(k, fmt.Sprintf("%s%s%s", v, fileName, ";"))
			} else {
				ss.Store(k, fmt.Sprintf("%s%s", fileName, ";"))
			}
		}
	}
}

type taskTaxtIndexing struct {
	fileName string
	lbq      *TestProject.LinkedBlockingQueue
}

func (task *taskTaxtIndexing) Run() error {
	wc := parse(filepath.Join(PATH, task.fileName))
	doc := &Document{
		fileName: task.fileName,
		wc:       wc,
	}
	task.lbq.Put(doc)
	return nil
}

func TextIndexingParallel() {
	files, err := ioutil.ReadDir(PATH)
	if err != nil {
		log.Fatal("open dir error,", err)
	}
	start := time.Now().UnixNano()
	var invertedIndex sync.Map
	rejected := NewRejectedHandler(func() {
		log.Fatal("pool closed,rejected task")
	})
	lbq := TestProject.NewLinkedBlockingQueue(math.MaxInt32)
	go func(m sync.Map) {
		for {
			doc := lbq.Take().(*Document)
			updateInvertedIndexParallel(doc.wc, m, doc.fileName)
		}
	}(invertedIndex)
	pool := NewPoolRejectedHandler(int32(100), rejected)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".txt") {
			task := &taskTaxtIndexing{fileName: f.Name(), lbq: lbq}
			pool.Execute(task)
		}
	}
	pool.ShutDown()
	pool.WaitTermination()
	end := time.Now().UnixNano()
	fmt.Printf("Execution Time: %d\n", (end-start)/1000000)
	fmt.Printf("invertedIndex: %v\n", invertedIndex)
}

type taskIndexingGroup struct {
	fileNames []string
	lbq       *TestProject.LinkedBlockingQueue
}

func (task *taskIndexingGroup) Run() error {
	var docs []*Document
	for _, v := range task.fileNames {
		wc := parse(filepath.Join(PATH, v))
		doc := &Document{
			fileName: v,
			wc:       wc,
		}
		docs = append(docs, doc)
	}
	task.lbq.Put(docs)
	return nil
}

func TextIndexingGroup() {
	files, err := ioutil.ReadDir(PATH)
	if err != nil {
		log.Fatal("open dir error,", err)
	}
	start := time.Now().UnixNano()
	var invertedIndex sync.Map
	rejected := NewRejectedHandler(func() {
		log.Fatal("pool closed,rejected task")
	})
	lbq := TestProject.NewLinkedBlockingQueue(math.MaxInt32)
	go func(m sync.Map) {
		for {
			doc := lbq.Take().([]*Document)
			for _, v := range doc {
				updateInvertedIndexParallel(v.wc, m, v.fileName)
			}
		}
	}(invertedIndex)
	pool := NewPoolRejectedHandler(int32(100), rejected)
	var fileNames []string
	for i, f := range files {
		if strings.HasSuffix(f.Name(), ".txt") {
			if (i+1)%100 != 0 {
				fileNames = append(fileNames, f.Name())
			} else {
				task := &taskIndexingGroup{fileNames: fileNames, lbq: lbq}
				pool.Execute(task)
				fileNames = nil
			}

		}
	}
	if len(fileNames) > 0 {
		task := &taskIndexingGroup{fileNames: fileNames, lbq: lbq}
		pool.Execute(task)
		fileNames = nil
	}
	pool.ShutDown()
	pool.WaitTermination()
	end := time.Now().UnixNano()
	fmt.Printf("Execution Time: %d\n", (end-start)/1000000)
	fmt.Printf("invertedIndex: %v\n", invertedIndex)
}
