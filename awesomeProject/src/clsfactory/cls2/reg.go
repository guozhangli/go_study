package cls2

import (
	"fmt"
	"go_study/awesomeProject/src/clsfactory/base"
)

type Class2 struct {
}

func (c *Class2) Do() {
	fmt.Println("Class2")
}

func init() {
	base.Register("Class2", func() base.Class {
		return new(Class2)
	})
}
