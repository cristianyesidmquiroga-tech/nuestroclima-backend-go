package api

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

const (
	horasHistorialPorDefecto = 24
	horasHistorialMax        = 168
)

func (m *Manejador) Historial(w http.ResponseWriter, r *http.Request) {
	horas := horasHistorialPorDefecto
	if v := r.URL.Query().Get("horas"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			horas = n
		}
	}
	if horas < 1 {
		horas = 1
	}
	if horas > horasHistorialMax {
		horas = horasHistorialMax
	}

	ctx, cancel := context.WithTimeout(r.Context(), 8*time.Second)
	defer cancel()

	lecturas, err := m.datos.Historial(ctx, horas)
	if err != nil {
		escribirError(w, http.StatusInternalServerError, "no se pudo leer el historial")
		return
	}
	escribirJSON(w, http.StatusOK, lecturas)
}
