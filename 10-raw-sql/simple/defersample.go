package main

import "fmt"

func main() {
	namaFunc()
	defer fmt.Println("satu")
	defer fmt.Println("dua")
	defer fmt.Println("tiga")
	panic("error")
	fmt.Println("empat")
}

func namaFunc() {
	fmt.Println("func 1")
	defer fmt.Println("func 2")
	fmt.Println("func 3")
}
