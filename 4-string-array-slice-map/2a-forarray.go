package main

import "fmt"

func main() {
	primes := [7]int{2, 3, 5, 8, 9, 10, 12}

	// technique 1
	for index := 0; index < len(primes); index++ {
		fmt.Println(primes[index])
	}
	// fmt.Println(primes[0])
	// fmt.Println(primes[1])
	// fmt.Println(primes[2])
	// fmt.Println(primes[3])
	// fmt.Println(primes[4])
	// fmt.Println(primes[5])

	// technique 2
	for index, element := range primes {
		fmt.Println(index, "=>", element)
	}
	for _, value := range primes {
		fmt.Println(value)
	}
	// technique 3
	index := 0
	for range primes {
		fmt.Println(primes[index])
		index++
	}
}
