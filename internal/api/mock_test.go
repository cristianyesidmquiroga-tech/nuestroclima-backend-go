package api

import (
	"context"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/modelos"
)

type datosMock struct {
	saludErr       error
	ultima         *modelos.Lectura
	ultimaErr      error
	historial      []modelos.Lectura
	historialErr   error
	historialHoras *int
}

func (d datosMock) Salud(context.Context) error { return d.saludErr }

func (d datosMock) UltimaLectura(context.Context) (*modelos.Lectura, error) {
	return d.ultima, d.ultimaErr
}

func (d datosMock) Historial(_ context.Context, horas int) ([]modelos.Lectura, error) {
	if d.historialHoras != nil {
		*d.historialHoras = horas
	}
	return d.historial, d.historialErr
}
