package controllers

import (
	"github.com/ARTM2000/goselflearn/internal/common"

	"github.com/gofiber/fiber/v2"
)

func GetAppHealthStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(common.FormatResponse(common.ResponseData{
		TrackId: c.GetRespHeader(fiber.HeaderXRequestID),
		Message: "everything is operational",
		Data: fiber.Map{
			"running": true,
		},
	}))
}
