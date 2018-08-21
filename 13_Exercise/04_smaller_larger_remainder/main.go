package main

import "fmt"

func main() {
	var smaller int
	var larger int
	fmt.Println("You will need to enter two numbers of different values")
	fmt.Print("Please enter the smaller number: ")
	fmt.Scan(&smaller)
	fmt.Print("Please enter the larger number: ")
	fmt.Scan(&larger)
	fmt.Println("The remainder of", larger, "divided by", smaller, "is", larger%smaller)
}
