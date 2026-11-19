package api

import (
	"context"
	"crypto/hmac"
	"net/http"
	"time"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/estacion"
)

// Ingesta recibe los datos crudos de la estacion (cada 60s), valida la PASSKEY
// en tiempo constante, calibra la lectura y la guarda.
func (m *Manejador) Ingesta(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		escribirError(w, http.StatusBadRequest, "datos invalidos")
		return
	}
	if len(r.Form) == 0 {
		escribirError(w, http.StatusBadRequest, "sin datos")
		return
	}

	clave := m.cfg.StationPasskey
	recibido := r.Form.Get("PASSKEY")
	if clave == "" || !hmac.Equal([]byte(recibido), []byte(clave)) {
		escribirError(w, http.StatusForbidden, "no autorizado")
		return
	}

	lectura := estacion.ConstruirLectura(r.Form)

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	if err := m.datos.GuardarLectura(ctx, lectura); err != nil {
		escribirError(w, http.StatusInternalServerError, "no se pudo guardar la lectura")
		return
	}
	escribirJSON(w, http.StatusOK, map[string]string{"estado": "ok"})
}
