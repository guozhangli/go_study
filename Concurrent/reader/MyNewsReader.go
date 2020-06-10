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
)

const (
	source = "../data/sources.txt"
	output = "../data/output/"
	stoped = false
)

func NewsReader() {
	nb := NewNewsBuffer()
	go NewsWriter(nb)
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
		pool.Execute(func() error {

			return nil
		})

	}

}

func NewsWriter(nb *NewsBuffer) {
	for !stoped {
		item := nb.get()
		err := ioutil.WriteFile(output+item.getFileName(), []byte(item.toXML()), os.ModePerm)
		if err != nil {
			fmt.Errorf("%s", "write file to disk error")
		}
	}
}

func NewsTask(name, url string, buffer *NewsBuffer) {
	rss := NewRssDataCapturer(name, buffer)
	list := rss.load(url)
	length := list.Length()
	for i := 0; i < length; i++ {
		buffer.add(list.GetVaule(i).(*CommonInformationItem))
	}
}
