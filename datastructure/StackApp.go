package TestProject

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	Plus = 1
	Sub  = 1
	Mult = 2
	Div  = 2
)

func fabi(i int) int {
	if i < 2 {
		if i == 0 {
			return 0
		}
		return 1
	}
	return fabi(i-1) + fabi(i-2)
}

func Convert2PostfixExp(stack Stack, midExp string) string {
	var postfxExp string
	w := FindWordOrNum(midExp)
	for _, v := range w {
		var s = midExp[v[0]:v[1]]
		if isNumber(s) {
			PostfixExp(&postfxExp, s)
		} else {
			if s != ")" {
				p := GetPriority(s)
				top := stack.GetTop()
				if top != nil {
					tp := GetPriority(top.(string))
					if s == "(" {
						stack.Push(s)
					} else {
						if p > tp {
							stack.Push(s)
						} else {
							var top = stack.GetTop()
							for top != nil && top.(string) != "(" {
								var t = stack.Pop()
								v := reflect.ValueOf(t)
								if !v.IsNil() {
									top := t.(*Node).Data
									if top.(string) != "(" {
										PostfixExp(&postfxExp, top.(string))
									}
								} else {
									top = nil
								}
							}
							stack.Push(s)
						}
					}
				} else {
					stack.Push(s)
				}
			} else {
				var top = stack.GetTop()
				for top != nil && top.(string) != "(" {
					var t = stack.Pop()
					v := reflect.ValueOf(t)
					if !v.IsNil() {
						top := t.(*Node).Data
						if top.(string) != "(" {
							PostfixExp(&postfxExp, top.(string))
						}
					} else {
						top = nil
					}
				}
			}
		}
	}
	if !stack.IsEmpty() {
		var top = stack.GetTop()
		for top != nil {
			var t = stack.Pop()
			v := reflect.ValueOf(t)
			if !v.IsNil() {
				top := t.(*Node).Data
				PostfixExp(&postfxExp, top.(string))
			} else {
				top = nil
			}
		}
	}
	return postfxExp
}

func PostfixExp(postfxExp *string, s string) {
	if *postfxExp != "" {
		*postfxExp += " "
	}
	*postfxExp += s
}

func CalculateExpression(stack Stack, postfixExp string) int {
	str := strings.Split(postfixExp, " ")
	for _, v := range str {
		if isNumber(v) {
			stack.Push(v)
		} else {
			b := stack.Pop().(*Node).Data.(string)
			a := stack.Pop().(*Node).Data.(string)
			a_i, _ := strconv.Atoi(a)
			b_i, _ := strconv.Atoi(b)
			c_i := CalculateValue(a_i, b_i, v)
			c := strconv.Itoa(c_i)
			stack.Push(c)
		}
	}
	if !stack.IsEmpty() {
		res := stack.GetTop().(string)
		res_i, _ := strconv.Atoi(res)
		return res_i
	}
	return 0
}

func isNumber(s string) bool {
	reg := "\\d+"
	regexp := regexp.MustCompile(reg)
	result := regexp.Match([]byte(s))
	return result
}

func FindWordOrNum(s string) [][]int {
	reg := "[-+*/()]|\\d+"
	regexp := regexp.MustCompile(reg)
	result := regexp.FindAllStringIndex(s, -1)
	return result
}

func GetPriority(s string) int {
	switch s {
	case "+":
		return Plus
	case "-":
		return Sub
	case "*":
		return Mult
	case "/":
		return Div
	default:
		return 0
	}
}

func CalculateValue(a int, b int, s string) int {
	switch s {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}
