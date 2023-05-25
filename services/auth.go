package services

import (
	pb "apigateway/protos"
	"apigateway/utils"
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

func SignOut(c *fiber.Ctx, client pb.AuthServiceClient, infoUserAgent useragent.UserAgent) (*pb.DeleteRefreshTokenResponse, error) {
	var refreshToken string
	session_id := c.Query("session_id")

	if infoUserAgent.Desktop {
		refreshToken = c.Cookies("refresh_token")
	} else {
		refreshToken = c.Query("refresh_token")
	}

	response, err := client.SignOut(context.Background(), &pb.DeleteRefreshTokenRequest{RefreshToken: refreshToken, SessionId: session_id})

	if err == nil && session_id == "" {
		utils.DeleteCookie(c, infoUserAgent, "refresh_token")
		utils.DeleteCookie(c, infoUserAgent, "access_token")
	}

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

func Sessions(c *fiber.Ctx, client pb.AuthServiceClient, infoUserAgent useragent.UserAgent) (*pb.SessionResponse, error) {
	var refreshToken string
	var accessToken string

	if infoUserAgent.Desktop {
		refreshToken = c.Cookies("refresh_token")
		accessToken = c.Cookies("access_token")
	} else {
		refreshToken = c.Query("refresh_token")
		accessToken = c.Query("access_token")
	}

	response, err := client.Sessions(context.Background(), &pb.SessionRequest{RefreshToken: refreshToken, AccessToken: accessToken})

	return response, err
}
