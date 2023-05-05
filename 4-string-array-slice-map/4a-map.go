package main

import "fmt"

func main() {
	// long declaration with value
	var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
	fmt.Println(salary_a, len(salary_a))

	salary_a["nabilah"] = 7000 // assign value
	fmt.Println(salary_a)

	delete(salary_a, "iswanul") // remove value by key
	fmt.Println(salary_a)

	var datemonth = map[string]int{}
	datemonth["januari"] = 30
	datemonth["februari"] = 28
	datemonth["maret"] = 31
	datemonth["april"] = 30
	datemonth["mei"] = 0
	value, exist := datemonth["mei"] // check key is exist
	valueJan, isExistJan := datemonth["januari"]
	fmt.Println("datanya:", valueJan, "isExist:", isExistJan)
	fmt.Println("datanya:", value, "isExist:", exist)

	if exist == true {
		fmt.Println("key nya sudah ada cuy")
	} else {
		fmt.Println("key nya belum ada cuy")

	}

	for key, value := range salary_a { // loop over maps
		fmt.Println("->", key, value)
	}
}
