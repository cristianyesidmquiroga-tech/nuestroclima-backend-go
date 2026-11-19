package db

import (
	"context"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/modelos"
)

func (a *Almacen) StatsDia(ctx context.Context) (*modelos.StatsDia, error) {
	row := a.pool.QueryRow(ctx, `
		SELECT max(temp_exterior), min(temp_exterior),
		       max(humedad_exterior), min(humedad_exterior),
		       max(velocidad_viento), max(rafaga_viento),
		       max(lluvia_dia), max(uv_index)
		FROM lecturas
		WHERE "timestamp" >= now() - interval '24 hours'`)

	var s modelos.StatsDia
	err := row.Scan(&s.TempMax, &s.TempMin, &s.HumMax, &s.HumMin,
		&s.VientoMax, &s.RafagaMax, &s.LluviaTotal, &s.UVMax)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (a *Almacen) StatsComparativas(ctx context.Context) (*modelos.StatsComparativas, error) {
	res := &modelos.StatsComparativas{
		Mensual: []modelos.StatsMensual{},
		Semanal: []modelos.StatsSemanal{},
	}

	rowsM, err := a.pool.Query(ctx, `
		SELECT to_char("timestamp", 'YYYY-MM') AS mes, max(temp_exterior),
		       min(temp_exterior), max(lluvia_dia), sum(lluvia_hora)
		FROM lecturas GROUP BY mes ORDER BY mes DESC LIMIT 12`)
	if err != nil {
		return nil, err
	}
	for rowsM.Next() {
		var m modelos.StatsMensual
		if err := rowsM.Scan(&m.Mes, &m.TempMax, &m.TempMin,
			&m.LluviaMaxDiaria, &m.EstimacionLluviaTotal); err != nil {
			rowsM.Close()
			return nil, err
		}
		res.Mensual = append(res.Mensual, m)
	}
	rowsM.Close()
	if err := rowsM.Err(); err != nil {
		return nil, err
	}

	rowsS, err := a.pool.Query(ctx, `
		SELECT to_char("timestamp", 'IYYY-IW') AS semana, max(temp_exterior),
		       min(temp_exterior), max(lluvia_dia)
		FROM lecturas GROUP BY semana ORDER BY semana DESC LIMIT 12`)
	if err != nil {
		return nil, err
	}
	defer rowsS.Close()
	for rowsS.Next() {
		var s modelos.StatsSemanal
		if err := rowsS.Scan(&s.Semana, &s.TempMax, &s.TempMin,
			&s.LluviaMaxDiaria); err != nil {
			return nil, err
		}
		res.Semanal = append(res.Semanal, s)
	}
	return res, rowsS.Err()
}
