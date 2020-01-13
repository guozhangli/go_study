package main

var eventByName=make(map[string][]func(interface{}))

func RegisterEvent(name string,callback func(interface{})){
	list:=eventByName[name]
	list= append(list, callback)
	eventByName[name]=list
}

func CallEvent(name string,param interface{}){
	list:=eventByName[name]
	for _,callback:=range list{
		callback(param)
	}
}