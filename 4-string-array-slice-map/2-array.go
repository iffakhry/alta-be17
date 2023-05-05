package main

import "fmt"

func main() {
	var countries [5]string
	countries[0] = "Indonesia"
	countries[1] = "Malaysia"
	countries[2] = "Singapura"
	countries[3] = "Thailand"
	countries[4] = "Jepang"
	countries[3] = "Amerika"

	fmt.Println("countries:", countries)
	fmt.Println("panjang countries:", len(countries))

	var angka [10]int
	angka[0] = 1
	angka[1] = 2
	angka[9] = 10
	fmt.Println("angka:", angka)
	fmt.Println("len angka:", len(angka))

	odd_numbers := [5]int{1, 3, 5, 7, 9}      // Intialized with values
	var even_numbers [5]int = [5]int{0, 2, 4} // Partial assignment

	fmt.Println(odd_numbers)
	fmt.Println(even_numbers)

	primesnumber := [...]int{2, 3, 3, 4, 6, 8, 9, 10}
	fmt.Println(primesnumber)
	fmt.Println(len(primesnumber))

	kota := [8]string{0: "Surabaya", 2: "Malang"}
	fmt.Println(kota)
	number := [10]int{0: 20, 3: 12, 9: 77}
	fmt.Println(number[3])

}
