package main

import (
	"fmt"
	"sort"
)

func main() {
	elements := []int{12, 15, 15, 19, 24, 31, 53, 59, 60}
	value := 55
	findIndex := sort.SearchInts(elements, value)

	if value == elements[findIndex] {
		fmt.Println("value found in elemenets. index:", findIndex)
	} else {
		fmt.Println("value not found in elemenets. index:", findIndex)
	}
}
