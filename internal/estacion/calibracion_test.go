package estacion

import (
	"net/url"
	"testing"
)

func deref(p *float64) float64 {
	if p == nil {
		return -999
	}
	return *p
}

func TestConversiones(t *testing.T) {
	casos := []struct {
		nombre string
		got    *float64
		quiero float64
	}{
		{"F->C 32", FahrenheitACelsius("32"), 0},
		{"F->C 68", FahrenheitACelsius("68"), 20},
		{"mph->kmh 10", MphAKmh("10"), 16.1},
		{"pulg->mm 1", PulgadasAMm("1"), 25.4},
		{"inHg->hPa 29.92", InhgAHpa("29.92"), 1013.2},
	}
	for _, c := range casos {
		if deref(c.got) != c.quiero {
			t.Errorf("%s = %v, quiero %v", c.nombre, deref(c.got), c.quiero)
		}
	}
	if FahrenheitACelsius("") != nil || MphAKmh("abc") != nil {
		t.Error("valor vacio/invalido deberia ser nil")
	}
}

func TestCompensacionTermica(t *testing.T) {
	temp, rad, viento := 30.6, 1453.0, 1.6
	// rad>1200 -> base 8, viento<3 -> factor 1 -> offset 8 -> 22.6
	got := AplicarCompensacionTermica(&temp, &rad, &viento)
	if deref(got) != 22.6 {
		t.Fatalf("compensada = %v, quiero 22.6", deref(got))
	}
	// Sin radiacion no hay compensacion
	if r := AplicarCompensacionTermica(&temp, nil, &viento); deref(r) != 30.6 {
		t.Fatalf("sin radiacion = %v, quiero 30.6", deref(r))
	}
	// Viento fuerte reduce el offset (factor 0.4 -> offset 3.2 -> 27.4)
	vientoFuerte := 12.0
	if r := AplicarCompensacionTermica(&temp, &rad, &vientoFuerte); deref(r) != 27.4 {
		t.Fatalf("viento fuerte = %v, quiero 27.4", deref(r))
	}
}

func TestConstruirLecturaCalibra(t *testing.T) {
	datos := url.Values{}
	datos.Set("tempf", "87.08") // ~30.6 C crudo
	datos.Set("solarradiation", "1453")
	datos.Set("windspeedmph", "1")
	datos.Set("humidity", "60")

	l := ConstruirLectura(datos)
	if l.TempExterior == nil || *l.TempExterior >= 30 {
		t.Fatalf("temp_exterior = %v, deberia bajar por compensacion", deref(l.TempExterior))
	}
	if l.HumedadExterior == nil || *l.HumedadExterior <= 60 {
		t.Fatalf("humedad = %v, deberia subir al corregir por temp", deref(l.HumedadExterior))
	}
}
