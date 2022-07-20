package main

import "fmt"

//pertama kita aliasin si function didalam parameternya dulu
type Blacklist func(string) bool

//buat function untuk filter name di register account
func registerAccount(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Welcome", name)
	} else {
		fmt.Println("You are blocked. your name is:", name)
	}
}

func main() {
	//membuat anonymous function didalam suatu variable agar lebih mudah untuk memanggilnya
	blacklist := func(name string) bool {
		return name == "Admin"
	}
	registerAccount("admin", blacklist)
	registerAccount("Antoni", blacklist)

}

// func blacklistAdmin(name string)bool{
// 	return name == "admin"
// }

// func blacklistRoot(name string)bool{
// 	return name == "root"
// }
