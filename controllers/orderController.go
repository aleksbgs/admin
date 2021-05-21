package controllers

import (
	"../database"
	"../models"
	"github.com/gofiber/fiber"
	"strconv"
)

func AllOrders(c *fiber.Ctx) error {

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}
