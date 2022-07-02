package main

import "fmt"

func getProfile() (string, string, string) { // note: jika ingin mengambil value yg di tulis adalah tipe datanya, jika ada 2 data yang ingin di return harus menggunakan ()
	return "Antoni", "adalah seorang", "Hacker"
}

func main() {
	variable1, _, variable3 := getProfile()
	fmt.Println(variable1)
	// fmt.Println(variable2) // jika tidak ingin mengirim value dari variable2 maka tinggal ubah saja jadi _
	fmt.Println(variable3)
}
