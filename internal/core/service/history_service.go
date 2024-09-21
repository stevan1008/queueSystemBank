package service

import (
    "time"
    "github.com/stevan1008/queueSystemBank/internal/core/domain"
    "github.com/stevan1008/queueSystemBank/internal/core/port"
)

type InMemoryHistory struct {
    clients []domain.Client
}

func NewInMemoryHistory() port.History {
    return &InMemoryHistory{
        clients: []domain.Client{},
    }
}

func (h *InMemoryHistory) RegisterClient(client domain.Client, waitTime time.Duration) {
    h.clients = append(h.clients, client)
}

func (h *InMemoryHistory) GetHistory() []domain.Client {
    return h.clients
}