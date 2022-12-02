package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"rehoboam/controller"
	"rehoboam/database"
	"rehoboam/middleware"
	"rehoboam/model"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	db := database.Connect()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Entry{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/signup", controller.SignUp)
	publicRoutes.POST("/signin", controller.SignIn)

	protectedRoutes := router.Group("/api/v1")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
