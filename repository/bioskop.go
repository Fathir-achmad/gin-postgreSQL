package repository

import (
	"database/sql"
	"gin-postgresql/structs"
)

func CreateBioskop(db *sql.DB, bioskop structs.Bioskop) (int, error) {
	var id int
	err := db.QueryRow(
		"INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id",
		bioskop.Nama, bioskop.Lokasi, bioskop.Rating,
	).Scan(&id)
	return id, err
}

func GetAllBioskop(db *sql.DB) ([]structs.Bioskop, error) {
	rows, err := db.Query("SELECT id, nama, lokasi, rating FROM bioskop")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bioskops []structs.Bioskop
	for rows.Next() {
		var b structs.Bioskop
		rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)
		bioskops = append(bioskops, b)
	}
	return bioskops, nil
}

func GetBioskopByID(db *sql.DB, id string) (structs.Bioskop, error) {
	var bioskop structs.Bioskop
	err := db.QueryRow("SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1", id).
		Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)
	return bioskop, err
}

func UpdateBioskop(db *sql.DB, id string, bioskop structs.Bioskop) (int64, error) {
	res, err := db.Exec(
		"UPDATE bioskop SET nama=$1, lokasi=$2, rating=$3 WHERE id=$4",
		bioskop.Nama, bioskop.Lokasi, bioskop.Rating, id,
	)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DeleteBioskop(db *sql.DB, id string) (int64, error) {
	res, err := db.Exec("DELETE FROM bioskop WHERE id = $1", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
