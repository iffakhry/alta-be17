package main

import (
	"fmt"
	"math"
)

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Segitiga struct {
	alas   float64
	tinggi float64
}

func (r Rect) Luas() float64 {
	return r.width * r.height
}

func (c Circle) Luas() float64 {
	if c.radius == 0 {
		return 0.0
	}
	return math.Pi * c.radius * c.radius
}

func (s Segitiga) Luas() float64 {
	return (s.alas * s.tinggi) / 2
}

func Luas(tipe string) {
	if tipe == "circle" {
		//  3.14 * jari * jari
	} else if tipe == "rectangle" {
		// panjang * lebar

	} else if tipe == "segitiga" {
		// 0.5 * alas * tinggi
	}
}

func LuasLingkaran() {
	// if jari == 0
}

func LuasPersegiPanjang() {

}

func LuasSegitiga() {

}

// luas lingkaran = phi * r^2

func main() {
	// rect := Rect{5.0, 4.0}
	// var rect Rect
	// rect.width = 5.0
	// rect.height = 4.0
	rect := Rect{
		width:  5.0,
		height: 4.0,
	}

	cir := Circle{5.0}
	segitiga := Segitiga{
		alas:   10,
		tinggi: 15,
	}
	fmt.Printf("Area of rectangle rect = %0.2f\n", rect.Luas())
	fmt.Printf("Area of circle cir = %0.2f\n", cir.Luas())
	fmt.Printf("Area of segitiga = %0.2f\n", segitiga.Luas())
}
