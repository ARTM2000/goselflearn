package controllers

import (
	"fmt"

	"github.com/ARTM2000/goselflearn/internal/common"
	"github.com/ARTM2000/goselflearn/internal/controllers/dto"
	"github.com/ARTM2000/goselflearn/internal/models"
	"github.com/ARTM2000/goselflearn/internal/services"

	"github.com/gofiber/fiber/v2"
)

// @Summery     Register user
// @Description Register new user
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       body               body     dto.UserRegister         true "query params"
// @Success     200                {object} response.RegisterUser
// @Failure     400                {object} response.BaseError
// @Failure     422                {object} response.ValidationError
// @Router      /api/auth/register [post]
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

// @Summery     Login user
// @Description Login user
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       body            body     dto.UserLogin              true "query params"
// @Success     200             {object} response.LoginUser
// @Failure     401             {object} response.UnauthorizedError
// @Failure     400             {object} response.BaseError
// @Failure     422             {object} response.ValidationError
// @Router      /api/auth/login [post]
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

// @Summery     Get user info
// @Description Get user all information
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       Authorization header   string                     true "bearer access token"
// @Success     200           {object} response.GetUserInfo
// @Failure     401           {object} response.UnauthorizedError
// @Router      /api/user/me  [get]
func GetUserInfo(c *fiber.Ctx) error {
	userData := c.Locals("user").(*models.User)

	return c.Status(fiber.StatusOK).JSON(common.FormatResponse(common.ResponseData{
		TrackId: c.GetRespHeader(fiber.HeaderXRequestID),
		Data: fiber.Map{
			"user": userData,
		},
	}))
}
