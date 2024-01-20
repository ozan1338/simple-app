package controller

import (
	"motivational-reminder/delivery/http/dto/request"
	"motivational-reminder/domain"
	"motivational-reminder/shared/util"

	"github.com/gofiber/fiber/v2"
)

type QuoteController struct {
    domain domain.Domain
}

func NewQuoteController(domain domain.Domain) QuoteController {
    return QuoteController{
        domain: domain,
    }
}

func (this *QuoteController) GetQuote(ctx *fiber.Ctx) error {
    quote, err := this.domain.QuoteUsecase.GetRandomQuote()

    if err != nil {
        resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to fetch quote")
        return ctx.Status(statusCode).JSON(resp)
    }

    return ctx.Render("resource/views/home", fiber.Map{
        "Quote": quote.Text,
    })
}

func (this *QuoteController) SaveQuote(ctx *fiber.Ctx) error {
    var quote request.RequestQuoteDTO
    if err := ctx.BodyParser(&quote); err != nil {
        resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
        return ctx.Status(statusCode).JSON(resp)
    }

    if err := this.domain.QuoteUsecase.SaveQuote(quote.ToQuoteEntity()); err != nil {
        resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to save quote")
        return ctx.Status(statusCode).JSON(resp)
    }

    return ctx.Redirect("/")
}
