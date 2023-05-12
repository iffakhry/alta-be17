package main

import (
	"database/sql"
	"fakhry/rawsql/controllers"
	"fakhry/rawsql/entities"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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

	fmt.Println("Menu: \n1. Read \n2. Insert")
	fmt.Println("masukkan menu")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		dataProducts := controllers.GetAllProducts(db)
		// fmt.Println(dataProducts)
		for _, value := range dataProducts {
			fmt.Printf("id: %d, nama: %s, harga: %d \n", value.ID, value.NamaProduk, value.Harga)
		}

	case 2:
		var newproduct = entities.Product{
			NamaProduk: "buku",
			Harga:      10000,
			Stock:      10,
			Keterangan: "buku tulis",
		}
		idNew, err := AddProduct(db, newproduct)
		if err != nil {
			fmt.Println("error", err.Error())
		} else {
			fmt.Println("id porduct", idNew)
		}
	}

	// query select
	// dataProducts := GetAllProducts(db)
	// // fmt.Println(dataProducts)
	// for _, value := range dataProducts {
	// 	fmt.Printf("id: %d, nama: %s, harga: %d \n", value.ID, value.NamaProduk, value.Harga)
	// }
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

func AddProduct(db *sql.DB, newProduct entities.Product) (int64, error) {
	result, err := db.Exec("INSERT INTO products (nama_produk, harga, stock, keterangan) VALUES (?, ?,?,?)", newProduct.NamaProduk, newProduct.Harga, newProduct.Stock, newProduct.Keterangan)
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}

	// Get the new album's generated ID for the client.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}
	// Return the new album's ID.
	return id, nil
}
