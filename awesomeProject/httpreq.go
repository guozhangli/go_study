package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client:=&http.Client{}
	req,err:=http.NewRequest("POST","http://10.10.1.115:9080/common/exportGatherCountZx",
		strings.NewReader(`{"platform":"1007","appVersion":"1.0.3","apiVersion":"1.0.2","token":"eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIxMTExIiwic3ViIjoidG9rZW4iLCJpc3MiOiJ1c2VyIiwiaWF0IjoxNTcyNjczOTUyLCJleHAiOjE1NzI3NjAzNTJ9.yKZfGn_TUYzXAddxWfDRGzIt2T_ziZQmN4ecjMdEns8","userJGDM":"0","data":{"dateType":"day","typeCode":"1","dateBegin":"2019-10-27 00:00:00","dateEnd":"2019-11-02 23:59:59","fjjgdm":"0","queryType":"1"},"userId":"1111","userName":"admin"}`))

	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
		return
	}
	req.Header.Add("User-Agent","myclient")
	resp,err:=client.Do(req)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
		return
	}
	data,err:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	defer resp.Body.Close()
}
