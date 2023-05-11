package main

import (
	"fmt"
	"sort"
)

/*
pecahan: 10, 25, 5, 1 --> 25,10,5,1
uang saya 42

42: 1,1,1,1,1,1,1  ... 42 pecahan uang
42: 5,5,5,5,5,5,5,5,1,1 --> 10 pecahan
42: 10,10,10,10,1,1 --> 6 pecahan
42: 25,5,5,5,1,1 --> 6
42: 25,1,1,1,1,1,1,1 --> banyak
42: 25,10,5,1,1 --> 5 pecahan

15: 10,5
50: 25,25
60: 25,25,10
100: 25,25,25,25
*/

func coinChange(uang int, coinValue []int) []int {
	sort.SliceStable(coinValue, func(i, j int) bool {
		return coinValue[i] > coinValue[j]
	})
	// fmt.Println("coinvalue", coinValue)
	var result []int
	// membaca koin / membaca value dari slice coinValue
	// for _, value := range coinValue {
	// 	// perulangan, ketika uangnya lebih besar dari pecahan, maka tukar dg pecahan tsb
	// 	for uang >= value {
	// 		result = append(result, value)
	// 		uang = uang - value
	// 	}
	// }

	for i := 0; i < len(coinValue); i++ {
		fmt.Println("coin", coinValue[i])
		for uang >= coinValue[i] {
			result = append(result, coinValue[i])
			uang = uang - coinValue[i]
		}
	}
	return result
}

func main() {
	coinValue := []int{10, 25, 5, 1}
	uangSaya := 60
	hasil := coinChange(uangSaya, coinValue)
	fmt.Println("hasil:", hasil)
}
