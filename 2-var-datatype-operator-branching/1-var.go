package main

import "fmt"

func main() {
	// long declaration
	var age uint
	fmt.Println(age) // 0

	var negara string
	fmt.Println(negara)

	var bilangan1 int = 15
	fmt.Println(bilangan1) // 15

	var bilangan2 float32 = 10.5
	fmt.Println(bilangan2)

	//deklarasi variabel sekaligus memberikan value budi
	var firstName string = "Budi"
	fmt.Println("nama depan:", firstName)

	var status bool = true
	fmt.Print("statusnya: ", status, "\n")

	var nama1, nama2 string = "Budi 1", "Budi 2"
	fmt.Println(nama1, "dan", nama2)

	//short declaration
	//deklarasi variabel sekaligus memberikan value 10
	angka1 := 10
	fmt.Println(angka1)
	lastName := "Budii 3"
	fmt.Println(lastName)

	//deklarasi variabel dengan nama nilai1 yang bertipe int
	var nilai1 int
	fmt.Println(nilai1) // 0
	// memasukkan/memberikan value 90 ke variabel nilai1
	nilai1 = 90
	// memasukkan/memberikan (assign ulang) value 100 ke variabel nilai1
	nilai1 = 100
	fmt.Println("nilai 1 a: ", nilai1)

	//constanta
	const phi float64 = 3.14
	// phi = 100
	fmt.Println("phi", phi)
	const message string = "hello BE 17"
	// message = "halo gaes"
	fmt.Println(message)
}
