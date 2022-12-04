package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes() {
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/user")
	v1.Get("/user")
	v1.Patch("/user")
	v1.Delete("/user/:id")

	app.Listen(":3000")
	//router := gin.Default()
	//
	//publicRoutes := router.Group("/auth")
	//publicRoutes.POST("/signup", controllers2.SignUp)
	//publicRoutes.POST("/signin", controllers2.SignIn)
	//
	//protectedRoutes := router.Group("/api/v1")
	//protectedRoutes.Use(jwt.AuthMiddleware())
	//protectedRoutes.POST("/entry", controllers2.AddEntry)
	//protectedRoutes.GET("/entry", controllers2.GetAllEntries)
	//
	//router.Run(":8000")
	//fmt.Println("Server running on port 8000")
}
