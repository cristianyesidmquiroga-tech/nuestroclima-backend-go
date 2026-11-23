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

func pedirIntegridad(t *testing.T, fecha string, d datosMock) *httptest.ResponseRecorder {
	t.Helper()
	m := NuevoManejador(&config.Config{}, d)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/integridad/"+fecha, nil)
	Router(&config.Config{}, m).ServeHTTP(rec, req)
	return rec
}

func TestIntegridadDevuelveHash(t *testing.T) {
	temp := 20.0
	d := datosMock{lecturasDia: []modelos.Lectura{{ID: 1, TempExterior: &temp}}}
	rec := pedirIntegridad(t, "20261122", d)
	if rec.Code != http.StatusOK {
		t.Fatalf("codigo = %d, quiero 200", rec.Code)
	}
	var r modelos.IntegridadDia
	if err := json.NewDecoder(rec.Body).Decode(&r); err != nil {
		t.Fatalf("json: %v", err)
	}
	if r.FechaNum != 20261122 || r.Cantidad != 1 || len(r.Hash) != 64 {
		t.Errorf("respuesta invalida: %+v", r)
	}
}

func TestIntegridadFechaInvalida(t *testing.T) {
	rec := pedirIntegridad(t, "ayer", datosMock{})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("codigo = %d, quiero 400", rec.Code)
	}
}

func TestIntegridadErrorDeDatos(t *testing.T) {
	rec := pedirIntegridad(t, "20261122", datosMock{lecturasDiaErr: errors.New("db caida")})
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("codigo = %d, quiero 500", rec.Code)
	}
}

func TestIntegridadSinLecturasHashVacio(t *testing.T) {
	rec := pedirIntegridad(t, "20261122", datosMock{})
	var r modelos.IntegridadDia
	_ = json.NewDecoder(rec.Body).Decode(&r)
	if r.Cantidad != 0 || r.Hash != "" {
		t.Errorf("sin lecturas debe dar cantidad 0 y hash vacio: %+v", r)
	}
}
