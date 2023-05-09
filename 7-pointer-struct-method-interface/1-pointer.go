package main

import "fmt"

func main() {
	var nama string = "John Doe"
	fmt.Println("value nama:", nama)
	// mendapatkan alamat memory dari sebuah variabel
	fmt.Println("address nama:", &nama)

	//deklarasi variable pointer
	var pointNama *string = &nama
	fmt.Println("value pointNama", pointNama)
	//mengakses value yg tersimpan disebuah alamat memory
	fmt.Println("value dari alamat pointNama", *pointNama)
	*pointNama = "budi"
	fmt.Println("value dari alamat pointNama", *pointNama)
	fmt.Println("value nama:", nama)

	// var point = &pointNama
	// fmt.Println(point)
}
