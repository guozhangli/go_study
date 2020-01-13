package main

import "fmt"

type Alipay struct {

}
func (a *Alipay) CanUseFaceID(){

}
type Cash struct {

}
func (c *Cash) Stolen(){

}
type CantainCanUseFaceID interface {
	CanUseFaceID()
}
type ContainStolen interface {
	Stolen()
}
func print1(payMethod interface{}){
	switch payMethod.(type) {
	case CantainCanUseFaceID:
		fmt.Printf("%T can use faceid\n",payMethod)
	case ContainStolen:
		fmt.Printf("%T may be stolen\n",payMethod)
	}
}
func main() {
	print1(new(Cash))
	print1(new(Alipay))
}
