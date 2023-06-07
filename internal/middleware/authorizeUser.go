package middleware

import (
	"fmt"
	"strings"

	"github.com/ARTM2000/goselflearn/internal/common"
	"github.com/ARTM2000/goselflearn/internal/services"

	"github.com/gofiber/fiber/v2"
)

func AuthorizeUser(c *fiber.Ctx) error {
	authHeader := c.Get(fiber.HeaderAuthorization)
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return fiber.NewError(fiber.StatusUnauthorized, common.MessageUnauthorized)
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		return fiber.NewError(fiber.StatusUnauthorized, common.MessageUnauthorized)
	}

	userService := services.NewUserService()
	userData, err := userService.VerifyUserAccessToken(tokenString)
	if err != nil {
		fmt.Printf("token verification error: %s\n", err.Error())
		return fiber.NewError(fiber.StatusUnauthorized, common.MessageUnauthorized)
	}

	c.Locals("user", userData)
	fmt.Printf("request authorized. user: %s\n", userData)
	return c.Next()
}
