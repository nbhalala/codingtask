package main

import(
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nbhalala/codingtask/api/routes"
	"github.com/joho/godotenv"
)

func createRoutes(app *fiber.App){
	app.Get("/:url", routes.urlResolve)
	app.Post("/api/v1", routes.urlShorten)
}

func main(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	createRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
