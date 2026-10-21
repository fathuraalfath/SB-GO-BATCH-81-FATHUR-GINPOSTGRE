package routers

import (
	"net/http"

	bioskopcontroller "SB-GO-BATCH-81-FATHUR-GINPOSTGRE/bioskopController"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// Route root
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Server API Bioskop Berjalan!",
			"info":    "Silakan akses endpoint /bioskop",
		})
	})

	router.POST("/bioskop", bioskopcontroller.CreateBioskop)
	router.GET("/bioskop", bioskopcontroller.GetBioskop)
	router.GET("/bioskop/:id", bioskopcontroller.GetBioskop)
	router.PUT("/bioskop/:id", bioskopcontroller.UpdateBioskop)
	router.DELETE("/bioskop/:id", bioskopcontroller.DeleteBioskop)
	return router
}
