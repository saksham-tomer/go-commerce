package routes

import (
	controllers "github.com/saksham-tomer/go-commerce/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
    app.Post("/user", controllers.CreateUser)
    app.Get("/user/:userId", controllers.GetAUser)
    app.Put("/user/:userId", controllers.EditAUser)
    app.Delete("/user/:userId", controllers.DeleteAUser)
    app.Get("/users", controllers.GetAllUsers)
}
