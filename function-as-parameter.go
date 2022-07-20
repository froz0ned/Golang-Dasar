package main

import (
	"fmt"
)

type Filter func(string) string // untuk membuat alias dengan type declaration terhadap function yang bisa digunakan disebuah parameter jika parameter function memiliki banyak atribut

func chatMobileLegend(name string, filter Filter) { // filter func(string) string bentuk function as parameter normal
	nameFiltered := filter(name)
	fmt.Println(nameFiltered)
}

func spamFilter(name string) string {

	if name == "anjing" {
		return "******"
	} else if name == "peler" {
		return "*****"
	} else {
		return name
	}
}

func main() {
	fmt.Println("team:")

}
