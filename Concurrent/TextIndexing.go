package Concurrent

import (
	"errors"
	"fmt"
	TestProject "go_study/datastructure"
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

/**
倒排索引
*/
type Document struct {
	fileName string
	wc       map[string]int
}

const PATH = "data/text"

var stop = false
var mu sync.Mutex

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
			updateInvertedIndex(&wc, &invertedIndex, f.Name())
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

//map 接收结果（并发不安全的容器）
func updateInvertedIndex(wc *map[string]int, ss *map[string]string, fileName string) {
	for k, _ := range *wc {
		if len(k) >= 3 {
			if _, ok := (*ss)[k]; ok {
				//性能太差
				/*	buf := new(TestProject.Buffer)
					buf.Write((*ss)[k],fileName, ";")
					(*ss)[k] = buf.Read()*/
				(*ss)[k] = (*ss)[k] + fileName + ";"
			} else {
				/*	buf := new(TestProject.Buffer)
					buf.Write(fileName, ";")
					(*ss)[k] = buf.Read()*/
				(*ss)[k] = fileName + ";"
			}
		}
	}
}

//sync.Map 接收结果（并发安全的容器）
func updateInvertedIndexParallel(wc *map[string]int, ss *sync.Map, fileName string) {
	for k, _ := range *wc {
		if len(k) >= 3 {
			if v, ok := (*ss).Load(k); ok {
				//性能太差
				/*buf := new(TestProject.Buffer)
				buf.Write(v.(string), fileName, ";")
				(*ss).Store(k, buf.Read())*/
				(*ss).Store(k, v.(string)+fileName+";")
			} else {
				/*buf := new(TestProject.Buffer)
				buf.Write(fileName, ";")
				(*ss).Store(k, buf.Read())*/
				(*ss).Store(k, fileName+";")
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
	start := time.Now().UnixNano()
	files := readDir()
	invertedIndex := new(sync.Map)
	lbq := TestProject.NewLinkedBlockingQueue(math.MaxInt32)
	var wgTextIndexing sync.WaitGroup
	for i := 0; i < 2; i++ {
		wgTextIndexing.Add(1)
		go execParse(invertedIndex, lbq, &wgTextIndexing)
	}
	pool := NewPoolRejectedHandler(int32(100), NewRejectedHandler(func() {
		log.Fatal("pool closed,rejected task")
	}))
	execTask(pool, files, lbq)
	fmt.Printf("worker size: %d\n", pool.WorkerSize())
	pool.ShutDown()        //将关闭协程池，不再接收任务
	pool.WaitTermination() //等待所有的任务完成，清理协程池
	setStop()
	wgTextIndexing.Wait()

	end := time.Now().UnixNano()
	fmt.Printf("Execution Time: %d\n", (end-start)/1000000)
	s := saveSlice(invertedIndex)
	fmt.Printf("invertedIndex: %d\n", len(s))
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
	start := time.Now().UnixNano()
	files := readDir()
	invertedIndex := new(sync.Map)
	lbq := TestProject.NewLinkedBlockingQueue(math.MaxInt32)
	var wgTextIndexing sync.WaitGroup

	go execParseGroup(invertedIndex, lbq, &wgTextIndexing)
	go execParseGroup(invertedIndex, lbq, &wgTextIndexing)
	pool := NewPoolRejectedHandler(int32(100), NewRejectedHandler(func() {
		log.Fatal("pool closed,rejected task")
	}))
	execTaskGroup(pool, files, lbq)
	pool.ShutDown()
	pool.WaitTermination()
	setStop()
	wgTextIndexing.Wait()
	end := time.Now().UnixNano()
	fmt.Printf("Execution Time: %d\n", (end-start)/1000000)
	s := saveSlice(invertedIndex)
	fmt.Printf("invertedIndex: %d\n", len(s))
}

func readDir() []os.FileInfo {
	files, err := ioutil.ReadDir(PATH)
	if err != nil {
		log.Fatal("open dir error,", err)
	}
	return files
}
func execParse(m *sync.Map, lbq *TestProject.LinkedBlockingQueue, wgTextIndexing *sync.WaitGroup) {
	defer wgTextIndexing.Done()
	for !stop {
		d := lbq.Take()
		if d == nil {
			continue
		}
		doc := d.(*Document)
		updateInvertedIndexParallel(&doc.wc, m, doc.fileName)
	}
	for {
		d := lbq.Poll()
		if d == nil {
			break
		}
		doc := d.(*Document)
		updateInvertedIndexParallel(&doc.wc, m, doc.fileName)
	}

}

func execParseGroup(m *sync.Map, lbq *TestProject.LinkedBlockingQueue, wgTextIndexing *sync.WaitGroup) {
	defer wgTextIndexing.Done()
	wgTextIndexing.Add(1)
	for !stop {
		d := lbq.Take()
		if d == nil {
			continue
		}
		doc := d.([]*Document)
		for _, v := range doc {
			updateInvertedIndexParallel(&v.wc, m, v.fileName)
		}
	}
	for {
		d := lbq.Poll()
		if d == nil {
			break
		}
		doc := d.([]*Document)
		for _, v := range doc {
			updateInvertedIndexParallel(&v.wc, m, v.fileName)
		}
	}
}

func execTask(pool *Pool, files []os.FileInfo, lbq *TestProject.LinkedBlockingQueue) {
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".txt") {
			task := &taskTaxtIndexing{fileName: f.Name(), lbq: lbq}
			pool.Execute(task)
		}
	}
	fmt.Printf("send task finish:%d\n", len(files))
}

func execTaskGroup(pool *Pool, files []os.FileInfo, lbq *TestProject.LinkedBlockingQueue) {
	var fileNames []string
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".txt") {
			fileNames = append(fileNames, f.Name())
			if len(fileNames) == 100 {
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
}

func saveSlice(invertedIndex *sync.Map) []string {
	var s []string
	invertedIndex.Range(func(key, value interface{}) bool {
		s = append(s, key.(string))
		return true
	})
	return s
}

func setStop() {
	defer mu.Unlock()
	mu.Lock()
	stop = true
}
