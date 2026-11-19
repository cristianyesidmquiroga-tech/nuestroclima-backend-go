package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/config"
)

type datosFalsos struct{ err error }

func (d datosFalsos) Salud(context.Context) error { return d.err }

func peticionSalud(t *testing.T, err error) (int, map[string]any) {
	t.Helper()
	m := NuevoManejador(&config.Config{}, datosFalsos{err: err})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/salud", nil)
	Router(&config.Config{}, m).ServeHTTP(rec, req)

	var cuerpo map[string]any
	if e := json.NewDecoder(rec.Body).Decode(&cuerpo); e != nil {
		t.Fatalf("json: %v", e)
	}
	return rec.Code, cuerpo
}

func TestSaludOK(t *testing.T) {
	code, cuerpo := peticionSalud(t, nil)
	if code != http.StatusOK {
		t.Fatalf("codigo = %d, quiero 200", code)
	}
	if cuerpo["db"] != true {
		t.Fatalf("db = %v, quiero true", cuerpo["db"])
	}
}

func TestSaludDBCaida(t *testing.T) {
	code, cuerpo := peticionSalud(t, errors.New("sin conexion"))
	if code != http.StatusServiceUnavailable {
		t.Fatalf("codigo = %d, quiero 503", code)
	}
	if cuerpo["db"] != false {
		t.Fatalf("db = %v, quiero false", cuerpo["db"])
	}
}
