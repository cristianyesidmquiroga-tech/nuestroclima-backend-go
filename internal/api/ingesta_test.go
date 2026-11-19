package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/config"
	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/modelos"
)

func pedirIngesta(t *testing.T, clave, body string, d datosMock) *httptest.ResponseRecorder {
	t.Helper()
	cfg := &config.Config{StationPasskey: clave}
	m := NuevoManejador(cfg, d)
	req := httptest.NewRequest(http.MethodPost, "/data/report/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	Router(cfg, m).ServeHTTP(rec, req)
	return rec
}

func TestIngestaOK(t *testing.T) {
	var guardada *modelos.Lectura
	rec := pedirIngesta(t, "secreto", "PASSKEY=secreto&tempf=68&humidity=50",
		datosMock{guardada: &guardada})
	if rec.Code != http.StatusOK {
		t.Fatalf("codigo = %d, quiero 200", rec.Code)
	}
	if guardada == nil || guardada.TempExterior == nil || *guardada.TempExterior != 20 {
		t.Fatalf("no guardo la lectura calibrada: %+v", guardada)
	}
}

func TestIngestaPasskeyMala(t *testing.T) {
	rec := pedirIngesta(t, "secreto", "PASSKEY=otra&tempf=68", datosMock{})
	if rec.Code != http.StatusForbidden {
		t.Fatalf("codigo = %d, quiero 403", rec.Code)
	}
}

func TestIngestaSinPasskeyConfig(t *testing.T) {
	rec := pedirIngesta(t, "", "PASSKEY=cualquiera&tempf=68", datosMock{})
	if rec.Code != http.StatusForbidden {
		t.Fatalf("codigo = %d, quiero 403", rec.Code)
	}
}

func TestIngestaSinDatos(t *testing.T) {
	rec := pedirIngesta(t, "secreto", "", datosMock{})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("codigo = %d, quiero 400", rec.Code)
	}
}
