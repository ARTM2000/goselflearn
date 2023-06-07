package routers

import (
	"github.com/ARTM2000/goselflearn/internal/controllers"
	"github.com/ARTM2000/goselflearn/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type APIRoutes struct{}

func (a *APIRoutes) Path() string {
	return "/api"
}

func (a *APIRoutes) Routes() *fiber.App {
	api := fiber.New()

	api.Get("/healthz", controllers.GetAppHealthStatus).Name("health")

	api.Route("/auth", func(r fiber.Router) {
		r.Post("/register", controllers.RegisterUser)
		r.Post("/login", controllers.LoginUser)
	})

	api.Route("/user", func(r fiber.Router) {
		r.Use(middleware.AuthorizeUser)
		r.Get("/me", controllers.GetUserInfo)
	})

	api.Route("/posts", func(r fiber.Router) {
		r.Use(middleware.AuthorizeUser)
		r.Post("/", controllers.CreatePostForUser)
		r.Get("/me", controllers.GetUserPosts)
	})

	return api
}
