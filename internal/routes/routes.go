package routes

import (
	"github.com/iskhakmuhamad/todo-api/internal/handler"
	"github.com/iskhakmuhamad/todo-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	authHandler *handler.AuthHandler,
	todoHandler *handler.TodoHandler,
	categoryHandler *handler.CategoryHandler,
	authMiddleware *middleware.AuthMiddleware,
) {
	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Todo API is running",
		})
	})

	api := app.Group("/api/v1")

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/logout", authHandler.Logout)

	// Protected routes
	protected := api.Group("", authMiddleware.ValidateJWT)

	// Category routes
	categories := protected.Group("/categories")
	categories.Post("/", categoryHandler.Create)
	categories.Get("/", categoryHandler.GetAll)
	categories.Get("/:id", categoryHandler.GetByID)
	categories.Put("/:id", categoryHandler.Update)
	categories.Delete("/:id", categoryHandler.Delete)

	// Todo routes
	todos := protected.Group("/todos")
	todos.Post("/", todoHandler.Create)
	todos.Get("/", todoHandler.GetAll)
	todos.Get("/:id", todoHandler.GetByID)
	todos.Put("/:id", todoHandler.Update)
	todos.Delete("/:id", todoHandler.Delete)
	todos.Patch("/:id/toggle", todoHandler.ToggleStatus)
}
