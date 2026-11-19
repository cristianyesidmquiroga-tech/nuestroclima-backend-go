package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/api"
	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/config"
	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/db"
)

func main() {
	cfg, err := config.Cargar()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	pool, err := db.Abrir(cfg.DSN())
	if err != nil {
		log.Fatalf("db: %v", err)
	}
	defer pool.Close()

	almacen := db.NuevoAlmacen(pool)
	manejador := api.NuevoManejador(cfg, almacen)

	srv := &http.Server{
		Addr:              ":" + cfg.Puerto,
		Handler:           api.Router(cfg, manejador),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       20 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		log.Printf("escuchando en :%s", cfg.Puerto)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("servidor: %v", err)
		}
	}()

	paro := make(chan os.Signal, 1)
	signal.Notify(paro, syscall.SIGINT, syscall.SIGTERM)
	<-paro

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("cierre: %v", err)
	}
	log.Println("servidor detenido")
}
