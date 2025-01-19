package main

import (
	"github.com/saksham-tomer/go-commerce/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/saksham-tomer/go-commerce/config"
)

func main() {
    app := fiber.New()

    //run database
    config.ConnectDB()

    //routes
    routes.UserRoute(app) //add this

    app.Listen(":6000")
}

