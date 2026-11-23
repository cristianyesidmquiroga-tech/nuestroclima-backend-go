package modelos

type IntegridadDia struct {
	Fecha    string `json:"fecha"`
	FechaNum uint32 `json:"fecha_num"`
	Cantidad int    `json:"cantidad"`
	Hash     string `json:"hash"`
}
