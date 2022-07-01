package main

import "fmt"

func main() {
	a := 1
	for a <= 10 {
		fmt.Println("Perulangan ke", a)
		a++
	}

	// for dengan statement
	for b := 1; b <= 5; b++ {
		fmt.Println("Perulangan ke - ", b)
	}

	slice := []string{
		"Antoni",
		"Hasea",
		"Trigogo",
		"Aritonang",
	}
	// for dengan statement implementasi terhadap slice
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range implementasi terhadap  slice
	// for index(untuk tipe data map menggunakan key), name(ini adalah value) := range (bisa di isi dengan slice,map,array) {print}
	for i, value := range slice {
		fmt.Println("index", i, "=", value) // note : jika ingin print valuenya saja, index nya harus diubah menjadi _ karena di golang tidak bolehh ada variable yang tidak terpakai
	}

	//for range untuk map

	orang := make(map[string]string)
	orang["name"] = "Antoni"
	orang["title"] = "Programmer"

	for key, value := range orang {
		fmt.Println(key, "=", value)
	}
}
