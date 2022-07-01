package main

import "fmt"

func main() {
	laper := "tio"
	if laper == "makan" {
		fmt.Println("kenyang")
	} else if laper == "ngemil" {
		fmt.Println("lumayan kenyang")
	} else {
		fmt.Println("laper")
	}

	// var panjangString = len(laper) => deklarasi manual menggunakan variable baru untuk mengecek panjang dari string laper
	//           **short statement** untuk membuat pengecekan panjang string didalam if kondisi ini saja       //
	if panjangString := len(laper); panjangString > 5 {
		fmt.Println("Nama kepanjangan")
	} else {
		fmt.Println("Panjang nama sudah pas")
	}
}
