package main

import (
	"restful_api_mongodb_fiber/configs"
	"restful_api_mongodb_fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)

	app.Listen(":7000")
}
