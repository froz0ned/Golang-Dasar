package main

import "fmt"

func getData() (nama string, hobby string, umur int) {
	nama = "Antoni" // assign variable dengan value
	hobby = "Basketball"
	umur = 23

	return //disini bisa langsung tulis return aja dan otomatis langsung ke return semua values dari variable yang sudah di assigned
}

func main() {
	a, b, c := getData() // variable yang disini gak harus sama kayak diatas
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
