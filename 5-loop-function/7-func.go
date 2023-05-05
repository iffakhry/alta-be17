package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
	fmt.Println("semuanya")
}

func greeting(hour int) {
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	} else if hour < 18 {
		fmt.Println("Selamat Sore")
	} else {
		fmt.Println("Selamat Malam")
	}
}

func sapaOrang(nama string) {
	fmt.Println("halo", nama)
}

func main() {
	data := 10
	greeting(data)
	sayHello()
	var data1 = "adi"
	sapaOrang(data1)
}
