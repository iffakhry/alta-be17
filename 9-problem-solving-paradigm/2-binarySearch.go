package main

import "fmt"

func linierSearch(elements []int, x int) int {
	var counter = 0
	// alat bantu untuk membaca tiap element nya
	// alat bantu untuk representasi index
	for i := 0; i < len(elements); i++ {
		counter++
		// membandingkan data yng dicari dengan data elemen yg sedang dibaca
		if elements[i] == x {
			fmt.Println("count:", counter)
			// jika sama, maka return indexnya
			return i
		}
	}
	//jika tidak ketemu, maka return -1. (asumsi tidak ada index -1 di slice)
	return -1
}

func binarySearch(input []int, cari int) int {
	var kiri = 0
	var kanan = len(input) - 1
	var counter = 0 //var pembantu, untuk menghitung brp kali perulangan berjalan

	for kiri <= kanan {
		counter++
		// mencari index tengah
		var tengah = (kiri + kanan) / 2
		// membandingkan angka cari dengan value di index tengah
		if cari < input[tengah] {
			// jika lebih kecil, maka geser kanannya
			kanan = tengah - 1
		} else if cari > input[tengah] {
			kiri = tengah + 1
		} else if cari == input[tengah] {
			fmt.Println("count:", counter)
			return tengah
		}
	}
	fmt.Println("count:", counter)
	return -1

}

func main() {
	var data = []int{12, 15, 15, 19, 24, 31, 53, 59, 60, 75, 89, 90, 100, 102, 104, 106, 109}
	indexFound := binarySearch(data, 109)
	// indexFound := linierSearch(data, 109)
	fmt.Println("index:", indexFound)
}
