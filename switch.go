package main

import "fmt"

func main() {
	nama := "Jonoo"

	switch nama {
	case "Jono": //case memiliki sifat seperti if
		fmt.Println("Nama saya Jono")
	case "Dicky":
		fmt.Println("Nama saya Dicky")
	default: // default memiliki sifat seperti else
		fmt.Println("Halo, Perkenalkan nama saya", nama)
	}
	// short statement dengan menggunakan switch
	switch length := len(nama); length > 5 {
	case true:
		fmt.Println("Nama cukup panjang")
	case false:
		fmt.Println("Nama kurang panjang")
	}

	//switch tanpa menggunakan kondisi di switchnya melainkan di bagian case
	panjang := len(nama)
	switch {
	case panjang > 5:
		fmt.Println("Nama terlalu panjang")
	case panjang < 5:
		fmt.Println("Nama kurang panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
