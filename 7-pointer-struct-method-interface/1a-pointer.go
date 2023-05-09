package main

import "fmt"

// https://blog.penjee.com/passing-by-value-vs-by-reference-java-graphical/

// pass by value
func fillCupValue(cup string) string {
	fmt.Println("data awal", cup) // kosong
	fmt.Println("address data awal", &cup)
	cup = "kopi"
	fmt.Println("data akhir", cup) // kopi
	return cup                     //kopi
}

// pass by reference
func fillCupPointer(cup *string) string {
	fmt.Println("data address awal", cup)
	fmt.Println("data awal", *cup) //kosong
	*cup = "kopi"
	fmt.Println("data akhir", *cup) // kopi
	return *cup                     //kopi
}

func main() {
	var cup string = "kosong"
	fmt.Println("data awal main", cup)          //kosong
	fmt.Println("address data awal main", &cup) // 0x0000
	// hasil1 := fillCupValue(cup)
	hasil1 := fillCupPointer(&cup)
	fmt.Println("hasil1", hasil1) //kopi
	fmt.Println("cup main", cup)  // kopi
}
