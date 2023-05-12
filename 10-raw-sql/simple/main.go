package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID         uint
	NamaProduk string
	Harga      int
	Stock      int
	Keterangan string
}

func main() {
	// <username>:<password>@tcp(<hostname>:<portdb>)/<db_name>
	var connectionString = "root:qwerty123@tcp(127.0.0.1:3306)/db_raw"
	db, err := sql.Open("mysql", connectionString)
	// ketika errornya ada isinya --> berarti terjadi error
	if err != nil {
		log.Fatal("error open connection", err.Error())
	} else {
		fmt.Println("berhasil open")
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// verif apakah sudah berhasil connect ke mysql
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("ping:", pingErr)
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()

	// query select
	dataProducts := GetAllProducts(db)
	// fmt.Println(dataProducts)
	for _, value := range dataProducts {
		fmt.Printf("id: %d, nama: %s, harga: %d \n", value.ID, value.NamaProduk, value.Harga)
	}
	// rows, errSelect := db.Query("SELECT id, nama_produk, harga, stock, keterangan FROM products")
	// if errSelect != nil {
	// 	log.Fatal("error query select", errSelect.Error())
	// }

	// var allProduct []Product // menampung semua data
	// for rows.Next() {        // proses pembacaan per baris
	// 	var product Product                                                                                         // menampung data per baris nya
	// 	errScan := rows.Scan(&product.ID, &product.NamaProduk, &product.Harga, &product.Stock, &product.Keterangan) // mapping
	// 	if errScan != nil {
	// 		log.Fatal("error scan", errScan.Error())
	// 	}

	// 	allProduct = append(allProduct, product)
	// }

	// fmt.Println("all:", allProduct)

	// query insert

}

func GetAllProducts(db *sql.DB) []Product {
	rows, errSelect := db.Query("SELECT id, nama_produk, harga, stock, keterangan FROM products")
	if errSelect != nil {
		log.Fatal("error query select", errSelect.Error())
	}

	var allProduct []Product // menampung semua data
	for rows.Next() {        // proses pembacaan per baris
		var product Product                                                                                         // menampung data per baris nya
		errScan := rows.Scan(&product.ID, &product.NamaProduk, &product.Harga, &product.Stock, &product.Keterangan) // mapping
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())
		}

		allProduct = append(allProduct, product)
	}

	// fmt.Println("all:", allProduct)
	return allProduct
}
