package main

import "fmt"

func main() {
	// var orang map[string]string = map[string]string ini versi panjangnya
	orang := map[string]string{
		"nama":            "Antoni Hasea",
		"Umur":            "22",
		"Hobby":           "Basketball",
		"Kewarganegaraan": "Indonesia",
	}
	//menambahkan key & value baru dalam map
	orang["minuman"] = "ot"

	///deklarasi keseluruhan map
	fmt.Println(orang)

	//deklarasi yang benar
	fmt.Println(orang["nama"])
	fmt.Println(orang["Umur"])
	fmt.Println(orang["Hobby"])
	fmt.Println(orang["Kewarganegaraan"])
	fmt.Println(orang["minuman"])

	//membuat map baru make format "make(map[string]string)" karena sebelumnya sudah ada map yang di deklarasi :=
	var tempat map[string]string = make(map[string]string)
	tempat["Tinggal"] = "Rumah"
	tempat["Jualan"] = "Facebook"
	tempat["ups"] = "salah"

	// print sebelum didelete
	fmt.Println(tempat)

	//function delete di map formatnya delete(map,key)
	delete(tempat, "ups")

	//print sesudah di delete
	fmt.Println(tempat)

	//panjang map
	fmt.Println(len(tempat))
}
