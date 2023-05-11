package controllers

import (
	"apigateway/grcp"
	"apigateway/services"
	"apigateway/utils"

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

	utils.SetAccessTokenCookie(c, infoUserAgent, utils.CookieInfo{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	})
	utils.SetRefreshTokenCookie(c, infoUserAgent, utils.CookieInfo{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	})

	return c.JSON(fiber.Map{"access_token": response.AccessToken, "refresh_token": response.RefreshToken})
}

func SingUp(c *fiber.Ctx) error {
	conn := grcp.Setup()
	client := grcp.NewAuthServiceClient(conn)
	userAgent := c.Get("User-Agent")
	infoUserAgent := useragent.Parse(userAgent)

	response, err := services.SignUp(c, client)

	if err != nil {
		fiberError := grcp.GrpcErrorToFiberError(err)
		return c.Status(fiberError.Code).JSON(fiber.Map{"error": fiberError.Message})
	}

	utils.SetAccessTokenCookie(c, infoUserAgent, utils.CookieInfo{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	})
	utils.SetRefreshTokenCookie(c, infoUserAgent, utils.CookieInfo{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	})

	return c.JSON(fiber.Map{"access_token": response.AccessToken, "refresh_token": response.RefreshToken})
}

func RefreshToken(c *fiber.Ctx) error {
	conn := grcp.Setup()
	client := grcp.NewAuthServiceClient(conn)
	userAgent := c.Get("User-Agent")
	infoUserAgent := useragent.Parse(userAgent)

	response, err := services.RefreshToken(c, client, infoUserAgent)

	if err != nil {
		fiberError := grcp.GrpcErrorToFiberError(err)
		return c.Status(fiberError.Code).JSON(fiber.Map{"error": fiberError.Message})
	}

	return c.JSON(fiber.Map{"access_token": response.AccessToken})
}
