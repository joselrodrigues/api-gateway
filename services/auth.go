package services

import (
	pb "apigateway/protos"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/mileusna/useragent"
)

func SignIn(c *fiber.Ctx, client pb.AuthServiceClient) (*pb.Response, error) {
	email := c.Query("email")
	username := c.Query("username")
	password := c.Query("password")

	response, err := client.SignIn(context.Background(), &pb.Request{Email: email, Username: username, Password: password})

	return response, err

}

func SignUp(c *fiber.Ctx, client pb.AuthServiceClient) (*pb.Response, error) {
	email := c.Query("email")
	username := c.Query("username")
	password := c.Query("password")

	response, err := client.SignUp(context.Background(), &pb.Request{Email: email, Username: username, Password: password})

	return response, err

}

func RefreshToken(c *fiber.Ctx, client pb.AuthServiceClient, infoUserAgent useragent.UserAgent) (*pb.RefreshTokenResponse, error) {
	var refreshToken string

	if infoUserAgent.Desktop {
		refreshToken = c.Cookies("refresh_token")
	} else {
		refreshToken = c.Query("refresh_token")
	}

	response, err := client.RefreshToken(context.Background(), &pb.RefreshTokenRequest{RefreshToken: refreshToken})

	return response, err

}
