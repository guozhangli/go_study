package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var urls=[]string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"http://www.golangtc.com/",
	}

	for _,url:=range urls{
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			resp,err:=http.Get(url)
			defer resp.Body.Close()
			fmt.Println(url,err)
			str,err:=ioutil.ReadAll(resp.Body)
			fmt.Println(string(str))
		}(url)
		wg.Wait()
		fmt.Println("over")
	}
}
