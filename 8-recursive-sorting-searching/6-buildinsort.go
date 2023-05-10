package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	/*
		ASC : kecil ke besar <
		DESC: besar ke kecil >
	*/
	// data := []int{1, 4, 6, 8, 2, 3, 9, 10, 15, 13, 17, 12, 20, 21, 23, 24, 25, 26}
	// data := []string{"c", "az", "abc", "abb", "ac", "b"}
	data := []string{"3", "9", "7", "8"}
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	fmt.Println("data:   ", data)

}
