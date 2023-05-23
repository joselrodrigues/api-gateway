package controllers

import (
	"apigateway/config"
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

	response, err := services.SignIn(c, client, infoUserAgent)

	if err != nil {
		fiberError := grcp.GrpcErrorToFiberError(err)
		return c.Status(fiberError.Code).JSON(fiber.Map{"error": fiberError.Message})
	}

	cfg := c.Locals("cfg").(config.Config)

	utils.SetAccessTokenCookie(c, infoUserAgent, utils.CookieInfo{
		AccessToken: response.AccessToken,
		TtlAccess:   cfg.AccessTokenExpiresIn,
	})
	utils.SetRefreshTokenCookie(c, infoUserAgent, utils.CookieInfo{
		RefreshToken: response.RefreshToken,
		TtlRefresh:   cfg.RefreshTokenExpiresIn,
	})

	return c.JSON(fiber.Map{"access_token": response.AccessToken, "refresh_token": response.RefreshToken})
}

func SignUp(c *fiber.Ctx) error {
	conn := grcp.Setup()
	client := grcp.NewAuthServiceClient(conn)
	userAgent := c.Get("User-Agent")
	infoUserAgent := useragent.Parse(userAgent)

	response, err := services.SignUp(c, client, infoUserAgent)

	if err != nil {
		fiberError := grcp.GrpcErrorToFiberError(err)
		return c.Status(fiberError.Code).JSON(fiber.Map{"error": fiberError.Message})
	}

	cfg := c.Locals("cfg").(config.Config)

	utils.SetAccessTokenCookie(c, infoUserAgent, utils.CookieInfo{
		AccessToken: response.AccessToken,
		TtlAccess:   cfg.AccessTokenExpiresIn,
	})
	utils.SetRefreshTokenCookie(c, infoUserAgent, utils.CookieInfo{
		RefreshToken: response.RefreshToken,
		TtlRefresh:   cfg.RefreshTokenExpiresIn,
	})

	return c.JSON(fiber.Map{"access_token": response.AccessToken, "refresh_token": response.RefreshToken})
}

func SignOut(c *fiber.Ctx) error {
	conn := grcp.Setup()
	client := grcp.NewAuthServiceClient(conn)
	userAgent := c.Get("User-Agent")
	infoUserAgent := useragent.Parse(userAgent)

	_, err := services.SignOut(c, client, infoUserAgent)

	if err != nil {
		fiberError := grcp.GrpcErrorToFiberError(err)
		return c.Status(fiberError.Code).JSON(fiber.Map{"error": fiberError.Message})
	}

	return c.JSON(fiber.Map{"message": "successfully signed out"})
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
