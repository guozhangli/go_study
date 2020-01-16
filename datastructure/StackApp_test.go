package TestProject

import "testing"


func TestFabiFunc(t *testing.T) {
	i := 0
	for i < 6 {
		n := fabi(i)
		i++
		t.Log(n)
	}
}


//9+(3-1)*3+10/2

func TestPostfixExpresstion(t *testing.T){
	sl:=NewStackLinked()
	midExp:="9+(3-1)*3+10/2"
	postfixExp:=Convert2PostfixExp(sl,midExp)
	t.Log(postfixExp)
}


