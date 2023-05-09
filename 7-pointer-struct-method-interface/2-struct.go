package main

import "fmt"

type karyawan struct {
	nama     string
	posisi   string
	domisili string
	umur     uint
	nilai    []NilaiKaryawan
}

type NilaiKaryawan struct {
	Bulan      string
	Nilai      int
	NilaiHuruf string
}

type Orang struct {
	Nama   string
	Umur   uint
	Weight float32
}

type Mobil struct {
	Type       string
	CC         int
	BahanBakar string
	Merk       string
}

func main() {
	var karyawan1 karyawan
	karyawan1.nama = "Budi"
	karyawan1.posisi = "c-level"
	karyawan1.domisili = "Jakarta"
	karyawan1.umur = 20

	fmt.Println("karyawan1:", karyawan1.nama)

	var karyawan2 karyawan
	karyawan2.nama = "rudi"
	karyawan2.posisi = "manager"
	karyawan2.domisili = "surabaya"

	var orang3 = karyawan{"andi", "manager", "surabaya", 30, []NilaiKaryawan{}}

	var pegawai []karyawan
	pegawai = append(pegawai, karyawan1)
	pegawai = append(pegawai, karyawan2)
	pegawai = append(pegawai, orang3)
	pegawai = append(pegawai, karyawan{
		umur:     20,
		nama:     "Adi",
		posisi:   "senior",
		domisili: "surabaya",
	})
	fmt.Println("pegawai:", pegawai)
	for _, v := range pegawai {
		fmt.Println(v.nama)
	}
}
