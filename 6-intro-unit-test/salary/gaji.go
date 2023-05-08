package salary

func HitungGaji(posisi string) int {
	var gaji int
	if posisi == "c-level" {
		gaji = 50000000
	} else if posisi == "manager" {
		gaji = 30000000
	} else if posisi == "senior" {
		gaji = 30000000
	} else if posisi == "junior" {
		gaji = 10000000
	} else {
		gaji = 0
	}

	return gaji
}

func SalaryTotal(gaji ...int) int {
	var total int
	// for i := 0; i < len(gaji); i++ {
	// 	total = total + gaji[i]
	// }
	for _, value := range gaji {
		total = total + value
	}
	return total
}
