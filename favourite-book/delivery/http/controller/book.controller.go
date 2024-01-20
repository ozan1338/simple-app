package controller

import (
	"favourite-book/delivery/http/dto/request"
	"favourite-book/domain"
	"favourite-book/shared/util"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	domain domain.Domain
}

func NewBookController(domain domain.Domain) BookController {
	return BookController{
		domain: domain,
	}
}

func (this *BookController) GetAllBook(ctx *fiber.Ctx) error {
	books, err := this.domain.BookUsecase.GetBooks()

	if err != nil {
        resp , statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to fetch books")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Render("views/home/home", fiber.Map{
		"Books": books,
	})
}

func (this *BookController) AddBook(ctx *fiber.Ctx) error {
	var book request.RequestBookDTO
	if err := ctx.BodyParser(&book); err != nil {
        resp , statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := this.domain.BookUsecase.CreateBook(book.ToBookEntity()); err != nil {
        resp , statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to add book")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Redirect("/")
}
