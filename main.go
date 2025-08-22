package main

import (
    "database/sql"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "gin-postgresql/controllers"
    "gin-postgresql/database"
    "os"

    _ "github.com/lib/pq"
)

var (
    DB  *sql.DB
    err error
)

func main() {
    // Load .env
   err = godotenv.Load("config/.env")
    if err != nil {
       panic("Error loading .env file")
    }

    // Build connection string
    psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
       os.Getenv("PGHOST"),
       os.Getenv("PGPORT"),
       os.Getenv("PGUSER"),
       os.Getenv("PGPASSWORD"),
       os.Getenv("PGDATABASE"),
    )

    // Open DB
      DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
       panic(err)
    }

    // Test connection
    err = DB.Ping()
    if err != nil {
       panic(err)
    }

    // Run migration
    database.DBMigrate(DB)

   router := gin.Default()
	router.GET("/bioskop", func(c *gin.Context) { controllers.AllBioskop(c, DB) })
	router.POST("/bioskop", func(c *gin.Context) { controllers.CreateBioskop(c, DB) })
	router.GET("/bioskop/:id", func(c *gin.Context) { controllers.GetBioskopByID(c, DB) })
	router.PUT("/bioskop/:id", func(c *gin.Context) { controllers.UpdateBioskop(c, DB) })
	router.DELETE("/bioskop/:id", func(c *gin.Context) { controllers.DeleteBioskop(c, DB) })

	router.Run(":" + os.Getenv("PORT"))

}
