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

func pedirActual(t *testing.T, d datosMock) *httptest.ResponseRecorder {
	t.Helper()
	m := NuevoManejador(&config.Config{}, d)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/actual", nil)
	Router(&config.Config{}, m).ServeHTTP(rec, req)
	return rec
}

func TestActualConDato(t *testing.T) {
	temp := 21.5
	rec := pedirActual(t, datosMock{ultima: &modelos.Lectura{ID: 7, TempExterior: &temp}})
	if rec.Code != http.StatusOK {
		t.Fatalf("codigo = %d, quiero 200", rec.Code)
	}
	var l modelos.Lectura
	if err := json.NewDecoder(rec.Body).Decode(&l); err != nil {
		t.Fatalf("json: %v", err)
	}
	if l.TempExterior == nil || *l.TempExterior != 21.5 {
		t.Fatalf("temp_exterior = %v, quiero 21.5", l.TempExterior)
	}
}

func TestActualSinDatos(t *testing.T) {
	rec := pedirActual(t, datosMock{ultima: nil})
	if rec.Code != http.StatusNoContent {
		t.Fatalf("codigo = %d, quiero 204", rec.Code)
	}
}

func TestActualError(t *testing.T) {
	rec := pedirActual(t, datosMock{ultimaErr: errors.New("db caida")})
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("codigo = %d, quiero 500", rec.Code)
	}
}
