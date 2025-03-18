package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB //переменная db, которая будет хранить соединение с базой данных.

// соединяет с бд
func initDB() {
	dsn := "host=localhost user=postgres password=12345678 dbname=cafee port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Миграция схемы
	db.AutoMigrate(&Menu{})
}

func main() {
	initDB()
	router := gin.Default()

	router.POST("/login", login)
	router.POST("/register", register)

	protected := router.Group("/")
	protected.Use(authMiddleware())
	{
		protected.GET("/dishes", getDishes)
		protected.GET("/dishes/:id", getDishByID)
	}

	protected.Use(adminOnlyMiddleware())
	{

		protected.PUT("/dishes/:id", updateDish)
		protected.DELETE("/dishes/:id", deleteDish)
		protected.POST("/dishes", createDish)
	}

	router.Run(":8080")
}
