package controllers

import (
	"database/sql"
	"fakhry/rawsql/entities"
	"log"
	"os"
)

func GetAllProducts(db *sql.DB) []entities.Product {
	db, _ = sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	rows, errSelect := db.Query("SELECT id, nama_produk, harga, stock, keterangan FROM products")
	if errSelect != nil {
		log.Fatal("error query select", errSelect.Error())
	}

	var allProduct []entities.Product // menampung semua data
	for rows.Next() {                 // proses pembacaan per baris
		var product entities.Product                                                                                // menampung data per baris nya
		errScan := rows.Scan(&product.ID, &product.NamaProduk, &product.Harga, &product.Stock, &product.Keterangan) // mapping
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())
		}

		allProduct = append(allProduct, product)
	}

	// fmt.Println("all:", allProduct)
	return allProduct
}
