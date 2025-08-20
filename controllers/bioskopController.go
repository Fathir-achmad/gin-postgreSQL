package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Bioskop struct {
	ID     int     `json:"id"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float64 `json:"rating"`
}

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

func AllBioskop(ctx *gin.Context, db *sql.DB) {
	rows, _ := db.Query("SELECT id, nama, lokasi, rating FROM bioskop")
	defer rows.Close()

	var bioskops []Bioskop
	for rows.Next() {
		var b Bioskop
		rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)
		bioskops = append(bioskops, b)
	}

	ctx.JSON(http.StatusOK, bioskops)
}

