package main

import "fmt"

func main() {
	var nilai32 int32 = 130
	var nilai8 int8 = int8(nilai32)
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai8)
	fmt.Println(nilai64)

	//konversi dari tipe data byte ke string ketika mencari huruf menggunakan string array index di string.go
	var name = "Antoni"            //deklarasi variable name yang berisi string antoni
	var e byte = name[0]           //deklarasi variable e untuk mencari index yang pertama yaitu [0] dari variable name. golang array start from 0. hasil yang di return itu tipe datanya byte(uint8).
	var eString string = string(e) // deklarasi variable eString yang dimana berfungsi untuk mengubah pencarian index nama. yg tadinya akan mengeluarkan return byte untuk huruf a tapi di konversi kembali ke string.

	fmt.Println(name)
	fmt.Println(eString)
}
