package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed sql_migrations/*.sql
var dbMigrations embed.FS

var DbConnection *sql.DB

// ConnectDB bikin koneksi ke PostgreSQL
func ConnectDB() *sql.DB {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		// fallback kalau jalan lokal
		connStr = "host=127.0.0.1 port=5432 user=postgres password=fathiras1905 dbname=go_bioskop sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	// test ping
	if err := db.Ping(); err != nil {
		log.Fatal("Database tidak bisa di-ping:", err)
	}

	fmt.Println("Database terkoneksi 🚀")
	return db
}

// DBMigrate untuk jalankan migration
func DBMigrate(dbParam *sql.DB) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migrations",
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	DbConnection = dbParam
	fmt.Println("Migration success, applied", n, "migrations!")
}
