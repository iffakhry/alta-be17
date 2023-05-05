package main

import "fmt"

func main() {
	// N := 5
	// for i := 0; i < N; i++ {
	// 	for j := 0; j < N; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// for baris := 0; baris < 4; baris++ { // 4
	// 	for kolom := 0; kolom < 4; kolom++ { // 3
	// 		fmt.Printf("b: %d, k: %d ||", baris, kolom)
	// 	}
	// 	fmt.Println()
	// }

	N := 5
	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
