package reader

import (
	"bufio"
	"fmt"
	"go_study/Concurrent"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	source = "../data/sources.txt"
	output = "../data/output/"
)

type MyNewsReader struct {
	wg     sync.WaitGroup
	stoped bool
}

type task struct {
	reader *MyNewsReader
	data   []string
	nb     *NewsBuffer
}

func (t task) Run() error {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for range ticker.C {
			t.reader.newsTask(t.data[0], t.data[1], t.nb)
		}
	}()
	return nil
}

func NewMyNewsReader() *MyNewsReader {
	reader := new(MyNewsReader)
	reader.stoped = false
	return reader
}

func (reader *MyNewsReader) NewsReader() {
	reader.wg.Add(1)
	nb := NewNewsBuffer()
	go reader.newsWriter(nb)
	rejected := Concurrent.NewRejectedHandler(func() {
		log.Fatal("goroutine is closed,reject exectue task")
	})
	pool := Concurrent.NewPoolRejectedHandler(20, rejected)
	file, err := os.Open(source)
	if err != nil {
		log.Fatal("open file is error")
	}
	rd := bufio.NewReader(file)
	for {
		line, _, err := rd.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		data := strings.Split(string(line), ";")
		task := task{
			reader: reader,
			data:   data,
			nb:     nb,
		}
		pool.Execute(task)
	}
	reader.wg.Wait()
	log.Println("Shutting down the executor.")
	reader.stoped = true
	pool.ShutDown()
	log.Println("The system has finished.")
}

func (reader *MyNewsReader) newsWriter(nb *NewsBuffer) {
	for !reader.stoped {
		if item := nb.get(); item != nil {
			err := ioutil.WriteFile(output+item.getFileName(), []byte(item.toXML()), os.ModePerm)
			if err != nil {
				fmt.Errorf("%s", "write file to disk error")
			}
		}
	}
}

func (reader *MyNewsReader) newsTask(name, url string, buffer *NewsBuffer) {
	rss := NewRssDataCapturer(name, buffer)
	list, err := rss.load(url)
	if err != nil {
		log.Println("read error form url:", url)
		return
	}
	for list.HasNode() {
		if list.Data() != nil {
			buffer.add(list.Data().(*CommonInformationItem))
		}
		list = list.NextNode()
	}
}

func (reader *MyNewsReader) shutdown() {
	reader.wg.Done()
}

func MyNewsReaderStart() {
	myNewsReader := NewMyNewsReader()
	go myNewsReader.NewsReader()
	time.Sleep(5 * time.Minute)
	myNewsReader.shutdown()
}
