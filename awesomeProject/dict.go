package main

import "fmt"

type Dictionary struct {
	data map[interface{}]interface{}
}

func (d *Dictionary) Get(key interface{})interface{}{
	return d.data[key]
}
func (d *Dictionary) Set(key,value interface{}){
	d.data[key]=value
}
func (d *Dictionary) Visit(callback func(k,v interface{}) bool){
	if callback==nil{
		return
	}
	for k,v:=range d.data{
		if !callback(k,v){
			return
		}
	}
}
func (d *Dictionary) Clear(){
	d.data=make(map[interface{}]interface{})
}
func NewDictionary() *Dictionary{
	d:=&Dictionary{}
	d.Clear()
	return d
}
func main() {
	dict:=NewDictionary()
	dict.Set("My Factory",60)
	dict.Set("Terra Craft",36)
	dict.Set("Don't Hungry",24)
	favorite:=dict.Get("Terra Craft")
	fmt.Println("favorite:",favorite)
	dict.Visit(func(k, v interface{}) bool {
		if v.(int)>40{
			fmt.Println(k,"is expensive")
			return true
		}
		fmt.Println(k,"is cheap")
		return true
	})
	printType(1024)
	printType("pig")
	printType(false)
}
func printType(v interface{}){
	switch v.(type) {
	case int:
		fmt.Println(v,"is int")
	case string:
		fmt.Println(v,"is string")
	case bool:
		fmt.Println(v,"is bool")
	}
}