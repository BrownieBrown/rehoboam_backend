package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"rehoboam/pkg/models"
)

func AddUser(context *fiber.Ctx) error {
	user := models.User{}

	err := context.BodyParser(&user)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}
}
