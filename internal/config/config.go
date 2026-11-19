package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Puerto         string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPass         string
	DBName         string
	StationPasskey string
	OrigenesCORS   []string
}

func Cargar() (*Config, error) {
	c := &Config{
		Puerto:         valor("APP_PORT", "8080"),
		DBHost:         valor("DB_HOST", "localhost"),
		DBPort:         valor("DB_PORT", "5432"),
		DBUser:         os.Getenv("DB_USER"),
		DBPass:         os.Getenv("DB_PASS"),
		DBName:         valor("DB_NAME", "nuestroclima"),
		StationPasskey: os.Getenv("STATION_PASSKEY"),
		OrigenesCORS:   lista(os.Getenv("ALLOWED_ORIGINS")),
	}

	var faltan []string
	if c.DBUser == "" {
		faltan = append(faltan, "DB_USER")
	}
	if c.DBPass == "" {
		faltan = append(faltan, "DB_PASS")
	}
	if len(faltan) > 0 {
		return nil, fmt.Errorf("faltan variables obligatorias: %s", strings.Join(faltan, ", "))
	}
	return c, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
}

func valor(clave, porDefecto string) string {
	if v := os.Getenv(clave); v != "" {
		return v
	}
	return porDefecto
}

func lista(v string) []string {
	if strings.TrimSpace(v) == "" {
		return nil
	}
	partes := strings.Split(v, ",")
	var r []string
	for _, p := range partes {
		if s := strings.TrimSpace(p); s != "" {
			r = append(r, s)
		}
	}
	return r
}
