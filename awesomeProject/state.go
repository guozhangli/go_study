package main

import "reflect"

type State interface {
	Name()string
	EnableSameTransit()bool
	OnBegin()
	OnEnd()
	CanTransitTo(name string)bool
}
func StateName(s State)string{
	if s==nil{
		return "none"
	}
	return reflect.TypeOf(s).Elem().Name()
}