package main

import (
	"fmt"
)

func main(){
	n :=large(43, 46, 87, 45, 126)
	fmt.Println(n)
}

func large(l ...int)int{
	var biggest int
	for _,v:=range l{
		if v>biggest {
			biggest = v
		}
	}
	return biggest
}