package filesearch

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
	for _, entry := range files {
		go func(filePath string, entry os.FileInfo) {

			subPath := filepath.Join(filePath, entry.Name())
			ch <- subPath

		}(filePath, entry)
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
	}
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

func CopyFileMain(filePath string) {
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
			CopyFileMain(subFile)
		} else {
			fileSrc := filepath.Join(filePath, entry.Name())
			fileDst := filepath.Join("D://20", entry.Name())
			CopyFile(fileSrc, fileDst)
		}
	}
}

func CopyFile(srcFileName string, dstFileName string) {
	//打开源文件
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		log.Fatalf("源文件读取失败,原因是:%v\n", err)
	}
	defer func() {
		err = srcFile.Close()
		if err != nil {
			log.Fatalf("源文件关闭失败,原因是:%v\n", err)
		}
	}()

	//创建目标文件,稍后会向这个目标文件写入拷贝内容
	distFile, err := os.Create(dstFileName)
	if err != nil {
		log.Fatalf("目标文件创建失败,原因是:%v\n", err)
	}
	defer func() {
		err = distFile.Close()
		if err != nil {
			log.Fatalf("目标文件关闭失败,原因是:%v\n", err)
		}
	}()
	//定义指定长度的字节切片,每次最多读取指定长度
	var tmp = make([]byte, 1024*4)
	//循环读取并写入
	for {
		n, err := srcFile.Read(tmp)
		n, _ = distFile.Write(tmp[:n])
		if err != nil {
			if err == io.EOF { //读到了文件末尾,并且写入完毕,任务完成返回(关闭文件的操作由defer来完成)
				return
			} else {
				log.Fatalf("拷贝过程中发生错误,错误原因为:%v\n", err)
			}
		}
	}
}
