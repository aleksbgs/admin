package controllers

import (
	"github.com/gofiber/fiber"
)
import "../models"
import "../database"

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.SetPassword("1234")

	database.DB.Create(&user)
	return c.JSON(user)
}
