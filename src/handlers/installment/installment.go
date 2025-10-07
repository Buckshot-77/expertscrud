package installment

import (
	"net/http"

	"github.com/Buckshot-77/expertscrud/src/interfaces"
	"github.com/Buckshot-77/expertscrud/src/services"
	"github.com/Buckshot-77/expertscrud/src/structs"
	"github.com/gofiber/fiber/v2"
)

type InstallmentHandler struct {
	router             fiber.Router
	installmentService interfaces.InstallmentService
}

func NewInstallmentHandler(router fiber.Router, serviceContainer services.ServiceContainer) InstallmentHandler {
	return InstallmentHandler{
		router:             router,
		installmentService: serviceContainer.InstallmentService,
	}
}

func (ih InstallmentHandler) SetRoutes() {
	group := ih.router.Group("/installment")

	group.Post("", ih.CreateInstallment)
}

func (ih InstallmentHandler) CreateInstallment(c *fiber.Ctx) error {
	var body structs.Installment

	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": "error while creating installment",
		})
	}

	return ih.installmentService.Create(body)
}
