## Raw SQL

## Setup
* Jalankan `go mod init namaproject`
* `go get -u github.com/go-sql-driver/mysql`
* buat file main.go
* mulai coding

## DB
* buat dulu databasenya
* buat tablenya

## Create ENV
* buat file dengan nama `.env`
* tulis configurasi connection string kalian disana
```bash
#contoh: 
export DB_CONNECTION='root:password@tcp(127.0.0.1:3306)/db_name'
```
* buat file dengan nama `.gitignore`
* tambahkan `.env` di gitignore tsb

## Menjalankan app golang dengan env
Untuk menjalankan app golang yang memanfaatkan env, maka kita harus load file env terlebih dahulu.
<br>
note: jalankan di terminal bash atau linux.

```bash
source .env

go run main.go
```