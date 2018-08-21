package main

import "fmt"

func foo(num ...int){
	fmt.Println(num)
}
func main(){
	aSlice := []int{1,2,3,4}
	foo(aSlice...)

}
