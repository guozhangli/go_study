package main

import "fmt"

type Weapon int
const(
	Arrow Weapon=iota
	Shuriken
	SniperFifle
	Rifle
	Blower
)
func main() {
	fmt.Println(Arrow,Shuriken,SniperFifle,Rifle,Blower)
}
