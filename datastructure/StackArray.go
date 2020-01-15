package TestProject

type StackArray struct {
	data []interface{}
	top  int
}

func NewStackArray(cap uint) *StackArray {
	stackArray := &StackArray{
		data: make([]interface{}, 0, cap),
		top:  -1,
	}
	return stackArray
}

func checkStack(stack *StackArray) {
	if stack == nil {
		panic("no stack created")
	}
}

func (stack *StackArray) IsEmpty() bool {
	checkStack(stack)
	if stack.top == -1 {
		return true
	}
	return false
}

func (stack *StackArray) Length() int {
	checkStack(stack)
	return len(stack.data)
}

func (stack *StackArray) Push(value interface{}) {
	checkStack(stack)
	stack.data = append(stack.data, value)
	stack.top++
}

func (stack *StackArray) Pop() interface{} {
	checkStack(stack)
	if stack.top == -1 {
		return nil
	}
	val := stack.data[stack.top]
	stack.data = append(stack.data[:stack.top], stack.data[stack.top+1:]...)
	stack.top--
	return val
}

func (stack *StackArray) Clear() {
	checkStack(stack)
	if stack.Length() > 0 {
		stack.data = stack.data[0:0]
		stack.top = -1
	}
}

func (stack *StackArray) Distroy(s **StackArray) {
	checkStack(stack)
	*s = nil
}
