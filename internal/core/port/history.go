package port

import (
	"time"

	"github.com/stevan1008/queueSystemBank/internal/core/domain"
)

type History interface {
    RegisterClient(client domain.Client, waitTime time.Duration)
    GetHistory() []domain.Client
}