package main

import (
	"database/sql"
	"gin-postgresql/routes"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=fathiras1905 dbname=bioskopdb sslmode=disable"
	var PORT = ":8080"

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

	routes.StartServer(db).Run(PORT)
}

// Get All 
// curl --location --request GET 'http://localhost:8080/bioskop/' \
// --header 'Content-Type: application/json' \

// Create Bioskop 
// curl --location 'http://localhost:8080/bioskop/' \
// --header 'Content-Type: application/json' \
// --data '{"nama":"Netflix","lokasi":"Mobile","rating":4.8}'