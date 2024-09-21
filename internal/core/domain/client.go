package domain

import "time"

const (
	PriorityBasic    = iota
	PriorityNormal
	PriorityVIP
	PriorityGerencial
)

type Client struct {
	ID          string
	Name        string
	Priority    int
	ArrivalTime time.Time
}