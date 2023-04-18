package main

import "fmt"

var data int

func main() {
	var bilangan1 int = 5
	var bilangan2 int = 3

	var hasil = bilangan1 + bilangan2
	fmt.Println("hasil tambah:", hasil)

	hasil1 := bilangan1 - bilangan2
	fmt.Println("hasil kurang:", hasil1)

	hasil2 := bilangan1 * bilangan2
	fmt.Println("hasil kali:", hasil2)

	hasil3 := bilangan1 / bilangan2
	fmt.Println("hasil bagi:", hasil3)

	hasil4 := bilangan1 % bilangan2
	fmt.Println("hasil modulo:", hasil4)

	// var hasil5 float64 = float64(bilangan1 / bilangan2)
	var hasil5 float64 = float64(bilangan1) / float64(bilangan2)
	fmt.Println("hasil bagi float:", hasil5)

	var nama1 string = "elon"
	var nama2 string = "musk"
	hasilnama := nama1 + "-" + nama2
	fmt.Println("hasilnama", hasilnama)

	data = 10
}
