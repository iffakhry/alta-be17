package main

import "fmt"

func main() {
	// long declaration
	var salary = map[string]int{}
	fmt.Println(salary)
	salary["adi"] = 10000
	salary["budi"] = 20000
	salary["rudi"] = 20000
	fmt.Println(salary)

	var data1 = map[int]int{}
	data1[9] = 90
	data1[12] = 120
	data1[9] = 129
	fmt.Println(data1)

	var bulanarr = [12]int{30, 28, 31, 30, 31, 30}
	fmt.Println(bulanarr[0]) // membaca jumlah hari di bulan januari
	fmt.Println(bulanarr[1]) // membaca jumlah hari di bulan februari
	var datemonth = map[string]int{}
	datemonth["januari"] = 30
	datemonth["februari"] = 28
	datemonth["maret"] = 31
	datemonth["april"] = 30
	fmt.Println(datemonth["januari"])
	fmt.Println(datemonth["februari"])

	var data2 = map[string]any{}
	data2["nama"] = "Budi"
	data2["umur"] = 20
	fmt.Println(data2)

	// long declaration with value
	var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
	fmt.Println(salary_a)

	// short declaration
	salary_b := map[string]int{}
	fmt.Println(salary_b)

	// using make
	var salary_c = make(map[string]int)
	salary_c["doe"] = 7000 // assign value
	fmt.Println(salary_c)
}
