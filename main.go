package main

import (
	"errors"
	"fmt"
	"goselflearn/internal/common"
	"goselflearn/internal/initializers"
	"goselflearn/internal/routers"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func init() {
	initializers.LoadConfigurationFromDotEnv(".")
	initializers.DBConnect()
}

func main() {
	config := fiber.Config{
		CaseSensitive:                true,
		ServerHeader:                 "none",
		AppName:                      "GoSelfLearn",
		DisablePreParseMultipartForm: true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			// check that if error was an fiber NewError and got status code,
			// specify that in error handler
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

			return c.Status(code).JSON(common.FormatResponse(common.ResponseData{
				Message: err.Error(),
				TrackId: c.GetRespHeader(fiber.HeaderXRequestID),
				IsError: true,
			}))
		},
	}
	app := fiber.New(config)

	/**
	 * General configuration
	 */
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(helmet.New())
	// Here we only allow `application/json` content-type to treat a valid
	app.Use(func(c *fiber.Ctx) error {
		contentType := c.Get("Content-Type")
		if c.Method() != "GET" && contentType != "application/json" {
			return fiber.NewError(fiber.StatusBadRequest, "Request body must be in 'application/json' format")
		}
		return c.Next()
	})

	/**
	 * Specify routes
	 */
	apiRoutes := routers.APIRoutes{}
	app.Mount(apiRoutes.Path(), apiRoutes.Routes())

	// to gracefully shutdown fiber web server
	go shutdown(app)

	port := 3010
	if initializers.Config.Port != nil {
		port = *initializers.Config.Port
	}
	app.Listen(fmt.Sprintf(":%d", port))
}

func shutdown(app *fiber.App) {
	sigs := make(chan os.Signal, 1)

	fmt.Println("shutdown signal registered")

	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	<-sigs
	fmt.Println("\nshutdown signal received")
	app.Shutdown()
}
