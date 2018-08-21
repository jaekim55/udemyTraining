package main

import "fmt"

func half (number int) (float64, bool){
	return float64(number)/2, number%2==0
}

func main() {
	num, even := half(8)
	fmt.Println(num, even)
}

