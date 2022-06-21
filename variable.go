package main

import "fmt"

func main() {
	var name string

	name = "Antoni Aritonang"
	fmt.Println(name)

	name = "Antoni Hasea"
	fmt.Println(name)

	var city = "Kota Depok"
	fmt.Println(city)

	city = "Kota Jakarta"
	fmt.Println(city)

	var age = 23
	fmt.Println(age)

	hobby := "Basketball"
	fmt.Println(hobby)

	hobby = "Beatbox"
	fmt.Println(hobby)

	var (
		firstname = "Antoni"
		lastname  = "Hasea"
	)
	fmt.Println(firstname)
	fmt.Println(lastname)

	//const itu sama kayak var sifat dari const adalah ketika di dideklarasi dia tidak dipanggil tidak apa-apa,
	// dan const adalah variable yang tidak bisa diubah ketika valuenya sudah di deklarasi
	const (
		nickname           = "Gogo"
		umur               = 23
		basketballposition = "Small Forward"
	)
	fmt.Println(nickname)
	fmt.Println(umur)
	fmt.Println(basketballposition)
}
