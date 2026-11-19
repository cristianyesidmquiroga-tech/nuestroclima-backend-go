package api

import (
	"context"
	"net/http"
	"time"
)

func (m *Manejador) Salud(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	estado, codigo, dbOK := "ok", http.StatusOK, true
	if err := m.datos.Salud(ctx); err != nil {
		estado, codigo, dbOK = "degradado", http.StatusServiceUnavailable, false
	}
	escribirJSON(w, codigo, map[string]any{"estado": estado, "db": dbOK})
}
