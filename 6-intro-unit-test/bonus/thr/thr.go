package thr

func HitungBonusThr(posisi string, lama int) int {
	var bonus int
	if posisi == "c-level" {
		bonus = 1000000 * lama
	} else if posisi == "manager" {
		bonus = 500000 * lama
	} else if posisi == "senior" {
		bonus = 500000 * lama
	} else if posisi == "junior" {
		bonus = 200000 * lama
	} else {
		bonus = 0
	}

	return bonus
}
