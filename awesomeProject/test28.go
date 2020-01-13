package main

import (
	"fmt"
	"sort"
)

type HeroKind int
const(
	None HeroKind=iota
	Tank
	Assassin
	Mage
)
type Hero struct {
	Name string
	Kind HeroKind
}
type Heros []*Hero
func(s Heros)Len()int{
	return len(s)
}
func (s Heros) Less(i,j int)bool{
	if s[i].Kind!=s[j].Kind{
		return s[i].Kind<s[j].Kind
	}
	return s[i].Name<s[j].Name
}
func (s Heros) Swap(i,j int){
	s[i],s[j]=s[j],s[i]
}
func main() {
	heros:=Heros{
		&Hero{"吕布",Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}
	sort.Sort(heros)
	for _,v:=range heros{
		fmt.Printf("%+v\n",v)
	}

	heros2:=[]*Hero{
		{"吕布",Tank},
		{"李白", Assassin},
		{"妲己", Mage},
		{"貂蝉", Assassin},
		{"关羽", Tank},
		{"诸葛亮", Mage},
	}
	sort.Slice(heros2, func(i, j int) bool {
		if heros2[i].Kind!=heros2[j].Kind{
			return heros2[i].Kind<heros2[j].Kind
		}
		return heros2[i].Name<heros2[j].Name
	})
	for _,v:=range heros2{
		fmt.Printf("%+v\n",v)
	}
}
