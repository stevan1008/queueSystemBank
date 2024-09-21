package util

import "github.com/stevan1008/queueSystemBank/internal/core/domain"

func GetPriorityLabel(priority int) string {
	switch priority {
	case domain.PriorityGerencial:
		return "Gerencial"
	case domain.PriorityVIP:
		return "VIP"
	case domain.PriorityNormal:
		return "Normal"
	default:
		return "Básico"
	}
}