package main

import (
	"fmt"
	"sort"
)

func main() {
	var team[3] string
	team[0]="hammer"
	team[1]="soldier"
	team[2]="mum"
	fmt.Println(team)

	var team2=[3]string{"hammer","soldier","mum"}
	fmt.Println(team2)

	var team3=[...]string{"hammer","soldier","mum"}
	fmt.Println(team3)

	for k,v:=range team3{
		fmt.Println(k,v)
	}
	//切片 地址，大小，容量

	var highRiseBuilding [30]int
	for i:=0;i<30;i++{
		highRiseBuilding[i]=i+1
	}
	fmt.Println(highRiseBuilding[10:15])
	fmt.Println(highRiseBuilding[15:])
	fmt.Println(highRiseBuilding[:15])
	fmt.Println(highRiseBuilding[0:0])

	var strList []string
	var numList []int
	var numListEmpty=[]int{}
	fmt.Println(strList,numList,numListEmpty)
	fmt.Println(len(strList),len(numList),len(numListEmpty))
	fmt.Println(strList==nil)
	fmt.Println(numList==nil)
	fmt.Println(numListEmpty==nil)
	//make
	a:=make([]int,2)//内存分配
	b:=make([]int,2,10)
	fmt.Println(a,b)
	fmt.Println(len(a),len(b))
	//append
	var numbers []int
	for i:=0;i<10;i++{
		numbers=append(numbers,i)
		fmt.Printf("len: %d cap: %d pointer: %p\n",len(numbers),cap(numbers),numbers)
	}

	var car []string
	car=append(car,"OldDirver")
	car=append(car,"Ice","Sniper","Monk")

	team4:=[]string{"Pig","Flyingcake","Chicken"}
	car=append(car,team4...)
	fmt.Println(car)
	//copy
	const elementCount=1000
	srcData:=make([]int,elementCount)
	for i:=0;i<elementCount;i++{
		srcData[i]=i
	}
	refData:=srcData
	copyData:=make([]int,elementCount)
	copy(copyData,srcData)
	srcData[0]=999
	fmt.Println(refData[0])
	fmt.Println(copyData[0],copyData[elementCount-1])
	copy(copyData,srcData[4:6])
	for i:=0;i<5;i++{
		fmt.Printf("%d\n",copyData[i])
	}
    //切片删除
	seq:=[]string{"a","b","c","d","e"}
	index:=2
	fmt.Println(seq[:index],seq[index+1:])
	seq=append(seq[:index],seq[index+1:]...)
	fmt.Println(seq)

	//map
	scene:=make(map[string]int)
	scene["route"]=66
	fmt.Println(scene["route"])
	v:=scene["route2"]
	fmt.Println(v)
	v,ok:=scene["route"]
	fmt.Println(ok)

	m:=map[string]string{
		"W":"forward",
		"A":"left",
		"D":"right",
		"S":"backward",
	}
	fmt.Println(m)

	scene["brazil"]=4
	scene["china"]=960
	for k,v:=range scene{
		fmt.Println(k,v)
	}
	for k,v:=range m{
		fmt.Println(k,v)
	}
	for _,v:=range m{
		fmt.Println(v)
	}
	for k:=range m{
		fmt.Println(k)
	}
	for _,v:=range scene{
		fmt.Println(v)
	}
	for k:=range scene{
		fmt.Println(k)
	}
	var sceneList []string
	for k:=range scene{
		sceneList=append(sceneList,k)
	}
	sort.Strings(sceneList)
	fmt.Println(sceneList)

	delete(scene,"brazil")
	for k,v:=range scene{
		fmt.Println(k,v)
	}

	m1:=make(map[int]int)
	go func(){
		for{
			m1[1]=1
		}
	}()
	go func() {
		for{
			_=m1[1]
		}
	}()
	var i int
	for{
		i++
	}
}
