package Concurrent

import (
	"errors"
	"fmt"
	TestProject "go_study/datastructure"
	"io/ioutil"
	"log"
	"math"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Document struct {
	fileName string
	wc       map[string]int
}

var stop = false
var wgTextIndexing sync.WaitGroup

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
			wc, err := parse(filepath.Join(PATH, f.Name()))
			if err != nil {
				continue
			}
			updateInvertedIndex(wc, invertedIndex, f.Name())
		}
	}
	end := time.Now().UnixNano()
	fmt.Printf("Execution Time: %d\n", (end-start)/1000000)
	fmt.Printf("invertedIndex: %d\n", len(invertedIndex))
}

var reg = regexp.MustCompile("\\P{L}+")

func parse(filePath string) (map[string]int, error) {
	var wc = make(map[string]int)
	datas, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("read file error,", err)
		return nil, errors.New("read file error")
	}
	data := reg.Split(string(datas), -1)
	for _, s := range data {
		if s != "" {
			s = strings.ToLower(s)
			if _, ok := wc[s]; !ok {
				wc[s] = 1
			} else {
				wc[s]++
			}
		}
	}
	return wc, nil
}

func updateInvertedIndex(wc map[string]int, ss map[string]string, fileName string) {
	for k, _ := range wc {
		if len(k) >= 3 {
			if _, ok := ss[k]; ok {
				buf := new(TestProject.Buffer)
				buf.Write(fileName, ";")
				ss[k] += buf.Read()
			} else {
				buf := new(TestProject.Buffer)
				buf.Write(fileName, ";")
				ss[k] = buf.Read()
			}
		}
	}
}

func updateInvertedIndexParallel(wc map[string]int, ss *sync.Map, fileName string) {
	for k, _ := range wc {
		if len(k) >= 3 {
			if v, ok := ss.Load(k); ok {
				buf := new(TestProject.Buffer)
				buf.Write(v.(string), fileName, ";")
				ss.Store(k, buf.Read())
			} else {
				buf := new(TestProject.Buffer)
				buf.Write(fileName, ";")
				ss.Store(k, buf.Read())
			}
		}
	}
}

type taskTaxtIndexing struct {
	fileName string
	lbq      *TestProject.LinkedBlockingQueue
}

func (task *taskTaxtIndexing) Run() error {
	wc, err := parse(filepath.Join(PATH, task.fileName))
	if err != nil {
		return errors.New("error")
	}
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
	var invertedIndex = new(sync.Map)
	rejected := NewRejectedHandler(func() {
		log.Fatal("pool closed,rejected task")
	})
	lbq := TestProject.NewLinkedBlockingQueue(math.MaxInt32)
	go func(m *sync.Map) {
		wgTextIndexing.Add(1)
		for !stop {
			doc := lbq.Take().(*Document)
			updateInvertedIndexParallel(doc.wc, m, doc.fileName)
		}
		wgTextIndexing.Done()
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
	stop = true
	wgTextIndexing.Wait()
	end := time.Now().UnixNano()
	fmt.Printf("Execution Time: %d\n", (end-start)/1000000)

	var len int32
	invertedIndex.Range(func(key, value interface{}) bool {
		atomic.AddInt32(&len, 1)
		return true
	})
	fmt.Printf("invertedIndex: %d\n", len)
}

type taskIndexingGroup struct {
	fileNames []string
	lbq       *TestProject.LinkedBlockingQueue
}

func (task *taskIndexingGroup) Run() error {
	var docs []*Document
	for _, v := range task.fileNames {
		wc, err := parse(filepath.Join(PATH, v))
		if err != nil {
			return errors.New("error")
		}
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
	var invertedIndex = new(sync.Map)
	rejected := NewRejectedHandler(func() {
		log.Fatal("pool closed,rejected task")
	})
	lbq := TestProject.NewLinkedBlockingQueue(math.MaxInt32)
	go func(m *sync.Map) {
		wgTextIndexing.Add(1)
		for !stop {
			doc := lbq.Take().([]*Document)
			for _, v := range doc {
				updateInvertedIndexParallel(v.wc, m, v.fileName)
			}
		}
		wgTextIndexing.Done()
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
	stop = true
	wgTextIndexing.Wait()
	end := time.Now().UnixNano()
	fmt.Printf("Execution Time: %d\n", (end-start)/1000000)
	var len int32
	invertedIndex.Range(func(key, value interface{}) bool {
		atomic.AddInt32(&len, 1)
		return true
	})
	fmt.Printf("invertedIndex: %d\n", len)
}
