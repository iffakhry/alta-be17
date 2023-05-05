package main

import "fmt"

func main() {
	var dataInput [3]int
	for i := 0; i < len(dataInput); i++ {
		fmt.Println("Masukkan data index ke", i)
		fmt.Scanln(&dataInput[i])
	}

	fmt.Println(dataInput)
}
