package user

import "github.com/gofiber/fiber/v2"

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var u CreateUserReq

	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	res, err := h.Service.CreateUser(c.Context(), &u)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(res)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var user LoginUserReq
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	u, err := h.Service.Login(c.Context(), &user)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    u.accessToken,
		MaxAge:   3600,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	res := &LoginUserRes{ // no token
		ID:       u.ID,
		Username: u.Username,
	}

	return c.Status(200).JSON(res)
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		MaxAge:   -1,
		Path:     "",
		Domain:   "",
		Secure:   false,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.SendStatus(200)
}
