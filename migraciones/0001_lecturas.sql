CREATE TABLE IF NOT EXISTS lecturas (
    id                BIGSERIAL PRIMARY KEY,
    timestamp         TIMESTAMPTZ NOT NULL DEFAULT now(),
    temp_exterior     DOUBLE PRECISION,
    temp_interior     DOUBLE PRECISION,
    humedad_exterior  DOUBLE PRECISION,
    humedad_interior  DOUBLE PRECISION,
    velocidad_viento  DOUBLE PRECISION,
    rafaga_viento     DOUBLE PRECISION,
    direccion_viento  INTEGER,
    lluvia_hora       DOUBLE PRECISION,
    lluvia_dia        DOUBLE PRECISION,
    lluvia_evento     DOUBLE PRECISION,
    presion_relativa  DOUBLE PRECISION,
    presion_absoluta  DOUBLE PRECISION,
    uv_index          INTEGER,
    radiacion_solar   DOUBLE PRECISION,
    punto_rocio       DOUBLE PRECISION,
    sensacion_termica DOUBLE PRECISION,
    passkey           VARCHAR(100)
);

CREATE INDEX IF NOT EXISTS idx_lecturas_timestamp ON lecturas ("timestamp" DESC);
