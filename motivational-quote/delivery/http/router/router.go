package router

import (
    "github.com/gofiber/fiber/v2"
    "motivational-reminder/delivery/http/controller"
    "motivational-reminder/domain"
)

func NewRouter(app *fiber.App, domain domain.Domain) {
    quoteController := controller.NewQuoteController(domain)

    app.Get("/", quoteController.GetQuote)
    app.Post("/", quoteController.SaveQuote)
}
