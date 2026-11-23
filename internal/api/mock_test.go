package api

import (
	"context"
	"time"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/modelos"
)

type datosMock struct {
	saludErr        error
	ultima          *modelos.Lectura
	ultimaErr       error
	historial       []modelos.Lectura
	historialErr    error
	historialHoras  *int
	lecturasDia     []modelos.Lectura
	lecturasDiaErr  error
	statsDia        *modelos.StatsDia
	statsDiaErr     error
	comparativas    *modelos.StatsComparativas
	comparativasErr error
	guardarErr      error
	guardada        **modelos.Lectura
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

func (d datosMock) LecturasDelDia(_ context.Context, _, _ time.Time) ([]modelos.Lectura, error) {
	return d.lecturasDia, d.lecturasDiaErr
}

func (d datosMock) StatsDia(context.Context) (*modelos.StatsDia, error) {
	return d.statsDia, d.statsDiaErr
}

func (d datosMock) StatsComparativas(context.Context) (*modelos.StatsComparativas, error) {
	return d.comparativas, d.comparativasErr
}

func (d datosMock) GuardarLectura(_ context.Context, l *modelos.Lectura) error {
	if d.guardada != nil {
		*d.guardada = l
	}
	return d.guardarErr
}
