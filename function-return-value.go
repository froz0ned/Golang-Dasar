package main

import "fmt"

func getValue(nama string) string {
	if nama == "" {
		return "Hello Brodie"
	} else {
		return "Hello " + nama
	}
}

func main() {

	postValue := getValue("Gogo") // bisa menggunakan variable
	fmt.Println(postValue)
	fmt.Println(getValue("")) // atau langsung print functionnya dan masukin value parameternya

}
