package handler

import (
	"github.com/stevan1008/queueSystemBank/internal/core/domain"
	"github.com/stevan1008/queueSystemBank/internal/core/port"
	"github.com/stevan1008/queueSystemBank/internal/core/util"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	queue   port.Queue
	history port.History
}

func NewHandler(queue port.Queue, history port.History) *Handler {
	return &Handler{
		queue:   queue,
		history: history,
	}
}

func (h *Handler) AddClient(c *fiber.Ctx) error {
	var client domain.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing client data")
	}

	if client.Priority < domain.PriorityBasic || client.Priority > domain.PriorityGerencial {
		return c.Status(fiber.StatusBadRequest).SendString("Prioridad inválida. Los valores deben estar entre 0 (Básico) y 3 (Gerencial).")
	}

	client.ArrivalTime = util.TimeNow()
	h.queue.AddClient(client)

	priorityLabel := util.GetPriorityLabel(client.Priority)
	return c.Status(fiber.StatusCreated).SendString("Cliente " + client.Name + " (" + priorityLabel + ") añadido a la cola")
}

func (h *Handler) ProcessClient(c *fiber.Ctx) error {
	if h.queue.IsEmpty() {
		return c.Status(fiber.StatusNoContent).SendString("No hay clientes en la cola")
	}

	client := h.queue.ProcessClient()
	waitTime := util.ElapsedTimeSince(client.ArrivalTime)
	h.history.RegisterClient(client, waitTime)

	priorityLabel := util.GetPriorityLabel(client.Priority)
	return c.Status(fiber.StatusOK).SendString("Cliente " + client.Name + " (" + priorityLabel + ") ha sido atendido (esperó " + util.FormatDuration(waitTime) + ")")
}

func (h *Handler) GetHistory(c *fiber.Ctx) error {
	history := h.history.GetHistory()

	if len(history) == 0 {
		return c.Status(fiber.StatusNoContent).SendString("No hay historial disponible")
	}

	return c.JSON(history)
}

func (h *Handler) GetQueue(c *fiber.Ctx) error {
	if h.queue.IsEmpty() {
		return c.Status(fiber.StatusNoContent).SendString("La cola está vacía")
	}

	return c.JSON(h.queue.GetClients())
}