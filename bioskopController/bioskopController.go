package bioskopcontroller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// inisialisasi struct bioskop
type Bioskop struct {
	BioskopID int64   `json:"bioskop_id"`
	Nama      string  `json:"nama"`
	Lokasi    string  `json:"lokasi"`
	Rating    float64 `json:"rating"`
}

var BioskopData = []Bioskop{}

// CreateBioskop menambahkan data bioskop baru
func CreateBioskop(ctx *gin.Context) {
	var newBioskop Bioskop
	if err := ctx.ShouldBindJSON(&newBioskop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBioskop.BioskopID = int64(len(BioskopData) + 1)
	BioskopData = append(BioskopData, newBioskop)
	ctx.JSON(http.StatusCreated, gin.H{"message": "data bioskop berhasil ditambahkan", "data": newBioskop})
}

// GetAllBioskop mengambil seluruh data bioskop
func GetAllBioskop(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": BioskopData,
	})
}

// GetBioskop mengambil seluruh data atau berdasarkan ID jika id diberikan
func GetBioskop(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"data": BioskopData,
		})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":         "invalid id",
			"error_message": "ID bioskop harus berupa angka",
		})
		return
	}

	for _, bioskop := range BioskopData {
		if bioskop.BioskopID == id {
			ctx.JSON(http.StatusOK, gin.H{
				"data": bioskop,
			})
			return
		}
	}

	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"error":         "data not found",
		"error_message": fmt.Sprintf("data dengan id: %d tidak ditemukan", id),
	})
}

// UpdateBioskop memperbarui data bioskop berdasarkan ID
func UpdateBioskop(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "error_message": "ID bioskop harus berupa angka"})
		return
	}

	var updatedBioskop Bioskop
	if err := ctx.ShouldBindJSON(&updatedBioskop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	condition := false
	for i, v := range BioskopData {
		if v.BioskopID == id {
			BioskopData[i].Nama = updatedBioskop.Nama
			BioskopData[i].Lokasi = updatedBioskop.Lokasi
			BioskopData[i].Rating = updatedBioskop.Rating
			updatedBioskop.BioskopID = id
			condition = true
			break
		}
	}

	if !condition {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "data bioskop berhasil di update", "data": updatedBioskop})
}

// DeleteBioskop menghapus data bioskop berdasarkan ID
func DeleteBioskop(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":         "invalid id",
			"error_message": "ID bioskop harus berupa angka",
		})
		return
	}

	condition := false
	var bioskopIndex int

	for i, v := range BioskopData {
		if v.BioskopID == id {
			condition = true
			bioskopIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":         "data not found",
			"error_message": fmt.Sprintf("data dengan id: %d tidak ditemukan", id),
		})
		return
	}

	copy(BioskopData[bioskopIndex:], BioskopData[bioskopIndex+1:])
	BioskopData[len(BioskopData)-1] = Bioskop{}
	BioskopData = BioskopData[:len(BioskopData)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("bioskop dengan id %v berhasil dihapus", id),
	})
}
