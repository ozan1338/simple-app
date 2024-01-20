package controller

import (
    "github.com/gofiber/fiber/v2"
    "system-gift-point/delivery/http/dto/request"
    "system-gift-point/domain"
    "system-gift-point/shared/util"
)

type SystemPointController struct {
    domain domain.Domain
}

func NewSystemPointController(domain domain.Domain) SystemPointController {
    return SystemPointController{
        domain: domain,
    }
}

func (this *SystemPointController) GetAllPoint(ctx *fiber.Ctx) error {
    actions, totalPoint, err := this.domain.ActionUsecase.GetActions()

    if err != nil {
        resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to fetch actions")
        return ctx.Status(statusCode).JSON(resp)
    }

    return ctx.Render("views/home/home", fiber.Map{
        "Actions":    actions,
        "TotalPoint": totalPoint,
    })
}

func (this *SystemPointController) SaveAction(ctx *fiber.Ctx) error {
    var actions request.RequestsystemPointDTO
    if err := ctx.BodyParser(&actions); err != nil {
        resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
        return ctx.Status(statusCode).JSON(resp)
    }

    if err := this.domain.ActionUsecase.SaveAction(actions.ToActionEntity()); err != nil {
        resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to save action")
        return ctx.Status(statusCode).JSON(resp)
    }

    return ctx.Redirect("/")
}
