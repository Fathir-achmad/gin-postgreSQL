package main

import (
	"database/sql"
	"gin-postgresql/routes"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Railway kasih DATABASE_URL
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		// fallback buat lokal
		connStr = "host=localhost port=5432 user=postgres password=fathiras1905 dbname=bioskopdb sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}
	defer db.Close()

	createTable := `
	CREATE TABLE IF NOT EXISTS bioskop (
		id SERIAL PRIMARY KEY,
		nama VARCHAR(100) NOT NULL,
		lokasi VARCHAR(100) NOT NULL,
		rating FLOAT
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Gagal membuat tabel:", err)
	}

	// Railway juga kasih PORT (default 8080 kalau lokal)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routes.StartServer(db).Run(":" + port)
}
