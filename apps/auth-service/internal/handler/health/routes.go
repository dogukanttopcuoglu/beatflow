package health

import (
	"github.com/dogukanttopcuoglu/beatflow/apps/auth-service/internal/transport/httpadapter"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, handler *CheckHandler) {
	router.Get("/healthz", httpadapter.Handle[CheckRequest, CheckResponse](handler))
	router.Get("/readyz", httpadapter.Handle[CheckRequest, CheckResponse](handler))
}
