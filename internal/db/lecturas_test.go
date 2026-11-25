package db

import (
	"regexp"
	"testing"
)

func TestGuardarLecturaFijaLaHora(t *testing.T) {
	cols := regexp.MustCompile(`\(\s*"timestamp",`)
	vals := regexp.MustCompile(`VALUES \(now\(\),`)
	if !cols.MatchString(sqlGuardarLectura) || !vals.MatchString(sqlGuardarLectura) {
		t.Fatal("el INSERT debe guardar \"timestamp\" con now(); sin eso la lectura queda sin hora")
	}
}
