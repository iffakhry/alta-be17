package main

import (
	"fmt"
	"math"
)

// singe return value
func calculateSquare(side int) int {
	hasil := side * side
	return hasil
}

// multiple return value
func calculateCircle(diameter float64) (float64, float64) {
	var keliling = math.Pi * math.Pow(diameter/2, 2)
	var luas = math.Pi * diameter
	// return 2 value
	return keliling, luas
}

func main() {
	var side = 5
	wide := calculateSquare(side)
	fmt.Printf("Luas persegi empat: %d \n\n", wide)

	var sisi = 10
	data1 := calculateSquare(sisi)
	fmt.Println("luas persegi 1:", data1)

	var diameter float64 = 15
	keliling, luas := calculateCircle(diameter)
	fmt.Printf("Luas lingkaran: %.2f \n", keliling)
	fmt.Printf("Keliling lingkaran: %.2f \n", luas)
}
