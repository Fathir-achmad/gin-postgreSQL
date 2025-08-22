package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gin-postgresql/repository"
	"gin-postgresql/structs"
	"net/http"
	"strconv"
)

// ---------------- CREATE ----------------
func CreateBioskop(ctx *gin.Context, db *sql.DB) {
	var bioskop structs.Bioskop

	if err := ctx.ShouldBindJSON(&bioskop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	id, err := repository.CreateBioskop(db, bioskop)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}
	bioskop.ID = id

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bioskop berhasil ditambahkan",
		"data":    bioskop,
	})
}

// ---------------- READ ALL ----------------
func AllBioskop(ctx *gin.Context, db *sql.DB) {
	bioskops, err := repository.GetAllBioskop(db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	ctx.JSON(http.StatusOK, bioskops)
}

// ---------------- READ ONE ----------------
func GetBioskopByID(ctx *gin.Context, db *sql.DB) {
	id := ctx.Param("id")

	bioskop, err := repository.GetBioskopByID(db, id)
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
	var bioskop structs.Bioskop

	if err := ctx.ShouldBindJSON(&bioskop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	rowsAffected, err := repository.UpdateBioskop(db, id, bioskop)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data"})
		return
	}
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

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

	rowsAffected, err := repository.DeleteBioskop(db, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bioskop berhasil dihapus",
	})
}
