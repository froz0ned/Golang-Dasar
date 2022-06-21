package main

import "fmt"

func main() {
	var bulan = [12]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = bulan[5:10] // 5 sebagai nilai low start dari slice, dan 10 sebagai value high dari slice. high value ini hanya sebagai pembatas slice value terhitung dari low sampai high - 1
	// fmt.Println(slice1)

	fmt.Println(len(slice1))
	// function mencari panjang dari slice1

	fmt.Println(cap(slice1))
	// function menghitung capacity dari sebuah slice, dihitung dari low sampai akhir dari array.

	// *mengubah array akan berpengaruh juga ke slice !! BERHATI-HATI DALAM MENGUBAH VALUE DARI ARRAY !!
	// bulan[5] = "Gemini"
	// fmt.Println(slice1)

	// *mengubah value dari slice dengan dimulai dari low slice yaitu 0 dimana 0 itu adalah index ke 6 dari array
	// slice1[0] = "Inimeg"
	// fmt.Println(bulan)

	var slice2 = bulan[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "gogo")
	fmt.Println(slice3)

	// membuat slice langsung tanpa array lain, yang dimana arraynya langsung terbuat didalam sebuah slice
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Gogo"
	newSlice[1] = "Keren"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice beserta atribut dan valuenya
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// perbedaan slice dgn array
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
