package handlers

import "github.com/gofiber/fiber/v2"

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

func (h *AuthServiceHamdler) CreateUser(c *fiber.Ctx) error {
	var user UseAccount
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	resp, err := h.SVC.CreateUser(&user)
	if err != nil {
		return h.Http.HttpResponseInternalServerErrorRequest(c, err)
	}
	return h.Http.HttpResponseok(c, resp)
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
func (h *AuthServiceHamdler) GetUserByEmail(c *fiber.Ctx) error {
	var user UseAccount
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	resp, err := h.SVC.GetUserByEmail(user.Email)
	if err != nil {
		return h.Http.HttpResponseInternalServerErrorRequest(c, err)
	}
	return h.Http.HttpResponseok(c, resp)
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

func (h *AuthServiceHamdler) GetUserByID(c *fiber.Ctx) error {
	var user UseAccount
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	resp, err := h.SVC.GetUserByID(user.ID)
	if err != nil {
		return h.Http.HttpResponseInternalServerErrorRequest(c, err)
	}
	return h.Http.HttpResponseok(c, resp)
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

func (h *AuthServiceHamdler) GetUserByGoogleID(c *fiber.Ctx) error {
	var user UseAccount
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	resp, err := h.SVC.GetUserByGoogleID(user.GoogleID)
	if err != nil {
		return h.Http.HttpResponseInternalServerErrorRequest(c, err)
	}
	return h.Http.HttpResponseok(c, resp)
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

func (h *AuthServiceHamdler) UpdateUser(c *fiber.Ctx) error {
	var user UseAccount
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	resp, err := h.SVC.UpdateUser(&user)
	if err != nil {
		return h.Http.HttpResponseInternalServerErrorRequest(c, err)
	}
	return h.Http.HttpResponseok(c, resp)
}

func (h *AuthServiceHamdler) DeleteUser(c *fiber.Ctx) error {
	var user UseAccount
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	resp, err := h.SVC.DeleteUser(&user)
	if err != nil {
		return h.Http.HttpResponseInternalServerErrorRequest(c, err)
	}
	return h.Http.HttpResponseok(c, resp)
}
