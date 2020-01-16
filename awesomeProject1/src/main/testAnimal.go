package main

import "log"

func Factory(name string) Animal {
	switch name {
	case "dog":
		return &Dog{MaxAge: 20}
	case "cat":
		return &Cat{MaxAge: 10}
	default:
		panic("No such animal")
	}
}

func main() {
	animal := Factory("dog")
	log.Fatal(animal)
}
