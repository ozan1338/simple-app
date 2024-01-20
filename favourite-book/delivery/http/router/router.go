package router

import (
    "favourite-book/delivery/http/controller"
    "favourite-book/domain"
    "github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, domain domain.Domain) {
    bookController := controller.NewBookController(domain)

    app.Get("/", bookController.GetAllBook)
    app.Post("/", bookController.AddBook)
}
