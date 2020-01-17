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

func TestPostfixExpresstion(t *testing.T) {
	sl := NewStackLinked()
	midExp := "9+(3-1)*3+10/2"
	postfixExp := Convert2PostfixExp(sl, midExp)
	t.Log(postfixExp)
}

func TestFindWord(t *testing.T) {
	midExp := "9+(3-1)*3+10/2"
	flag := FindWordOrNum(midExp)
	t.Log(flag)
}

func TestCalculaeExpression(t *testing.T) {
	sl := NewStackLinked()
	midExp := "(9+(3-1)*3+10/2)*5"
	//midExp := "9+(3-1)*3+10/2"
	postfixExp := Convert2PostfixExp(sl, midExp)
	t.Log(postfixExp)
	value := CalculateExpression(sl, postfixExp)
	t.Log(value)
}
