package routes

import (
	c "apigateway/controllers"

	"github.com/gofiber/fiber/v2"
)

func auth(app *fiber.App) {
	app.Get("/signin", c.SignIn)
	// app.Get("/singup", s.SingUp)
	// app.Get("/refreshtoken", s.RefreshToken)
}
