package api

import (
	"context"
	"net/http"
	"time"
)

func (m *Manejador) Stats(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 8*time.Second)
	defer cancel()

	s, err := m.datos.StatsDia(ctx)
	if err != nil {
		escribirError(w, http.StatusInternalServerError, "no se pudieron calcular las estadisticas")
		return
	}
	escribirJSON(w, http.StatusOK, s)
}

func (m *Manejador) StatsComparativas(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	s, err := m.datos.StatsComparativas(ctx)
	if err != nil {
		escribirError(w, http.StatusInternalServerError, "no se pudieron calcular las estadisticas")
		return
	}
	escribirJSON(w, http.StatusOK, s)
}
