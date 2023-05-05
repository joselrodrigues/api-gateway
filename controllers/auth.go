package controllers

import (
	"apigateway/grcp"
	"apigateway/services"

	"github.com/gofiber/fiber/v2"
	"github.com/mileusna/useragent"
)

func SignIn(c *fiber.Ctx) error {

	conn := grcp.Setup()
	client := grcp.NewAuthServiceClient(conn)
	userAgent := c.Get("User-Agent")
	infoUserAgent := useragent.Parse(userAgent)

	response, err := services.SignIn(c, client)

	if err != nil {
		fiberError := grcp.GrpcErrorToFiberError(err)
		return c.Status(fiberError.Code).JSON(fiber.Map{"error": fiberError.Message})
	}

	if infoUserAgent.Desktop {
		cookieAccessToken := new(fiber.Cookie)
		cookieAccessToken.Name = "access_token"
		cookieAccessToken.Value = response.AccessToken
		cookieAccessToken.HTTPOnly = true
		c.Cookie(cookieAccessToken)

		cookieRefreshToken := new(fiber.Cookie)
		cookieRefreshToken.Name = "refresh_token"
		cookieRefreshToken.Value = response.RefreshToken
		cookieRefreshToken.HTTPOnly = true
		c.Cookie(cookieRefreshToken)
	}

	return c.JSON(fiber.Map{"access_token": response.AccessToken, "refresh_token": response.RefreshToken})
}
