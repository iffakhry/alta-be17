package main

import "fmt"

/*
1,4,6,8,2,3,9,10,15,13,17,12,20,21,23,24,25,26

cari : 9 --> ada d index 6
cari : 11 --> tidak ada


*/

func linierSearch(elements []int, x int) int {
	// alat bantu untuk membaca tiap element nya
	// alat bantu untuk representasi index
	for i := 0; i < len(elements); i++ {
		// membandingkan data yng dicari dengan data elemen yg sedang dibaca
		if elements[i] == x {
			// jika sama, maka return indexnya
			return i
		}
	}
	//jika tidak ketemu, maka return -1. (asumsi tidak ada index -1 di slice)
	return -1
}

func main() {
	data := []int{1, 4, 6, 8, 2, 3, 9, 10, 15, 13, 17, 12, 20, 21, 23, 24, 25, 26}
	cari := 11
	index := linierSearch(data, cari)
	fmt.Println(index)
	if index < 0 {
		fmt.Println("data yg dicari tidak ada")
	} else {
		fmt.Println("data ketemu")
	}
}
