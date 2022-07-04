package main

import "fmt"

func getNama(nama string) string {
	return "Hello " + nama
}

func main() {
	dataYangDikirimKeParameter := getNama("Gogo") // sebuah function yang dipakai sebagai value dari sebuah variable
	fmt.Println(dataYangDikirimKeParameter)
}
