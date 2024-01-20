package router

import (
	"system-gift-point/delivery/http/controller"
	"system-gift-point/domain"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, domain domain.Domain) {
    systemPointController := controller.NewSystemPointController(domain)

    app.Get("/", systemPointController.GetAllPoint)
    app.Post("/", systemPointController.SaveAction)
}
