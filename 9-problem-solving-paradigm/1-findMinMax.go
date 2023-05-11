package main

import "fmt"

/*
10, 7, 3, 5, 8, 2,12, 9

min = 10
max = 10
*/

func findMinMax(input []int) (min int, max int, minIndex int, maxIndex int) {
	min = input[0]
	max = input[0]
	minIndex, maxIndex = 0, 0

	// proses membaca per data nya
	for i := 1; i < len(input); i++ {
		// membandingkan elemen input ke i dengan min
		if input[i] < min {
			//jika value input ke i lebih kecil dari min, mkaa nilai min diganti
			min = input[i]
			minIndex = i
		}
		if input[i] > max {
			max = input[i]
			maxIndex = i
		}
	}

	return min, max, minIndex, maxIndex
}

func main() {
	var angka = []int{10, 7, 3, 5, 8, 2, 12, 9}
	min, max, minIndex, maxIndex := findMinMax(angka)
	fmt.Println("min", min, "max", max, "minIndex", minIndex, "maxIndex", maxIndex)
}
