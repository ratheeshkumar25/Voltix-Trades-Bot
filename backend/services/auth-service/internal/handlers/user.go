package handlers

import (
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UseAccount struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//@Id Create User
//@Summary Create User
//@Tags User
//@Accept json
//@Produce json
//@Param user body UseAccount true "User Object"
//@Success 200 {object} HttpResponse
//@Failure 400 {object} HttpResponse
//@Failure 500 {object} HttpResponse
//@Router api/v1/user [post]

func (h *AuthServiceHandler) CreateUser(c *fiber.Ctx) error {
	var user UseAccount
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	resp, err := h.SVC.CreateUser(user.Email, user.Password, nil)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(resp)
}

// @Id Get User By Email
// @Summary Get User By Email
// @Tags User
// @Accept json
// @Produce json
// @Param email path string true "User Email"
// @Success 200 {object} HttpResponse
// @Failure 400 {object} HttpResponse
// @Failure 500 {object} HttpResponse
// @Router api/v1/user/{email} [get]
func (h *AuthServiceHandler) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	if email == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Email is required"})
	}
	resp, err := h.SVC.GetUserByEmail(email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

//@Id Get User By ID
//@Summary Get User By ID
//@Tags User
//@Accept json
//@Produce json
//@Param id path string true "User ID"
//@Success 200 {object} HttpResponse
//@Failure 400 {object} HttpResponse
//@Failure 500 {object} HttpResponse
//@Router api/v1/user/{id} [get]

func (h *AuthServiceHandler) GetUserByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	resp, err := h.SVC.GetUserByID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

//@Id Get User By Google ID
//@Summary Get User By Google ID
//@Tags User
//@Accept json
//@Produce json
//@Param google_id path string true "User Google ID"
//@Success 200 {object} HttpResponse
//@Failure 400 {object} HttpResponse
//@Failure 500 {object} HttpResponse
//@Router api/v1/user/{google_id} [get]

func (h *AuthServiceHandler) GetUserByGoogleID(c *fiber.Ctx) error {
	googleID := c.Params("google_id")
	resp, err := h.SVC.GetUserByGoogleID(googleID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

//@Id Update User
//@Summary Update User
//@Tags User
//@Accept json
//@Produce json
//@Param user body UseAccount true "User Object"
//@Success 200 {object} HttpResponse
//@Failure 400 {object} HttpResponse
//@Failure 500 {object} HttpResponse
//@Router api/v1/user [put]

func (h *AuthServiceHandler) UpdateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	// Ensure ID is set or get from context if authenticated
	// For now assuming body has ID
	resp, err := h.SVC.UpdateUser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

func (h *AuthServiceHandler) DeleteUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	resp, err := h.SVC.DeleteUser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
