package main

import "fmt"

type calculateInterface interface {
	luas() int
	keliling() int
	// volume(angka int) string
}

type persegi struct {
	sisi int
}

type segitiga struct {
	alas   int
	tinggi int
}

func (p persegi) luas() int {
	return p.sisi * p.sisi
}

func (p persegi) keliling() int {
	return 4 * p.sisi
}

func (s segitiga) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitiga) keliling() int {
	return (s.alas * s.tinggi) / 2
}

func main() {
	var p1 = persegi{
		sisi: 10,
	}
	fmt.Println("luas persegi", p1.luas())

	var Ipersegi calculateInterface
	Ipersegi = persegi{
		sisi: 10,
	}
	fmt.Println("luas persegi", Ipersegi.luas())

	var Isegitiga calculateInterface
	Isegitiga = segitiga{
		alas:   10,
		tinggi: 15,
	}
	fmt.Println("luas segitiga", Isegitiga.luas())
}
