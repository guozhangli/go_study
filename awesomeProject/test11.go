package main

import (
	"fmt"
	"os"
	"sync"
)
var(
	valueByKey=make(map[string]int)
	valueByKeyGuard sync.Mutex
)
func main() {
	fmt.Println("defer begin")

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	fmt.Println("defer end")
	valueByKey["a"]=11
	valueByKey["b"]=22
	fmt.Println(readValue("a"))
	fmt.Println(fileSize(`E:\captureRecord_20191017121530219_1_2dd3fdc931bb4f5980cda446b45c6083\1-2-3-captureRecord_20191017121530219_1_2dd3fdc931bb4f5980cda446b45c6083.irisking`))
}

func readValue(key string)int{
	valueByKeyGuard.Lock()
	defer valueByKeyGuard.Unlock()
	defer fmt.Println("test defer")
	return valueByKey[key]
}

func fileSize(filename string) int64{
	f,err:=os.Open(filename)
	if err!=nil{
		return 0
	}
	defer f.Close()
	info,err:=f.Stat()
	if err!=nil{
		return 0
	}
	size:=info.Size()
	return size
}
