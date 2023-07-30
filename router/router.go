package router

import (
	"github.com/zumosik/jwt-auth-golang/internal/user"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func InitRouter(userHandler *user.Handler) {
	app = fiber.New()

	// User endpoints
	app.Post("/user/signup", userHandler.CreateUser)
	app.Post("/user/login", userHandler.Login)
	app.Get("/user/logout", userHandler.Logout)

}

func Start(addr string) error {
	return app.Listen(addr)
}
