package main

import (
	"database/sql"
	"fmt"

	"SB-GO-BATCH-81-FATHUR-GINPOSTGRE/routers"

	_ "github.com/lib/pq"
)

// koneksi database
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "bioskop"
)

var (
	db  *sql.DB
	err error
)

func main() {
	// Format string koneksi PostgreSQL
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Buka koneksi ke database
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Verifikasi koneksi ke database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Pesan sukses koneksi
	fmt.Println("Database berhasil terkoneksi")

	// Menjalankan server Gin
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
