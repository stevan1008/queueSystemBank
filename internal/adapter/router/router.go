package router

import (
	"github.com/stevan1008/queueSystemBank/internal/adapter/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, h *handler.Handler) {
	app.Post("/clients", h.AddClient)
	app.Post("/clients/next", h.ProcessClient)
	app.Get("/history", h.GetHistory)
	app.Get("/queue", h.GetQueue)
}