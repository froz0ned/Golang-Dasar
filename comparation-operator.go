package main

import "fmt"

func main() {
	var name1 = "Gogo"
	var name2 = "Gogo"

	var result bool = name1 == name2
	fmt.Println("ini string :", result)

	var angka1 = 100
	var angka2 = 200

	fmt.Println("angka1 > angka2 :", angka1 > angka2)
	fmt.Println("angka1 < angka2 :", angka1 < angka2)
	fmt.Println("angka1 == angka2 :", angka1 == angka2)
	fmt.Println("angka1 != angka2 :", angka1 != angka2)

}
