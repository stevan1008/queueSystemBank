package util

import (
	"time"

	"golang.org/x/exp/rand"
)

// RandomDelay genera un retraso aleatorio entre min y max segundos.
// Se usará para simular la llegada de clientes en tiempos aleatorios.
func RandomDelay(min int, max int) time.Duration {
    if min >= max {
        return time.Duration(min) * time.Second
    }
    randomSeconds := rand.Intn(max-min) + min
    return time.Duration(randomSeconds) * time.Second
}

// TimeNow retorna la hora actual.
// Se usará para estandarizar cómo se captura el tiempo en la aplicación.
func TimeNow() time.Time {
    return time.Now()
}

// FormatDuration formatea la duración en un string entendible (ejemplo: "2m30s")
// Será útil para mostrar el tiempo que un cliente ha esperado en la cola.
func FormatDuration(d time.Duration) string {
    return d.String()
}

// ElapsedTimeSince calcula el tiempo transcurrido desde una hora dada.
// Se usará para calcular cuánto tieempo ha esperado un cliente en la cola.
func ElapsedTimeSince(t time.Time) time.Duration {
    return time.Since(t)
}