package main

import "fmt"

func selectionSort(elements []int) []int {
	n := len(elements)
	for k := 0; k < n; k++ {
		minimal := k
		for j := k + 1; j < n; j++ {
			if elements[j] < elements[minimal] {
				minimal = j
			}
		}
		//proses pertukanran datananya
		elements[k], elements[minimal] = elements[minimal], elements[k]
	}
	return elements
}

func main() {
	data := []int{1, 4, 6, 8, 2, 3, 9, 10, 15, 13, 17, 12, 20, 21, 23, 24, 25, 26}
	result := selectionSort(data)
	fmt.Println("hasil", result)

	// var a = "teh"
	// var b = "kopi"

	// var c = b
	// b = a
	// a = c
	// fmt.Println("a", a, "b", b)
}
