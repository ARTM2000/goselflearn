package controllers

import (
	"fmt"
	"goselflearn/internal/common"
	"goselflearn/internal/controllers/dto"
	"goselflearn/internal/models"
	"goselflearn/internal/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	newUser := dto.UserRegister{}
	if err := c.BodyParser(&newUser); err != nil {
		fmt.Println(err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate user data
	err := newUser.Validate()
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			common.FormatResponse(common.ResponseData{
				Message: err.Message,
				Data:    fiber.Map{"error": err},
				TrackId: c.GetRespHeader(fiber.HeaderXRequestID),
				IsError: true,
			}),
		)
	}

	// save user data to db
	userService := services.NewUserService()
	savedUser, registerError := userService.RegisterUser(&newUser)
	if registerError != nil {
		return fiber.NewError(
			fiber.StatusUnprocessableEntity, 
			registerError.Error(),
		)
	}

	// return new user information

	return c.Status(fiber.StatusOK).JSON(common.FormatResponse(common.ResponseData{
		Data: fiber.Map{
			"user": savedUser,
		},
		Message: common.MessageUserCreated,
		TrackId: c.GetRespHeader(fiber.HeaderXRequestID),
	}))
}

func LoginUser(c *fiber.Ctx) error {

	loginData := dto.UserLogin{}
	if err := c.BodyParser(&loginData); err != nil {
		fmt.Println(err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate user data
	err := loginData.Validate()
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			common.FormatResponse(common.ResponseData{
				Message: err.Message,
				Data:    fiber.Map{"error": err},
				TrackId: c.GetRespHeader(fiber.HeaderXRequestID),
				IsError: true,
			}),
		)
	}

	// save user data to db
	userService := services.NewUserService()
	accessToken, tokenErr := userService.LoginUser(&loginData)
	if tokenErr != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, tokenErr.Error())
	}

	return c.Status(fiber.StatusOK).JSON(common.FormatResponse(common.ResponseData{
		TrackId: c.GetRespHeader(fiber.HeaderXRequestID),
		Data: fiber.Map{
			"access_token": accessToken,
		},
		Message: common.MessageSuccessfulLogin,
	}))
}

func GetUserInfo(c *fiber.Ctx) error {
	userData := c.Locals("user").(*models.User)

	return c.Status(fiber.StatusOK).JSON(common.FormatResponse(common.ResponseData{
		TrackId: c.GetRespHeader(fiber.HeaderXRequestID),
		Data: fiber.Map{
			"user": userData,
		},
	}))
}
