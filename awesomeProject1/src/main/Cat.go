package main

type Cat struct {
	MaxAge int
}

func (c *Cat) Sleep() {
	panic("implement me")
}

func (c *Cat) Age() int {
	panic("implement me")
}

func (c *Cat) Type() string {
	panic("implement me")
}
