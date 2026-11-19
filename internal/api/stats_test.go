package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/config"
	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/modelos"
)

func pedir(t *testing.T, ruta string, d datosMock) *httptest.ResponseRecorder {
	t.Helper()
	m := NuevoManejador(&config.Config{}, d)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, ruta, nil)
	Router(&config.Config{}, m).ServeHTTP(rec, req)
	return rec
}

func TestStatsDia(t *testing.T) {
	tmax := 28.0
	rec := pedir(t, "/api/stats", datosMock{statsDia: &modelos.StatsDia{TempMax: &tmax}})
	if rec.Code != http.StatusOK {
		t.Fatalf("codigo = %d, quiero 200", rec.Code)
	}
	var s modelos.StatsDia
	if err := json.NewDecoder(rec.Body).Decode(&s); err != nil {
		t.Fatalf("json: %v", err)
	}
	if s.TempMax == nil || *s.TempMax != 28.0 {
		t.Fatalf("temp_max = %v, quiero 28", s.TempMax)
	}
}

func TestStatsDiaError(t *testing.T) {
	rec := pedir(t, "/api/stats", datosMock{statsDiaErr: errors.New("db")})
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("codigo = %d, quiero 500", rec.Code)
	}
}

func TestStatsComparativas(t *testing.T) {
	d := datosMock{comparativas: &modelos.StatsComparativas{
		Mensual: []modelos.StatsMensual{{Mes: "2026-11"}},
		Semanal: []modelos.StatsSemanal{{Semana: "2026-47"}},
	}}
	rec := pedir(t, "/api/stats/comparative", d)
	if rec.Code != http.StatusOK {
		t.Fatalf("codigo = %d, quiero 200", rec.Code)
	}
	var s modelos.StatsComparativas
	if err := json.NewDecoder(rec.Body).Decode(&s); err != nil {
		t.Fatalf("json: %v", err)
	}
	if len(s.Mensual) != 1 || len(s.Semanal) != 1 {
		t.Fatalf("mensual=%d semanal=%d, quiero 1 y 1", len(s.Mensual), len(s.Semanal))
	}
}
