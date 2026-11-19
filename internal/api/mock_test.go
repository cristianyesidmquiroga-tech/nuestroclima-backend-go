package api

import (
	"context"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/modelos"
)

type datosMock struct {
	saludErr  error
	ultima    *modelos.Lectura
	ultimaErr error
}

func (d datosMock) Salud(context.Context) error { return d.saludErr }

func (d datosMock) UltimaLectura(context.Context) (*modelos.Lectura, error) {
	return d.ultima, d.ultimaErr
}
