package controllers

import (
	"fmt"
	"goselflearn/internal/common"
	"goselflearn/internal/controllers/dto"
	"goselflearn/internal/models"
	"goselflearn/internal/services"

	"github.com/gofiber/fiber/v2"
)

func CreatePostForUser(c *fiber.Ctx) error {
	var createPostData dto.CreatePost

	if err := c.BodyParser(&createPostData); err != nil {
		fmt.Println(err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := createPostData.Validate()
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

	userData := c.Locals("user").(*models.User)

	postService := services.NewPostService()
	post, cErr := postService.CreatePost(&createPostData, userData.ID)
	if cErr != nil {
		return fiber.NewError(fiber.StatusInternalServerError, cErr.Error())
	}

	return c.Status(fiber.StatusOK).JSON(common.FormatResponse(common.ResponseData{
		Message: common.MessageNewPostCreated,
		Data:    fiber.Map{"post": post},
		TrackId: c.GetRespHeader(fiber.HeaderXRequestID),
	}))
}

func GetUserPosts(c *fiber.Ctx) error {
	userData := c.Locals("user").(*models.User)

	postService := services.NewPostService()
	posts, err := postService.FindUserPosts(userData.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(common.FormatResponse(common.ResponseData{
		Message: common.MessageNewPostCreated,
		Data:    fiber.Map{"posts": posts},
		TrackId: c.GetRespHeader(fiber.HeaderXRequestID),
	}))
}
