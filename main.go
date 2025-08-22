package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gin-postgresql/routes"

	_ "github.com/lib/pq"
)

func main() {
	// Railway kasih DATABASE_URL, kalau lokal fallback manual
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "host=127.0.0.1 port=5432 user=postgres password=fathiras1905 dbname=bioskopdb sslmode=disable"
	}

	// Buka koneksi database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}
	defer db.Close()

	// Tes koneksi database
	if err := db.Ping(); err != nil {
		log.Fatalf("Database tidak bisa di-ping: %v", err)
	}

	fmt.Println("✅ Berhasil konek ke database!")

	// Buat tabel kalau belum ada
	createTable := `
	CREATE TABLE IF NOT EXISTS bioskop (
		id SERIAL PRIMARY KEY,
		nama VARCHAR(100) NOT NULL,
		lokasi VARCHAR(100) NOT NULL,
		rating FLOAT
	);`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatalf("Gagal membuat tabel: %v", err)
	}

	fmt.Println("✅ Tabel bioskop siap digunakan")

	// Railway juga kasih PORT, default ke 8080 kalau lokal
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 Server jalan di port %s...\n", port)
	if err := routes.StartServer(db).Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
