package api

import (
	"context"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/config"
)

type Datos interface {
	Salud(ctx context.Context) error
}

type Manejador struct {
	cfg   *config.Config
	datos Datos
}

func NuevoManejador(cfg *config.Config, datos Datos) *Manejador {
	return &Manejador{cfg: cfg, datos: datos}
}
