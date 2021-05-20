package controllers

import (
	"../database"
	"../models"
	"../util"
	"github.com/gofiber/fiber"
	"strconv"
)

func AllUsers(c *fiber.Ctx) error {

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.User{}, page))
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
func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Preload("Role").Find(&user)
	return c.JSON(user)
}
func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}
	database.DB.Model(&user).Updates(user)
	return c.JSON(user)

}
func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		Id: uint(id),
	}
	database.DB.Delete(&user)
	return nil
}

func UpdateInfo(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id:        uint(userId),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	database.DB.Model(&user).Updates(user)
	return c.JSON(user)

}
func UpdatePassword(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password is not match",
		})

	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}
	user.SetPassword(data["password"])

	database.DB.Model(&user).Updates(user)
	return c.JSON(user)
}
