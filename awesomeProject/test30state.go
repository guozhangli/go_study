package main

import "fmt"

type IdleState struct {
	StateInfo
}
func (i *IdleState) OnBegin(){
	fmt.Println("IdleState begin")
}
func (i *IdleState) OnEnd(){
	fmt.Println("IdleState end")
}
type MoveState struct {
	StateInfo
}
func (m *MoveState) OnBegin(){
	fmt.Println("MoveState begin")
}

func (m *MoveState) EnableSameTransit() bool{
	return true
}

type JumpState struct {
	StateInfo
}
func (j *JumpState) OnBegin(){
	fmt.Println("JumpState begin")
}
func (j *JumpState) CanTransitTo(name string)bool{
	return name!="MoveState"
}

func main() {
	sm:=NewStateManager()
	sm.OnChange=func(from,to State){
		fmt.Printf("%s ------->%s\n\n",StateName(from),StateName(to))
	}
	sm.Add(new(IdleState))
	sm.Add(new(MoveState))
	sm.Add(new(JumpState))
	transitAndReport(sm,"IdleState")
	transitAndReport(sm,"MoveState")
	transitAndReport(sm,"MoveState")
	transitAndReport(sm,"JumpState")
	transitAndReport(sm,"JumpState")
	transitAndReport(sm,"IdleState")

}

func transitAndReport(sm *StateManager,target string){
	if err:=sm.Transit(target);err!=nil{
		fmt.Printf("FAILED ! %s------->%s,%s\n\n",sm.CurrState().Name(),target,err.Error())
	}
}
