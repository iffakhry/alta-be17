package main

import (
	"fmt"
	"math"
)

/*
sekolah skolah

1 buku tulis 10000
2 pulpen 5000
3 Tas 50000

buku tulis 2
pulpen 4

total yang harus dia bayar?
buku tulis 2 * 10000 = 20000
pulpen 4 * 5000 = 20000
total nya = 40000

bilangan prima adalah bil yang hanya bisa dibagi oleh dirinya sendiri dan 1

sample:
2,3,5,7,11 13,17

input 5
output true

10
output false

15
output false
*/

// func checkPrime(number int) bool {
// 	// var result bool = true

// 	if number == 2 {
// 		return true
// 	} else if number%2 == 0 {
// 		return false
// 	} else if number <= 1 {
// 		return false
// 	}

// 	for i := 3; i < number; i = i + 2 {
// 		if number%i == 0 {
// 			return false
// 		}
// 	}
// 	return true

// }

// func checkPrime(number int) bool {
// 	if number < 2 {
// 		return false
// 	}
// 	for i := 2; i <= number/2; i++ {
// 		if number%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func checkPrime(number int) bool {
	if number < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrtNumber; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	result := checkPrime(9)
	fmt.Println("result", result)
}
