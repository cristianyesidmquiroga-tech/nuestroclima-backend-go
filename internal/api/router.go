package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/config"
)

func Router(cfg *config.Config, m *Manejador) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	origenes := cfg.OrigenesCORS
	if len(origenes) == 0 {
		origenes = []string{"http://localhost:5173"}
	}
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   origenes,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/salud", m.Salud)
	r.Get("/api/actual", m.Actual)
	r.Get("/api/historial", m.Historial)
	r.Get("/api/stats", m.Stats)
	r.Get("/api/stats/comparative", m.StatsComparativas)

	return r
}
