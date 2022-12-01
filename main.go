package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"rehoboam/controller"
	"rehoboam/database"
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
	publicRoutes.POST("/signUp", controller.SignUp)
	publicRoutes.POST("/signIn", controller.SignIn)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
