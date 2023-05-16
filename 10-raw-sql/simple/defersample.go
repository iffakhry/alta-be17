package main

import "fmt"

type Mobil struct {
	Tipe  string
	Merk  string
	Warna string
	Harga int
}

func main() {
	namaFunc()
	defer fmt.Println("satu")
	defer fmt.Println("dua")
	defer fmt.Println("tiga")
	panic("error")
	fmt.Println("empat")

	var myObj Mobil
	myObj.Tipe = "Toyota"

	var myObj1 Mobil
	myObj1.Tipe = "Suzuki"

	// var nama string
}

func namaFunc() {
	fmt.Println("func 1")
	defer fmt.Println("func 2")
	fmt.Println("func 3")
}
