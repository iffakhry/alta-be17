package salary

func HitungGajiBonus(posisi string) int {
	var gaji int
	if posisi == "c-level" {
		gaji = 10000000
	} else if posisi == "manager" {
		gaji = 5000000
	} else if posisi == "senior" {
		gaji = 5000000
	} else if posisi == "junior" {
		gaji = 1000000
	} else {
		gaji = 0
	}

	return gaji
}
