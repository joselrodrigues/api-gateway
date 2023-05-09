package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mileusna/useragent"
)

type CookieInfo struct {
	AccessToken  string
	RefreshToken string
}

func SetAccessTokenCookie(c *fiber.Ctx, userAgentInfo useragent.UserAgent, cookieInfo CookieInfo) {
	if userAgentInfo.Desktop {
		cookie := new(fiber.Cookie)
		cookie.Name = "access_token"
		cookie.Value = cookieInfo.AccessToken
		cookie.HTTPOnly = true
		c.Cookie(cookie)
	}
}

func SetRefreshTokenCookie(c *fiber.Ctx, userAgentInfo useragent.UserAgent, cookieInfo CookieInfo) {
	if userAgentInfo.Desktop {
		cookie := new(fiber.Cookie)
		cookie.Name = "refresh_token"
		cookie.Value = cookieInfo.RefreshToken
		cookie.HTTPOnly = true
		c.Cookie(cookie)
	}
}
