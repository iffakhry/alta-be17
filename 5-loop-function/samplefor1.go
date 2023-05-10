package main

import "fmt"

/*
tampilkan bilangan genap dari 1 sampai n
n adalah inputan
contoh:
n = 10
output: 2,4,6,8,10

1,2,3,4,5,6,7,8,9,10

10,9,8,7 ... 1

n=20
output: 2,4,6,8,10,12,14,16,18,20
*/

func TampilGenap(n int) []int {
	// variabel penampung hasil, untuk return value
	var hasil []int
	var temp float32 = 0.0
	// untuk memunculkan angka" berurutan mulai dari 1 sampai n
	// for i := 1; i <= n; i++ {
	for i := n; i > 0; i = i - 1 {

		// fmt.Println(i)
		// mengecek apakah value i habis dibagi 2 (genap)
		if i%2 == int(temp) {
			// kalau terpenuhi, maka masukkan value i ke variabel penampung
			hasil = append(hasil, i)
		}
	}
	//return
	return hasil
}

func main() {
	result := TampilGenap(10)
	fmt.Println("result:", result)
	if len(result) == 0 {
		fmt.Println("gak ada isinya")
	}
}
