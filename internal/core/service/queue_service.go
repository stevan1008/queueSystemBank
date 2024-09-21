package service

import (
    "sort"
    "github.com/stevan1008/queueSystemBank/internal/core/domain"
    "github.com/stevan1008/queueSystemBank/internal/core/port"
)

type InMemoryQueue struct {
    queue domain.Queue
}

func NewInMemoryQueue() port.Queue {
    return &InMemoryQueue{
        queue: domain.Queue{
            Clients: []domain.Client{},
        },
    }
}

func (q *InMemoryQueue) AddClient(client domain.Client) {
    q.queue.Clients = append(q.queue.Clients, client)
    sort.SliceStable(q.queue.Clients, func(i, j int) bool {
        if q.queue.Clients[i].Priority == q.queue.Clients[j].Priority {
            return q.queue.Clients[i].ArrivalTime.Before(q.queue.Clients[j].ArrivalTime)
        }
        return q.queue.Clients[i].Priority > q.queue.Clients[j].Priority
    })
}

func (q *InMemoryQueue) ProcessClient() domain.Client {
    if len(q.queue.Clients) == 0 {
        return domain.Client{}
    }
    client := q.queue.Clients[0]
    q.queue.Clients = q.queue.Clients[1:]
    return client
}

func (q *InMemoryQueue) IsEmpty() bool {
    return len(q.queue.Clients) == 0
}

func (q *InMemoryQueue) GetClients() []domain.Client {
	return q.queue.Clients
}