package main

func main() {
	c := make(chan struct{})
	<-c
}
