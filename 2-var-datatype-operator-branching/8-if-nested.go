package main

import "fmt"

func main() {
	date := 32
	if date < 1 || date > 31 {
		fmt.Println("input salah")
	} else {
		fmt.Println("inputan valid")

		if date >= 25 {
			fmt.Println("Gajian bro")
			if date == 25 {
				fmt.Println("Gajian tepat waktu")
			}
		} else if date == 24 {
			fmt.Println("mau gajian bro")
		} else if date >= 1 && date <= 5 {
			fmt.Println("awal bulan bro")
		} else if date == 30 { // block ini tidak akan pernah dijalankan
			fmt.Println("akhir bulan bro")
		} else {
			fmt.Println("belum gajian bro")
		}
	}

	var nama string = "budi"
	if nama == "budi" {
		fmt.Println("halo guys")
	} else {
		fmt.Println("ok")
	}
}
