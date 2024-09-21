package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stevan1008/queueSystemBank/internal/core/service"
	"github.com/stretchr/testify/assert"
)

func TestAddClientHandler(t *testing.T) {
	queue := service.NewInMemoryQueue()
	history := service.NewInMemoryHistory()
	h := NewHandler(queue, history)

	app := fiber.New()
	app.Post("/clients", h.AddClient)

	clientData := `{"ID":"1", "Name":"Marlon Muete", "Priority":1}`

	req := httptest.NewRequest("POST", "/clients", bytes.NewBuffer([]byte(clientData)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	responseBody := buf.String()

	assert.Contains(t, responseBody, "Cliente Marlon Muete (Normal) añadido a la cola")
}