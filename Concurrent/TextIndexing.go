package Concurrent

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
		line, is_prefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF && !is_prefix {
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
