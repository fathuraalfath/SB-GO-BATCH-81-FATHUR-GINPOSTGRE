package main

import (
	"database/sql"
	"fmt"
	"os"

	"SB-GO-BATCH-81-FATHUR-GINPOSTGRE/routers"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

// Helper untuk membaca environment variable dengan nilai default (fallback)
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func main() {
	// Muat file .env jika ada (untuk development lokal)
	_ = godotenv.Load()

	// Membaca variabel database dari environment
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "bioskop")
	sslmode := getEnv("DB_SSLMODE", "disable")

	// Format string koneksi PostgreSQL
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

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

	// Menjalankan server Gin dengan port dinamis dari environment
	portServer := getEnv("PORT", "8080")
	routers.StartServer().Run(":" + portServer)
}
