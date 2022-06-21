package main

import "fmt"

func main() {
	//type deklarasi, memeberi alias untuk tipe data
	type ktptype string
	type marriedStatus bool

	var NomorKtpGogo ktptype = "1212121398"
	var StatusPernikahan marriedStatus = false

	fmt.Println(NomorKtpGogo)
	fmt.Println(StatusPernikahan)
}
