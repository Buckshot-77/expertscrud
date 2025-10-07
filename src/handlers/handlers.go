package handlers

import (
	"github.com/Buckshot-77/expertscrud/src/handlers/installment"
	"github.com/Buckshot-77/expertscrud/src/services"
	"github.com/gofiber/fiber/v2"
)

func NewHandlerContainer(router fiber.Router, serviceContainer services.ServiceContainer) {
	installment.NewInstallmentHandler(router, serviceContainer).SetRoutes()
}
