package main

import "fmt"

// Operasi && (and)
//--------------------
// true && true = true
// true && false = false
// false && true = false
// false && false = false
//-----------------------

// Operasi || (or)
//----------------------
// false || false = false
// true || false = true
// true || true = true
// false || true = true

// Operasi ! (negasi) hanya untuk satu operator boolean klo != itu untuk 2 operator != itu adalah (bukan samadengan)
//-----------------
// !true = false
// !false = true
//------------------

func main() {

	var nilaiAkhir = 90   // deklarasi nilai siswanya
	var absensiTotal = 30 // deklarasi jumlah absensi siswanya

	// cara manual
	// var kriteriaLulusNilaiAkhir = nilaiAkhir > 80
	// var kriteriaLulusAbsensi = absensiTotal > 25

	// var SyaratSiswaLulus bool = kriteriaLulusNilaiAkhir && kriteriaLulusAbsensi

	// fmt.Println(SyaratSiswaLulus)

	fmt.Println(nilaiAkhir > 80 && absensiTotal > 25)

}
