package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:9080/essearch/identity/gatherRecord",
		strings.NewReader(` {"platform":"1007","appVersion":"1.0.3","apiVersion":"1.0.2","token":"666666","data":{"userJGDM":"110000000000","queryType":"2","isDeleted":2,"sjsjStart":"2020-01-01 00:00:00","sjsjEnd":"2020-12-31 23:59:59","hmcjCylx":"3"},"page":"1","limit":"1000"}`))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	//req.Header.Add("User-Agent","myclient")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	d := m["data"].(map[string]interface{})["data"].([]interface{})
	fmt.Printf("es 查询记录数量:%d\n", len(d))
	defer resp.Body.Close()

	req2, err := http.NewRequest("POST", "http://localhost:9080/bigScreen/queryDoubtfulRecords",
		strings.NewReader(`{
    "page": 1,
    "limit": 1000,
    "data": {
        "fssjStart": "2020-01-01 00:00:00",
        "fssjEnd": "2020-12-31 23:59:59",
        "zjlx": "",
        "zjhm": "",
        "xm": "",
        "ryfl": "",
        "cylx": "3",
        "userJGDM": "110000000000",
        "queryType": "2",
        "includesub": true
    },
    "platform": "",
    "appversion": "1.0.3",
    "apiversion": "1.0.2",
    "mac": "12345678",
    "userName": "hyy",
    "userId": "12",
    "ip": "10.10.1.66 ",
    "token": "666666"
}`))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	resp2, err := client.Do(req2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	data2, err := ioutil.ReadAll(resp2.Body)
	var m2 map[string]interface{}
	err = json.Unmarshal(data2, &m2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	d2 := m2["data"].(map[string]interface{})["records"].([]interface{})
	fmt.Printf("mysql 查询记录数量:%d\n", len(d2))
	defer resp2.Body.Close()
	fmt.Println("es compare mysql")
	for _, v := range d {
		vv := v.(map[string]interface{})
		if v1, ok := vv["cjsbbh"]; ok {
			var flag bool
			for _, v2 := range d2 {
				vv2 := v2.(map[string]interface{})
				if v3, ok := vv2["cjsbbh"]; ok {
					if v1 == v3 {
						flag = true
					}
				}
			}
			if !flag {
				fmt.Printf("未找到cjsbbh：%s\n", v1)
			}
		}
	}
	fmt.Println("mysql compare es")
	for _, v := range d2 {
		vv := v.(map[string]interface{})
		if v1, ok := vv["cjsbbh"]; ok {
			var flag bool
			for _, v2 := range d {
				vv2 := v2.(map[string]interface{})
				if v3, ok := vv2["cjsbbh"]; ok {
					if v1 == v3 {
						flag = true
					}
				}
			}
			if !flag {
				fmt.Printf("未找到cjsbbh：%s\n", v1)
			}
		}
	}
}
