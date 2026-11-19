package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/config"
	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/modelos"
)

func pedirHistorial(t *testing.T, query string, d datosMock) *httptest.ResponseRecorder {
	t.Helper()
	m := NuevoManejador(&config.Config{}, d)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/historial"+query, nil)
	Router(&config.Config{}, m).ServeHTTP(rec, req)
	return rec
}

func TestHistorialDevuelveSerie(t *testing.T) {
	temp := 18.0
	d := datosMock{historial: []modelos.Lectura{{ID: 1, TempExterior: &temp}, {ID: 2}}}
	rec := pedirHistorial(t, "", d)
	if rec.Code != http.StatusOK {
		t.Fatalf("codigo = %d, quiero 200", rec.Code)
	}
	var ls []modelos.Lectura
	if err := json.NewDecoder(rec.Body).Decode(&ls); err != nil {
		t.Fatalf("json: %v", err)
	}
	if len(ls) != 2 {
		t.Fatalf("len = %d, quiero 2", len(ls))
	}
}

func TestHistorialLimitaHoras(t *testing.T) {
	casos := map[string]int{"?horas=500": 168, "?horas=0": 1, "": 24, "?horas=48": 48}
	for query, esperado := range casos {
		var horas int
		pedirHistorial(t, query, datosMock{historialHoras: &horas})
		if horas != esperado {
			t.Errorf("%q -> horas = %d, quiero %d", query, horas, esperado)
		}
	}
}
