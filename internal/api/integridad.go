package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/integridad"
	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/modelos"
)

func (m *Manejador) Integridad(w http.ResponseWriter, r *http.Request) {
	inicio, fin, err := integridad.RangoDia(chi.URLParam(r, "fecha"))
	if err != nil {
		escribirError(w, http.StatusBadRequest, "fecha invalida, use AAAAMMDD o AAAA-MM-DD")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 8*time.Second)
	defer cancel()

	lecturas, err := m.datos.LecturasDelDia(ctx, inicio, fin)
	if err != nil {
		escribirError(w, http.StatusInternalServerError, "no se pudo leer las lecturas del dia")
		return
	}

	escribirJSON(w, http.StatusOK, modelos.IntegridadDia{
		Fecha:    inicio.Format("20060102"),
		FechaNum: integridad.NumeroFecha(inicio),
		Cantidad: len(lecturas),
		Hash:     integridad.Hash(lecturas),
	})
}
