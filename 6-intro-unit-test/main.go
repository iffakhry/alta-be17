package main

import (
	salaryBonus "alta/project/bonus/salary"
	"alta/project/bonus/thr"
	"alta/project/salary"
	"fmt"
)

// 1. aplikasi untuk menghitung penjumlahan 2 bilangan + 5
func add2Numbers(number1, number2 int) int {
	var hasil = number1 + number2 + 5
	return hasil
}

/*
2. aplikasi untuk menghitung gaji karyawan disebuah perusahaan.
klasifikas gaji:
a. c-level = 55 juta + 10 jt
b. manager = 35 juta + 5 jt
c. senior = 30 juta + 5jt
d. junior = 10 juta + 2 jt
*/

func main() {
	var nama1 = "Budi"
	var posisi1 = "c-level"
	var gaji1 int = salary.HitungGaji(posisi1)
	var gaji1Bonus int = salaryBonus.HitungGajiBonus(posisi1)
	gajithr1 := thr.HitungBonusThr(posisi1, 1)
	gajitotal1 := salary.SalaryTotal(gaji1, gaji1Bonus, gajithr1)
	fmt.Println("gaji total 1:", gajitotal1)
	// if posisi1 == "c-level" {
	// 	gaji1 = 50000000
	// } else if posisi1 == "manager" {
	// 	gaji1 = 35000000
	// } else if posisi1 == "senior" {
	// 	gaji1 = 30000000
	// } else if posisi1 == "junior" {
	// 	gaji1 = 10000000
	// } else {
	// 	gaji1 = 0
	// }

	fmt.Printf("nama: %s, posisi: %s, gaji: %d, gajithr: %d\n ", nama1, posisi1, gaji1, gajithr1)

	var nama2 = "Rudi"
	var posisi2 = "senior"
	var gaji2 int = salary.HitungGaji(posisi2)
	// if posisi2 == "c-level" {
	// 	gaji2 = 50000000
	// } else if posisi2 == "manager" {
	// 	gaji2 = 35000000
	// } else if posisi2 == "senior" {
	// 	gaji2 = 30000000
	// } else if posisi2 == "junior" {
	// 	gaji2 = 10000000
	// } else {
	// 	gaji2 = 0
	// }
	fmt.Printf("nama: %s, posisi: %s, gaji: %d\n ", nama2, posisi2, gaji2)

	// var bil1 = 10
	// var bil2 = 50

	// var output = Add2Numbers(bil1, bil2)
	// fmt.Println(output)

	// var angka1 = 20
	// var angka2 = 30
	// var output2 = Add2Numbers(angka1, angka2)
	// fmt.Println(output2)

	// var angka1a = 40
	// var angka1b = 50
	// var output3 = Add2Numbers(angka1a, angka1b)
	// fmt.Println(output3)
}
