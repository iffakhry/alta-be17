package main

import "fmt"

func main() {
	var nameInput string
	fmt.Println("Masukkan nama anda!:")
	fmt.Scanln(&nameInput)
	fmt.Println("nama:", nameInput)

	var bilangan int
	fmt.Println("Masukkan angka:")
	fmt.Scanln(&bilangan)
	var hasil = bilangan + 15
	fmt.Printf("Bilangan inputan: %d , Bilangan hasil: %d %T \n", bilangan, hasil, bilangan)
}
