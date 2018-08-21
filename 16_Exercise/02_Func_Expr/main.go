package main

import "fmt"

func main(){
	halfexp := func(number int) (float64, bool) {
		return float64(number) / 2, number%2 == 0
	}
	h, even := halfexp(20)
	fmt.Println(h, even)
}
