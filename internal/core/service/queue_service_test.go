package service

import (
	"testing"
	"time"
	"github.com/stevan1008/queueSystemBank/internal/core/domain"

	"github.com/stretchr/testify/assert"
)

func TestQueueService(t *testing.T) {
	queue := NewInMemoryQueue()

	client1 := domain.Client{
		ID:          "1",
		Name:        "Marlon Muete",
		Priority:    domain.PriorityNormal, // Prioridad Normal (1)
		ArrivalTime: time.Now(),
	}

	client2 := domain.Client{
		ID:          "2",
		Name:        "Jerson Monterroso",
		Priority:    domain.PriorityBasic, // Prioridad Básica (0)
		ArrivalTime: time.Now().Add(1 * time.Second), // Llegada después
	}

	queue.AddClient(client1)
	queue.AddClient(client2)

	client := queue.ProcessClient()
	assert.Equal(t, "1", client.ID, "El cliente con mayor prioridad debería ser procesado primero")
	assert.Equal(t, domain.PriorityNormal, client.Priority, "El cliente debería tener prioridad Normal")

	client = queue.ProcessClient()
	assert.Equal(t, "2", client.ID, "El siguiente cliente debería ser procesado en orden")
	assert.Equal(t, domain.PriorityBasic, client.Priority, "El cliente debería tener prioridad Básica")
}