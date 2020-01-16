package main

type Dog struct {
	MaxAge int
}

func (d *Dog) Sleep() {
	panic("implement me")
}

func (d *Dog) Age() int {
	panic("implement me")
}

func (d *Dog) Type() string {
	panic("implement me")
}
