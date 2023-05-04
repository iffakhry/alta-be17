package main

import "fmt"

func main() {
	var myAge = 20
	if myAge == 17 {
		fmt.Println("Umurnya 17 tahun")
	}

	if myAge == 20 {
		fmt.Println("umurnya 20 tahun")
	}

	if myAge > 10 {
		fmt.Println("umunya diatas 10 tahun")
		myAge = myAge + 10
	}

	if myAge > 10 && myAge == 17 {
		fmt.Println("sweet 17")
	}

	fmt.Println("selesai")
}
