package integridad

import (
	"testing"
	"time"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/modelos"
)

func pf(v float64) *float64 { return &v }

func TestRangoDia(t *testing.T) {
	for _, c := range []string{"20261122", "2026-11-22"} {
		ini, fin, err := RangoDia(c)
		if err != nil {
			t.Fatalf("%s: %v", c, err)
		}
		if ini.Year() != 2026 || ini.Month() != 11 || ini.Day() != 22 {
			t.Errorf("%s: inicio mal: %v", c, ini)
		}
		if fin.Sub(ini) != 24*time.Hour {
			t.Errorf("%s: el rango no es un dia: %v", c, fin.Sub(ini))
		}
	}
	if _, _, err := RangoDia("ayer"); err == nil {
		t.Error("se esperaba error con fecha invalida")
	}
}

func TestNumeroFecha(t *testing.T) {
	ini, _, _ := RangoDia("2026-11-22")
	if NumeroFecha(ini) != 20261122 {
		t.Errorf("num = %d, quiero 20261122", NumeroFecha(ini))
	}
}

func TestHashDeterministaEIndependienteDelOrden(t *testing.T) {
	ts := time.Date(2026, 11, 22, 10, 0, 0, 0, time.UTC)
	a := modelos.Lectura{Timestamp: ts, TempExterior: pf(20)}
	b := modelos.Lectura{Timestamp: ts.Add(time.Hour), TempExterior: pf(21)}

	h1 := Hash([]modelos.Lectura{a, b})
	h2 := Hash([]modelos.Lectura{b, a})
	if h1 == "" || h1 != h2 {
		t.Errorf("hash no determinista: %q vs %q", h1, h2)
	}
	if len(h1) != 64 {
		t.Errorf("sha256 hex debe medir 64, mide %d", len(h1))
	}

	c := modelos.Lectura{Timestamp: ts, TempExterior: pf(99)}
	if Hash([]modelos.Lectura{c, b}) == h1 {
		t.Error("el hash no cambio al cambiar los datos")
	}
	if Hash(nil) != "" {
		t.Error("sin lecturas el hash debe ser vacio")
	}
}
