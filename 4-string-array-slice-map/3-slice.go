package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Obtaining a slice from an array `array`
	// array[low:high]
	var primes = [5]int{2, 3, 5, 7, 11}

	// Creating a slice from the array
	var part_primes []int = primes[1:4]

	// part_primes = append(part_primes, 10000)
	// menambah data ke slice akan menambah data ke array juga

	fmt.Println(reflect.ValueOf(part_primes).Kind())
	fmt.Println(part_primes)

	// var numbers []int
	// fmt.Println(numbers)

	//long declaration
	var numbers = []int{1, 2, 3, 4, 5, 8}
	fmt.Println(numbers)
	fmt.Println(numbers[1])
	fmt.Println("len", len(numbers))

	//short declaration
	angka := []int{1, 2, 3, 4, 5, 8}
	fmt.Println(angka)

	var data1 = make([]int, 5, 10)
	fmt.Println(data1)

	bilangan := []int{1, 2, 3, 4}
	fmt.Println(bilangan)
	bilangan = append(bilangan, 9)
	bilangan = append(bilangan, 18)
	bilangan = append(bilangan, 20)
	fmt.Println(bilangan)

	var nama1 = "RUdi"
	nama := []string{"Adi", "Budi"}
	nama = append(nama, "Cindy")
	fmt.Println(nama)
	nama = append(nama, "Cindy2", "Doni")
	nama = append(nama, nama1)
	fmt.Println(nama)
	nama[2] = "cindy3"
	fmt.Println(nama)

}
