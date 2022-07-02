package main

import "fmt"

func sayhello(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}
func main() {
	sayhello("Antoni", "Hasea")
}
