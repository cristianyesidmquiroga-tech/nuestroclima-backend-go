package modelos

import "time"

type Lectura struct {
	ID               int64     `json:"id"`
	Timestamp        time.Time `json:"timestamp"`
	TempExterior     *float64  `json:"temp_exterior"`
	TempInterior     *float64  `json:"temp_interior"`
	HumedadExterior  *float64  `json:"humedad_exterior"`
	HumedadInterior  *float64  `json:"humedad_interior"`
	VelocidadViento  *float64  `json:"velocidad_viento"`
	RafagaViento     *float64  `json:"rafaga_viento"`
	DireccionViento  *int      `json:"direccion_viento"`
	LluviaHora       *float64  `json:"lluvia_hora"`
	LluviaDia        *float64  `json:"lluvia_dia"`
	LluviaEvento     *float64  `json:"lluvia_evento"`
	PresionRelativa  *float64  `json:"presion_relativa"`
	PresionAbsoluta  *float64  `json:"presion_absoluta"`
	UVIndex          *int      `json:"uv_index"`
	RadiacionSolar   *float64  `json:"radiacion_solar"`
	PuntoRocio       *float64  `json:"punto_rocio"`
	SensacionTermica *float64  `json:"sensacion_termica"`
}
