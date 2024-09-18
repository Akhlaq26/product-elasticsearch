package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"product-elasticsearch/internal/common"
	"product-elasticsearch/internal/config"
	"product-elasticsearch/internal/rest"
	"time"

	_ "product-elasticsearch/docs"

	"github.com/gofiber/fiber/v3"
	fiberlog "github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

// @title Product Elastic Search Skill-Test
// @version 1.0
// @description This is a swagger for the service
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.Init()
	es := config.InitES(cfg)
	// create new fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			// Fiber provides a way to retrieve the actual status code from an error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Set the status code
			c.Status(code)

			// Return a JSON response with the default error message
			return c.JSON(common.ErrorResponse{
				Error: err.Error(),
			})
		},
	})

	app.Use(logger.New(logger.Config{}), recover.New())
	rest.RegisterRoute(cfg, app, es)

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Start the server
	go func() {
		port := fmt.Sprintf(":%s", cfg.ServerPort)
		if err := app.Listen(port); err != nil {
			fiberlog.Errorf("Error starting server: %s\n", err)
			quit <- os.Interrupt
		}
	}()

	// Wait for the OS signal
	<-quit
	fiberlog.Info("Shutting down server...")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Gracefully shut down the server
	if err := app.ShutdownWithContext(ctx); err != nil {
		fiberlog.Errorf("Error shutting down server: %s\n", err)
	}

	fiberlog.Infof("Server stopped")
}
