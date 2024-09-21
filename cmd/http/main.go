package main

import (
	"log"
	"github.com/stevan1008/queueSystemBank/internal/adapter/handler"
	"github.com/stevan1008/queueSystemBank/internal/adapter/router"
	"github.com/stevan1008/queueSystemBank/internal/core/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	queue := service.NewInMemoryQueue()
	history := service.NewInMemoryHistory()

	h := handler.NewHandler(queue, history)
	app := fiber.New()
	router.SetupRouter(app, h)

	port := ":8084"
	log.Printf("Servidor escuchando en %s", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}