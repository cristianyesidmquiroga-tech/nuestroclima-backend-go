# AGENTS.md

Backend en Go de **Nuestro Clima**: la API del dashboard meteorologico de la
finca. Es una de varias implementaciones del mismo contrato
(`contrato/openapi.yaml`), pensada para comparar tecnologias de backend. El
frontend consume ese contrato sin importar que backend este detras.

## Stack

- Go 1.27, router `go-chi/chi`, `jackc/pgx` sobre PostgreSQL.
- Sin ORM: SQL directo contra el contrato.

## Estructura

- `cmd/api/` arranque del servidor.
- `internal/config/` carga de variables de entorno.
- `internal/db/` conexion y consultas (almacen).
- `internal/api/` router, manejadores y respuestas.
- `migraciones/` SQL de esquema.
- `contrato/openapi.yaml` fuente de verdad de la API.
- `docker/` Dockerfile y compose.

## Reglas

- Una contribucion por funcionalidad; commits en español `tipo: descripcion`.
- Todo cambio entra por Pull Request; se fusiona solo con `pruebas` y `codeql`
  en verde (squash y borrar rama). Nada de push directo a `main`.
- Nada de `git push --force` ni reescribir historial publicado; se corrige con
  `git revert` o un commit nuevo.
- Toda logica nueva lleva sus pruebas en el mismo PR.
- Dependencias siempre en su ultima version de parche con soporte; `govulncheck`
  debe pasar.
- El `.env` real nunca se sube; solo `.env.example`.
