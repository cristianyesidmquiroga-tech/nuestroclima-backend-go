package api

import (
	"context"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/config"
	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/modelos"
)

type Datos interface {
	Salud(ctx context.Context) error
	UltimaLectura(ctx context.Context) (*modelos.Lectura, error)
	Historial(ctx context.Context, horas int) ([]modelos.Lectura, error)
	StatsDia(ctx context.Context) (*modelos.StatsDia, error)
	StatsComparativas(ctx context.Context) (*modelos.StatsComparativas, error)
	GuardarLectura(ctx context.Context, l *modelos.Lectura) error
}

type Manejador struct {
	cfg   *config.Config
	datos Datos
}

func NuevoManejador(cfg *config.Config, datos Datos) *Manejador {
	return &Manejador{cfg: cfg, datos: datos}
}
