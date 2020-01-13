package main

import (
	"fmt"
	"sort"
)

type Interface interface {
	Len() int
	Less(i,j int) bool
	Swap(i,j int)
}
type MyStringList []string

func (m MyStringList) Swap(i, j int) {
	m[i],m[j]=m[j],m[i]
}

func (m MyStringList) Len()int{
	return len(m)
}
func (m MyStringList) Less(i,j int)bool{
	return m[i]<m[j]
}

func main() {
	names:=MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	fmt.Println(names.Len())
	sort.Sort(names)
	for _,v:= range  names{
		fmt.Printf("%s\n",v)
	}

}
