package main

import "fmt"

func main() {
	// continue adalah ketika kita mau skip kondisi tertentu yang sudah kita buat dan langsung melanjutkan ke post statement selanjutnya, berbeda dengan break yang dimana langsung mematika perulangan keseluruhan
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
