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

// La hora va explicita: la tabla importada de produccion no tiene DEFAULT en "timestamp".
const sqlGuardarLectura = `
	INSERT INTO lecturas ("timestamp", temp_exterior, temp_interior, humedad_exterior,
		humedad_interior, velocidad_viento, rafaga_viento, direccion_viento,
		lluvia_hora, lluvia_dia, lluvia_evento, presion_relativa, presion_absoluta,
		uv_index, radiacion_solar, punto_rocio, sensacion_termica, passkey)
	VALUES (now(),$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,'estacion')`

func (a *Almacen) GuardarLectura(ctx context.Context, l *modelos.Lectura) error {
	_, err := a.pool.Exec(ctx, sqlGuardarLectura,
		l.TempExterior, l.TempInterior, l.HumedadExterior, l.HumedadInterior,
		l.VelocidadViento, l.RafagaViento, l.DireccionViento, l.LluviaHora,
		l.LluviaDia, l.LluviaEvento, l.PresionRelativa, l.PresionAbsoluta,
		l.UVIndex, l.RadiacionSolar, l.PuntoRocio, l.SensacionTermica)
	return err
}

func (a *Almacen) Historial(ctx context.Context, horas int) ([]modelos.Lectura, error) {
	rows, err := a.pool.Query(ctx,
		`SELECT `+colsLectura+` FROM lecturas
		 WHERE "timestamp" >= now() - make_interval(hours => $1)
		 ORDER BY "timestamp" ASC`, horas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lecturas := make([]modelos.Lectura, 0)
	for rows.Next() {
		l, err := escanearLectura(rows)
		if err != nil {
			return nil, err
		}
		lecturas = append(lecturas, *l)
	}
	return lecturas, rows.Err()
}
