package main

import "fmt"

/*
input: sebuah deret bilangan
outputnya: munculkan angka yang genap saja
*/

func CetakGenap(input []int) {
	// baca deret bilangannya
	//cek tiap bilangan
	// for _, value := range input {
	// 	// fmt.Println("z:", z)
	// 	// fmt.Println(input[z])
	// 	if value%2 == 0 {
	// 		fmt.Println(value)
	// 	}
	// }
	// for z := 0; z < len(input); z++ {
	for z := len(input) - 1; z >= 0; z-- {
		// fmt.Println("z:",z)
		// fmt.Println(input[z])
		if input[z]%2 == 0 {
			fmt.Println(input[z])
		}
	}
	// fmt.Println(input[0])
	// fmt.Println(input[1])
	// fmt.Println(input[2])
	// fmt.Println(input[3])
	// fmt.Println(input[4])
	// fmt.Println(input[5])
	// fmt.Println(input[6])
}

func main() {
	var data = []int{10, 2, 9, 4, 11, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	CetakGenap(data)
	// ketika tidak ada angka genap sama sekali, maka munculkan pesan "tidak ada bilangan genap"
}
