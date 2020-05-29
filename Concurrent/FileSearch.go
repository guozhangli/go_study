package Concurrent

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Result struct {
	Path  string
	Found bool
}

func NewResult() *Result {
	return &Result{
		Path:  "",
		Found: false,
	}
}

func FileSearchSerial(filePath, fileName string, result *Result) {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		return
	}
	if len(files) == 0 || files == nil {
		return
	}
	for _, entry := range files {
		if entry.IsDir() {
			subFile := filepath.Join(filePath, entry.Name())
			FileSearchSerial(subFile, fileName, result)
		} else {
			if entry.Name() == fileName {
				result.Path = filepath.Join(filePath, entry.Name())
				result.Found = true
				return
			}
		}
		if result.Found {
			return
		}
	}
}

func FileSearchParallel(filePath, fileName string, result *Result) {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		return
	}
	if len(files) == 0 || files == nil {
		return
	}
	ch := make(chan string, len(files))
	done := make(chan struct{})
	go func(filePath string, files []os.FileInfo) {
		for _, entry := range files {
			subPath := filepath.Join(filePath, entry.Name())
			ch <- subPath
		}
	}(filePath, files)
	go func(fileName string, ch chan string, result *Result, done chan struct{}) {
		for {
			select {
			case filePath, ok := <-ch:
				if !ok {
					return
				}
				walkdir(filePath, fileName, result, done)
			case <-done:
				return
				fmt.Println("exit chan")
			}
		}
	}(fileName, ch, result, done)
	<-done
}

func walkdir(filePath, fileName string, result *Result, done chan struct{}) {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		return
	}
	if len(files) == 0 || files == nil {
		return
	}
	for _, entry := range files {
		if entry.IsDir() {
			subFile := filepath.Join(filePath, entry.Name())
			walkdir(subFile, fileName, result, done)
		} else {
			if entry.Name() == fileName {
				result.Path = filepath.Join(filePath, entry.Name())
				result.Found = true
				done <- struct{}{}
				return
			}
		}
		if result.Found {
			return
		}
	}
}
