package services

import (
	pb "apigateway/protos"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mileusna/useragent"
	"google.golang.org/grpc/metadata"
)

func SignIn(c *fiber.Ctx, client pb.AuthServiceClient, infoUserAgent useragent.UserAgent) (*pb.Response, error) {
	email := c.Query("email")
	username := c.Query("username")
	password := c.Query("password")

	md := metadata.Pairs("user-device", fmt.Sprintf("%s(%s)", infoUserAgent.Name, infoUserAgent.OS))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	response, err := client.SignIn(ctx, &pb.Request{Email: email, Username: username, Password: password})

	return response, err

}

func SignUp(c *fiber.Ctx, client pb.AuthServiceClient, infoUserAgent useragent.UserAgent) (*pb.Response, error) {
	email := c.Query("email")
	username := c.Query("username")
	password := c.Query("password")

	md := metadata.Pairs("user-device", fmt.Sprintf("%s(%s)", infoUserAgent.Name, infoUserAgent.OS))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	response, err := client.SignUp(ctx, &pb.Request{Email: email, Username: username, Password: password})

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
