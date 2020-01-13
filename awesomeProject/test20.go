package main

import (
	"encoding/json"
	"fmt"
)

type Screen struct {
	Size float32
	ResX,ResY int
}
type Battery struct {
	Capactiy int
}
func genJsonData()[]byte{
	raw:=&struct {
		Screen
		Battery
		HasTouchID bool
	}{
		Screen:Screen{
			Size:5.5,
			ResX:1920,
			ResY:1080,
		},
		Battery:Battery{
			2910,
		},
		HasTouchID:true,
	}
	jsonData,err:=json.Marshal(raw)
	if err!=nil{
		fmt.Printf("json parse error")
		return nil
	}
	return jsonData

}
func main() {
	jsonData:=genJsonData()
	fmt.Println(string(jsonData))

	screenAndTouch:=struct{
		Screen
		HasTouchID bool
	}{}
	json.Unmarshal(jsonData,&screenAndTouch)
	fmt.Printf("%+v\n",screenAndTouch)

	batteryAndTouch:= struct {
		Battery
		HasTouchID bool
	}{}
	json.Unmarshal(jsonData,&batteryAndTouch)
	fmt.Printf("%+v\n",batteryAndTouch)
}
