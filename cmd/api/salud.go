package main

import (
	"net/http"
	"os"
	"time"
)

// La imagen distroless no trae shell ni curl, asi que el healthcheck del
// contenedor corre el mismo binario: `/api salud`.
func revisarSalud(url string, cliente *http.Client) int {
	resp, err := cliente.Get(url)
	if err != nil {
		return 1
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 1
	}
	return 0
}

func urlSalud() string {
	puerto := os.Getenv("APP_PORT")
	if puerto == "" {
		puerto = "8080"
	}
	return "http://127.0.0.1:" + puerto + "/salud"
}

func esRevisionDeSalud(args []string) bool {
	return len(args) > 1 && args[1] == "salud"
}

func correrRevisionDeSalud() {
	os.Exit(revisarSalud(urlSalud(), &http.Client{Timeout: 4 * time.Second}))
}
