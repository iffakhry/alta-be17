package main

import "fmt"

func main() {
	type Siswa struct {
		Nama  string
		Kelas string
		Nilai uint
	}

	var siswaData Siswa
	fmt.Println("Masukkan nama siswa:")
	fmt.Scanln(&siswaData.Nama)
	fmt.Println("Masukkan kelas siswa:")
	fmt.Scanln(&siswaData.Kelas)
	fmt.Println("Masukkan Nilai siswa:")
	fmt.Scanln(&siswaData.Nilai)

	fmt.Println(siswaData)
	fmt.Println(siswaData.Nama)
	fmt.Println(&siswaData.Nama)
	fmt.Println(&siswaData.Kelas)
}
