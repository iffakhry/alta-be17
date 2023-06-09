package controllers

import (
	"database/sql"
	"fakhry/rawsql/entities"
	"fmt"
	"log"
)

func GetAllProducts(db *sql.DB) []entities.Product {
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

func AddProduct(db *sql.DB, newProduct entities.Product) error {
	result, err := db.Exec("INSERT INTO products (nama_produk, harga, stock, keterangan) VALUES (?, ?,?,?)", newProduct.NamaProduk, newProduct.Harga, newProduct.Stock, newProduct.Keterangan)
	if err != nil {
		return fmt.Errorf("AddAlbum: %v", err)
	}

	// Get the new album's generated ID for the client.
	_, errId := result.LastInsertId()
	if errId != nil {
		return fmt.Errorf("AddAlbum: %v", errId)
	}
	// Return the new album's ID.
	return nil
}
