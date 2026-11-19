package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/modelos"
)

const colsLectura = `id, "timestamp", temp_exterior, temp_interior, humedad_exterior,
	humedad_interior, velocidad_viento, rafaga_viento, direccion_viento, lluvia_hora,
	lluvia_dia, lluvia_evento, presion_relativa, presion_absoluta, uv_index,
	radiacion_solar, punto_rocio, sensacion_termica`

type filaScan interface {
	Scan(dest ...any) error
}

func escanearLectura(f filaScan) (*modelos.Lectura, error) {
	var l modelos.Lectura
	err := f.Scan(&l.ID, &l.Timestamp, &l.TempExterior, &l.TempInterior,
		&l.HumedadExterior, &l.HumedadInterior, &l.VelocidadViento, &l.RafagaViento,
		&l.DireccionViento, &l.LluviaHora, &l.LluviaDia, &l.LluviaEvento,
		&l.PresionRelativa, &l.PresionAbsoluta, &l.UVIndex, &l.RadiacionSolar,
		&l.PuntoRocio, &l.SensacionTermica)
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func (a *Almacen) UltimaLectura(ctx context.Context) (*modelos.Lectura, error) {
	row := a.pool.QueryRow(ctx,
		`SELECT `+colsLectura+` FROM lecturas ORDER BY "timestamp" DESC LIMIT 1`)
	l, err := escanearLectura(row)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return l, err
}
