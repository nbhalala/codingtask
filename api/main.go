/*
 * Main package for executable commands
 * References:
 * 1. Fiber: https://docs.gofiber.io/api/app
 */

package main

// Importing packages from remote modules
import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nbhalala/codingtask/routes"
	"github.com/joho/godotenv"
)

// Routing: Fiber's component Feature
// Function Signature: app.Method(path string, ...func(*fiber.Ctx))
func SetupRoutes(app *fiber.App) {
	// Routes
	app.Get("/:url", routes.urlResolve)
	app.Post("/api/v1", routes.urlShorten)
}

// main function
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()
	app.Use(logger.New())
	SetupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
