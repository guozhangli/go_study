package main

import "awesomeProject/src/clsfactory/base"
import _ "awesomeProject/src/clsfactory/cls1"
import _ "awesomeProject/src/clsfactory/cls2"
func main(){
	c1:=base.Create("Class1")
	c1.Do()

	c2:=base.Create("Class2")
	c2.Do()
}