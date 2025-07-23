package handlers

import (
	//	"go_fiber/db"
	//	"go_fiber/models"
	"fmt"
	"go_fiber/config"
	"go_fiber/models"
	"go_fiber/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
	user := new(models.Users)

	if err := h.BodyParser(user); err != nil {
		return h.Status(500).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	createdUser, err := s.handler.CreateUser(user)

	if err != nil {
		return h.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return h.Status(fiber.StatusAccepted).JSON(createdUser)
}

func (s *UserHandler) UpdateUser(h *fiber.Ctx) error {
	var user models.Users

	errr := h.BodyParser(&user)
	if errr != nil {
		return h.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	updatedUser, _ := s.handler.UpdateUser(user)
	fmt.Print(updatedUser)
	return h.Status(fiber.StatusOK).JSON(updatedUser)
}

func (s *UserHandler) DeleteUser(h *fiber.Ctx) error {
	var user models.Users

	err := h.BodyParser(&user)

	if err != nil {
		return h.Status(400).JSON(fiber.Map{"error": "Cannot Can't Delete 1"})
	}
	errr := s.handler.DeleteUser(user)
	if errr != nil {
		return h.Status(400).JSON(fiber.Map{"error": "Cannot Can't Delete 2"})
	}
	return h.SendStatus(200)
}

// Creating accoutn for login
func (s *UserHandler) CreateUserAccount(h *fiber.Ctx) error {
	var account = new(models.Account)

	err := h.BodyParser(account)
	if err != nil {
		return h.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "404 Invalid request body",
		})
	}

	mess, err := s.handler.CreateAccountService(account)
	if err != nil {
		return err
	}
	if mess == "Exist" {
		return h.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": "Already Exist",
		})
	}
	return h.SendStatus(fiber.StatusAccepted)
}

// After login EU get token
func (s *UserHandler) LoginUserAccount(h *fiber.Ctx) error {
	var loginCredential models.Account

	err := h.BodyParser(&loginCredential)
	if err != nil {
		return err
	}

	resultMatching, mess  := s.handler.LoginUserAccountService(loginCredential)
	switch mess {
	case "User Does not exist exist":
		return h.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": 404,
		})

	case "Error in Database":
		return h.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": 500,
		})

	case "Password does not match":
		return h.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": 401,
		})

	case "Account match":
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id" : resultMatching.ID,
		"username":  resultMatching.Username,
		"role": resultMatching.Role,
		"exp":   time.Now().Add(time.Minute * 60).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Config("COOKIES_SECRET_KEY")))
	if err != nil {
		return h.SendStatus(fiber.StatusInternalServerError)
	}

	fmt.Print(t)
	//Creating Cookies struct may other way setcookie
	h.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		Expires:  time.Now().Add(60 * time.Minute),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
	})

	return h.Status(fiber.StatusOK).JSON(fiber.Map{
		"alert": "succesfull login",
	})

	/* return h.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": t,
	}) */
}

func LogoutUser(h *fiber.Ctx) error {
	h.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
	})

	return h.SendStatus(fiber.StatusOK)

}
