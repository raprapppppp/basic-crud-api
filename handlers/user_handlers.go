package handlers

import (
	"go_fiber/db"
	"go_fiber/models"

	"github.com/gofiber/fiber/v2"
)

func Getuser(h *fiber.Ctx) error {
	var users []models.User

	db.Database.Find(&users)
	return h.Status(fiber.StatusFound).JSON(users)
}

func GetUserById(h *fiber.Ctx) error {
	id, _ := h.ParamsInt("id")

	var user models.User

	db.Database.Find(&user, id)

	return h.Status(fiber.StatusOK).JSON(user)

}

func Adduser(h *fiber.Ctx) error {
	user := new(models.User)

	err := h.BodyParser(user)
	if err != nil {
		return h.Status(503).SendString(err.Error())
	}
	db.Database.Create(&user)

	return h.Status(fiber.StatusOK).JSON(user)

}

func UpdateUser(h *fiber.Ctx) error {
	user := new(models.User)
	id, _ := h.ParamsInt("id")

	err := h.BodyParser(user)
	if err != nil {
		return h.Status(504).SendString(err.Error())
	}

	db.Database.Where("id = ?", id).Updates(&user)
	return h.Status(fiber.StatusOK).JSON(user)
}

func DeleteUser(h *fiber.Ctx) error {
	id, _ := h.ParamsInt("id")
	var user models.User

	result := db.Database.Delete(&user, id)

	if result.RowsAffected == 0 {
		return h.SendStatus(404)
	}
	return h.SendStatus(200)

}

func AddLoginUser(h *fiber.Ctx) error {
	user := new(models.Login)

	err := h.BodyParser(user)
	if err != nil {
		return h.Status(503).SendString(err.Error())
	}
	db.Database.Create(&user)

	return h.Status(fiber.StatusOK).JSON(user)

}

func GetLoginUser(h *fiber.Ctx) error {
	return nil

}
