package port


import "github.com/stevan1008/queueSystemBank/internal/core/domain"

type Queue interface {
    AddClient(client domain.Client)
    ProcessClient() domain.Client
    IsEmpty() bool
	GetClients() []domain.Client 
}