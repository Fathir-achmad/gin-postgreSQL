package routes

import (
	"database/sql"
	"gin-postgresql/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.POST("/bioskop", func(ctx *gin.Context) {
		controllers.CreateBioskop(ctx, db)
	})
	router.GET("/bioskop", func(ctx *gin.Context) {
		controllers.AllBioskop(ctx, db)
	})

	return router
}
