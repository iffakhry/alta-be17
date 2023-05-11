package main

import (
	"database/sql"
	"fmt"
	"log"

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
		fmt.Println("berhasil")
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	// query
	db.Close()
}
