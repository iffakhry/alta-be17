package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i == 1 {
			fmt.Println("continue")
			continue
		}
		if i > 3 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}
}
