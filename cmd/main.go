package main

import (
	"log"

	"github.com/iskhakmuhamad/todo-api/internal/config"
	"github.com/iskhakmuhamad/todo-api/internal/handler"
	"github.com/iskhakmuhamad/todo-api/internal/middleware"
	"github.com/iskhakmuhamad/todo-api/internal/repository"
	"github.com/iskhakmuhamad/todo-api/internal/routes"
	"github.com/iskhakmuhamad/todo-api/internal/seeder"
	"github.com/iskhakmuhamad/todo-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	db := config.ConnectDB(cfg)

	// Auto migrate
	config.AutoMigrate(db)

	// Run seeder if enabled
	if cfg.RunSeeder {
		log.Println("Running database seeder...")
		if err := seeder.RunSeeder(db); err != nil {
			log.Printf("Seeder error: %v", err)
		} else {
			log.Println("Database seeder completed successfully")
		}
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	todoRepo := repository.NewTodoRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	todoService := service.NewTodoService(todoRepo)
	categoryService := service.NewCategoryService(categoryRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)
	todoHandler := handler.NewTodoHandler(todoService)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Global middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Setup routes
	routes.SetupRoutes(app, authHandler, todoHandler, categoryHandler, authMiddleware)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
