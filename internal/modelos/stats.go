package modelos

type StatsDia struct {
	TempMax     *float64 `json:"temp_max"`
	TempMin     *float64 `json:"temp_min"`
	HumMax      *float64 `json:"hum_max"`
	HumMin      *float64 `json:"hum_min"`
	VientoMax   *float64 `json:"viento_max"`
	RafagaMax   *float64 `json:"rafaga_max"`
	LluviaTotal *float64 `json:"lluvia_total"`
	UVMax       *int     `json:"uv_max"`
}

type StatsMensual struct {
	Mes                   string   `json:"mes"`
	TempMax               *float64 `json:"temp_max"`
	TempMin               *float64 `json:"temp_min"`
	LluviaMaxDiaria       *float64 `json:"lluvia_max_diaria"`
	EstimacionLluviaTotal *float64 `json:"estimacion_lluvia_total"`
}

type StatsSemanal struct {
	Semana          string   `json:"semana"`
	TempMax         *float64 `json:"temp_max"`
	TempMin         *float64 `json:"temp_min"`
	LluviaMaxDiaria *float64 `json:"lluvia_max_diaria"`
}

type StatsComparativas struct {
	Mensual []StatsMensual `json:"mensual"`
	Semanal []StatsSemanal `json:"semanal"`
}
