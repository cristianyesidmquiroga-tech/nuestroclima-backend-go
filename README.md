# Nuestro Clima — Backend (Go)

API del dashboard meteorológico de la finca: recibe las lecturas de la estación
y las sirve al frontend. Es una implementación del contrato común
(`contrato/openapi.yaml`); la idea es montar el mismo contrato en varias
tecnologías y comparar cuál rinde mejor.

<details>
<summary><b>Read this in English</b></summary>

Weather dashboard backend for the farm: it receives the station's readings and
serves them to the frontend. It implements a shared contract
(`contrato/openapi.yaml`) so the same API can be built in several technologies
and compared.

</details>

## Stack

Go 1.27 · go-chi · pgx · PostgreSQL · Docker.

## Desarrollo

```bash
cp .env.example .env   # completa DB_USER y DB_PASS
go run ./cmd/api
```

Con Docker:

```bash
cd docker && docker compose up --build
```

## Contrato

`contrato/openapi.yaml` define cada endpoint y su respuesta. Hoy cubre:
`/salud`, `/data/report/`, `/api/actual`, `/api/historial`, `/api/stats`,
`/api/stats/comparative`.

## Pruebas

```bash
go test -race ./...
```
