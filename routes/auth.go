package routes

import (
	c "apigateway/controllers"

	"github.com/gofiber/fiber/v2"
)

func auth(app *fiber.App) {
	app.Get("/signin", c.SignIn)
	app.Get("/signup", c.SignUp)
	app.Get("/refreshtoken", c.RefreshToken)
	app.Get("/signout", c.SignOut)
	app.Get("/sessions", c.Sessions)
}
