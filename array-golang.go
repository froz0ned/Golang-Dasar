package main

import "fmt"

func main() {
	//dibawah ini adalah cara panjang untuk deklarasi array

	// var names [4]string
	// names[0] = "Antoni"
	// names[1] = "Hasea"
	// names[2] = "Trigogo"
	// names[3] = "Aritonang"

	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])
	// fmt.Println(names[3])

	//dibawah ini adalah cara cepet deklarasi array

	var values = [4]string{ // pakai cara ini index tetep start dari 0 tapi dia kasih reserved 1 index kosong
		"Antoni",
		"Hasea",
		"Trigogo",
		"Aritonang",
	}
	values[1] = "greek"      // Function untuk mengubah array index yang kita mau disini saya ubah index ke 1
	fmt.Println(len(values)) //Function untuk menghitung panjang array yang sudah kita buat diatas. walaupun tidak ada datanya tapi jika arraynya sudah di buat untuk menyimpan 5 array maka akan tercetak 5 sebagai panjang array tsb

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[3])
}
