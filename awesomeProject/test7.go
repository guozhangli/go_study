package main

import "fmt"

func main() {
	c:=make(chan int)
	go func(){
		c<-1
		c<-2
		c<-3
		close(c)
	}()

	for v:=range c{
		fmt.Println(v)
	}

	for x:=0;x<10;x++{
		for y:=0;y<10;y++{
			if y==2{
				goto breakHere
			}
		}
	}
	return

	breakHere:
	fmt.Println("done")

	outLoop:
		for i:=0;i<5;i++{
			for j:=0;j<5;j++{
				switch j {
				case 2:
					fmt.Println(i,j)
					break outLoop
				case 3:
					fmt.Println(i,j)
					break outLoop

				}
			}
		}


	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			switch j {
			case 2:
				fmt.Println(i,j)
				continue

			}
		}
	}
}
