package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Bioskop struct {
	ID     int     `json:"id"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float64 `json:"rating"`
}

// ---------------- CREATE ----------------
func CreateBioskop(ctx *gin.Context, db *sql.DB) {
	var bioskop Bioskop

	if err := ctx.ShouldBindJSON(&bioskop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	err := db.QueryRow(
		"INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id",
		bioskop.Nama, bioskop.Lokasi, bioskop.Rating,
	).Scan(&bioskop.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bioskop berhasil ditambahkan",
		"data":    bioskop,
	})
}

// ---------------- READ ALL ----------------
func AllBioskop(ctx *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, nama, lokasi, rating FROM bioskop")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer rows.Close()

	var bioskops []Bioskop
	for rows.Next() {
		var b Bioskop
		rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)
		bioskops = append(bioskops, b)
	}

	ctx.JSON(http.StatusOK, bioskops)
}

// ---------------- READ ONE ----------------
func GetBioskopByID(ctx *gin.Context, db *sql.DB) {
	id := ctx.Param("id")
	var bioskop Bioskop

	err := db.QueryRow("SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1", id).
		Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)

	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	ctx.JSON(http.StatusOK, bioskop)
}

// ---------------- UPDATE ----------------
func UpdateBioskop(ctx *gin.Context, db *sql.DB) {
	id := ctx.Param("id")
	var bioskop Bioskop

	if err := ctx.ShouldBindJSON(&bioskop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	res, err := db.Exec(
		"UPDATE bioskop SET nama=$1, lokasi=$2, rating=$3 WHERE id=$4",
		bioskop.Nama, bioskop.Lokasi, bioskop.Rating, id,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	// kembalikan id sebagai response
	idInt, _ := strconv.Atoi(id)
	bioskop.ID = idInt

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bioskop berhasil diperbarui",
		"data":    bioskop,
	})
}

// ---------------- DELETE ----------------
func DeleteBioskop(ctx *gin.Context, db *sql.DB) {
	id := ctx.Param("id")

	res, err := db.Exec("DELETE FROM bioskop WHERE id = $1", id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bioskop berhasil dihapus",
	})
}
