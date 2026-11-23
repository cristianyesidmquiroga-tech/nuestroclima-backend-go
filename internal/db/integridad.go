package db

import (
	"context"
	"time"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/modelos"
)

func (a *Almacen) LecturasDelDia(ctx context.Context, inicio, fin time.Time) ([]modelos.Lectura, error) {
	rows, err := a.pool.Query(ctx,
		`SELECT `+colsLectura+` FROM lecturas
		 WHERE "timestamp" >= $1 AND "timestamp" < $2
		 ORDER BY "timestamp" ASC`, inicio, fin)
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
