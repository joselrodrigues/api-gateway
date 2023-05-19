package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mileusna/useragent"
)

type CookieInfo struct {
	AccessToken  string
	RefreshToken string
	TtlAccess    time.Duration
	TtlRefresh   time.Duration
}

func SetAccessTokenCookie(c *fiber.Ctx, userAgentInfo useragent.UserAgent, cookieInfo CookieInfo) {
	if userAgentInfo.Desktop {
		cookie := new(fiber.Cookie)
		cookie.Name = "access_token"
		cookie.Value = cookieInfo.AccessToken
		cookie.HTTPOnly = true
		cookie.Expires = time.Now().Add(cookieInfo.TtlAccess)
		c.Cookie(cookie)
	}
}

func SetRefreshTokenCookie(c *fiber.Ctx, userAgentInfo useragent.UserAgent, cookieInfo CookieInfo) {
	if userAgentInfo.Desktop {
		cookie := new(fiber.Cookie)
		cookie.Name = "refresh_token"
		cookie.Value = cookieInfo.RefreshToken
		cookie.HTTPOnly = true
		cookie.Expires = time.Now().Add(cookieInfo.TtlRefresh)
		c.Cookie(cookie)
	}
}
