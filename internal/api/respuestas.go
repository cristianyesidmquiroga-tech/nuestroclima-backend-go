package api

import (
	"encoding/json"
	"net/http"
)

func escribirJSON(w http.ResponseWriter, estado int, datos any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(estado)
	if datos == nil {
		return
	}
	_ = json.NewEncoder(w).Encode(datos)
}

func escribirError(w http.ResponseWriter, estado int, mensaje string) {
	escribirJSON(w, estado, map[string]string{"error": mensaje})
}
