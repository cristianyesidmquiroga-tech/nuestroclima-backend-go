package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func servidorQueResponde(codigo int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(codigo)
	}))
}

func TestRevisarSaludSanaDevuelveCero(t *testing.T) {
	srv := servidorQueResponde(http.StatusOK)
	defer srv.Close()
	if got := revisarSalud(srv.URL, srv.Client()); got != 0 {
		t.Fatalf("esperaba 0, llego %d", got)
	}
}

func TestRevisarSaludSinBaseDevuelveUno(t *testing.T) {
	srv := servidorQueResponde(http.StatusServiceUnavailable)
	defer srv.Close()
	if got := revisarSalud(srv.URL, srv.Client()); got != 1 {
		t.Fatalf("esperaba 1, llego %d", got)
	}
}

func TestRevisarSaludSinServidorDevuelveUno(t *testing.T) {
	srv := servidorQueResponde(http.StatusOK)
	url := srv.URL
	srv.Close()
	if got := revisarSalud(url, http.DefaultClient); got != 1 {
		t.Fatalf("esperaba 1, llego %d", got)
	}
}

func TestSoloElArgumentoSaludActivaLaRevision(t *testing.T) {
	if !esRevisionDeSalud([]string{"/api", "salud"}) {
		t.Fatal("salud deberia activar la revision")
	}
	if esRevisionDeSalud([]string{"/api"}) || esRevisionDeSalud([]string{"/api", "otro"}) {
		t.Fatal("sin argumento salud debe arrancar el servidor")
	}
}

func TestURLSaludUsaElPuertoConfigurado(t *testing.T) {
	t.Setenv("APP_PORT", "9090")
	if got := urlSalud(); got != "http://127.0.0.1:9090/salud" {
		t.Fatalf("url inesperada: %s", got)
	}
}
