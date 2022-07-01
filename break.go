package main

import "fmt"

func main() {
	// break adalah kondisi dimana ketika kondisi yang sudah kita buat terpenuhi maka keseluruhan pengulangan akan terhenti
	for i := 0; i < 10; i++ {
		// fmt.Println(i) jika print disini maka angka 5 nya akan sempet di print,
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

}
