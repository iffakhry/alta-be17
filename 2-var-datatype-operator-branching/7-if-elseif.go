package main

import "fmt"

// if elseif else
func main() {
	// date := 30
	var date int
	fmt.Println("Masukkan date:")
	fmt.Scanln(&date)

	if date >= 25 {
		fmt.Println("Gajian bro")
	} else if date == 24 {
		fmt.Println("mau gajian bro")
	} else if date >= 1 && date <= 5 {
		fmt.Println("awal bulan bro")
	} else if date == 30 { // block ini tidak akan pernah dijalankan
		fmt.Println("akhir bulan bro")
	} else {
		fmt.Println("belum gajian bro")
	}

	if date == 30 {
		fmt.Println("akhir bulan cuy")
	}

	fmt.Println("selesai")
}
