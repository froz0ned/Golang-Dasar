package main

import "fmt"

// recursive adalah function yang memanggil functionnya sendiri
func recursiveLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

func recursiveFunction(angka int) int {
	if angka == 1 {
		return 1
	} else {
		return angka * recursiveFunction(angka-1)
	}
}

func main() {
	loop := recursiveLoop(5)
	fmt.Println(loop)
	fmt.Println(recursiveFunction(3))
}
