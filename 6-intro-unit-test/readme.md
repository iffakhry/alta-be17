# Unit Testing Notes

## langkah-langkah melakukan proses unit test
* Jalankan `go mod init namaproject`

## menjalankan file test
* `go test ./salary/... -v`
* `go test ./salary/... -v -cover` (sekalligus menampilkan coverage nya)
* `go test ./salary/... -coverprofile=cover.out && go tool cover -html=cover.out`