package api

import (
	"context"
	"net/http"
	"time"
)

func (m *Manejador) Actual(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	l, err := m.datos.UltimaLectura(ctx)
	if err != nil {
		escribirError(w, http.StatusInternalServerError, "no se pudo leer la ultima lectura")
		return
	}
	if l == nil {
		escribirJSON(w, http.StatusNoContent, nil)
		return
	}
	escribirJSON(w, http.StatusOK, l)
}
