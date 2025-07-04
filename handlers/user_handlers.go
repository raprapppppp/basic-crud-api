package handlers

import (
	//	"go_fiber/db"
	//	"go_fiber/models"
	"fmt"
	"go_fiber/models"
	"go_fiber/services"

	"github.com/gofiber/fiber/v2"
)

// inject Interface serviceDependeciy
type UserHandler struct {
	handler services.UserServiceDepend
}

func NewUserHandler(s services.UserServiceDepend) *UserHandler {
	return &UserHandler{s}
}

func (s *UserHandler) Getuser(h *fiber.Ctx) error {
	users, err := s.handler.FindAll()

	if err != nil {
		return h.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return h.Status(fiber.StatusOK).JSON(users)
}

func (s *UserHandler) CreateUser(h *fiber.Ctx) error {
	var user models.User

	if err := h.BodyParser(&user); err != nil {
		return h.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	createdUser, err := s.handler.CreateUser(user)

	if err != nil {
		return h.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return h.Status(fiber.StatusAccepted).JSON(createdUser)
}

func (s *UserHandler) UpdateUser(h *fiber.Ctx) error {
	id, err := h.ParamsInt("id")
	var user models.User

	if err != nil {
		h.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Invalid request data",
			"detail": err.Error(),
		})
	}
	errr := h.BodyParser(&user)
	if errr != nil {
		return h.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	updatedUser, _ := s.handler.UpdateUser(user, id)
	fmt.Print(updatedUser)
	return h.Status(fiber.StatusOK).JSON(updatedUser)
}

func (s *UserHandler) DeleteUser(h *fiber.Ctx) error {
	id, err := h.ParamsInt("id")
	var user models.User

	if err != nil {
		return h.Status(400).JSON(fiber.Map{"error": "Cannot Can't Delete 1"})

	}

	errr := s.handler.DeleteUser(user, id)
	if errr != nil {
		return h.Status(400).JSON(fiber.Map{"error": "Cannot Can't Delete 2"})
	}

	return h.SendStatus(200)
}

// func Getuser(h *fiber.Ctx) error {
// 	var users []models.User

// 	db.Database.Find(&users)
// 	return h.Status(fiber.StatusFound).JSON(users)
// }

// func GetUserById(h *fiber.Ctx) error {
// 	id, _ := h.ParamsInt("id")

// 	var user models.User

// 	db.Database.Find(&user, id)

// 	return h.Status(fiber.StatusOK).JSON(user)

// }

// func Adduser(h *fiber.Ctx) error {
// 	user := new(models.User)

// 	err := h.BodyParser(user)
// 	if err != nil {
// 		return h.Status(503).SendString(err.Error())
// 	}
// 	db.Database.Create(&user)

// 	return h.Status(fiber.StatusOK).JSON(user)

// }

// func UpdateUser(h *fiber.Ctx) error {
// 	user := new(models.User)
// 	id, _ := h.ParamsInt("id")

// 	err := h.BodyParser(user)
// 	if err != nil {
// 		return h.Status(504).SendString(err.Error())
// 	}

// 	db.Database.Where("id = ?", id).Updates(&user)
// 	return h.Status(fiber.StatusOK).JSON(user)
// }

// func DeleteUser(h *fiber.Ctx) error {
// 	id, _ := h.ParamsInt("id")
// 	var user models.User

// 	result := db.Database.Delete(&user, id)

// 	if result.RowsAffected == 0 {
// 		return h.SendStatus(404)
// 	}
// 	return h.SendStatus(200)

// }
