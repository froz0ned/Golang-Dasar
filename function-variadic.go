package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10, 10, 10) // tipe data yang default akan terkirim oleh varargs adalah slice
	fmt.Println(total)

	//ketika slice sudah ada didalam suatu function yang sudah kita buat, slice masih akan
	slice := []int{1, 1, 1, 1, 1}
	slicePassing = sumAll(slice...) // khusus passing slice ke varargs harus menggunakan tambahan ...
	fmt.Println(slice)
}
