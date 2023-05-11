package routes

import (
	c "apigateway/controllers"

	"github.com/gofiber/fiber/v2"
)

func auth(app *fiber.App) {
	app.Get("/signin", c.SignIn)
	app.Get("/singup", c.SingUp)
	app.Get("/refreshtoken", c.RefreshToken)
}
