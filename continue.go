package main

import "fmt"

func main() {
	// continue adalah ketika kita mau skip kondisi tertentu yang sudah kita buat dan langsung melanjutkan ke post statement selanjutnya, berbeda dengan break yang dimana langsung mematika perulangan keseluruhan
	// buat lah bilangan genap dari 1 sampai 10
	for i := 1; i <= 10; i++ {
		if i%2 == 1 { // ketika hasil bagi dari i adalah 1 atau ganjil maka result dari i tidak akan di print
			continue
		}
		fmt.Println(i)
	}
	// buat bilangan ganjil dari 1 sampai 10
	for x := 1; x <= 10; x++ {
		if x%2 == 0 { // ketika x adalah genap maka akan terskip dan tidak di print
			continue
		}
		fmt.Println(x)
	}

}
