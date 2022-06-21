package main

import "fmt"

func main() {
	var a = 2
	var b = 19
	var c = a + b
	var x = b % a

	// augmented assignment
	x += 5 // x = x + 5

	fmt.Println(c)
	fmt.Println(x)

	// unary operator
	// x++ itu sama aja kayak x = x + 1
	// x-- itu sama aja kayak x = x - 1
	x++
	fmt.Println(x)

}
