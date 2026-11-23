package integridad

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/modelos"
)

// Colombia no tiene horario de verano; se usa un offset fijo para no depender
// de la base de datos de zonas horarias dentro del contenedor.
var zonaColombia = time.FixedZone("-05", -5*3600)

// RangoDia interpreta "AAAAMMDD" o "AAAA-MM-DD" como un dia en hora de Colombia
// y devuelve el inicio (inclusive) y el fin (exclusivo) de ese dia.
func RangoDia(fecha string) (time.Time, time.Time, error) {
	f := strings.ReplaceAll(fecha, "-", "")
	t, err := time.ParseInLocation("20060102", f, zonaColombia)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return t, t.AddDate(0, 0, 1), nil
}

// NumeroFecha devuelve la fecha como entero AAAAMMDD, el formato que espera el
// contrato para anclar el hash.
func NumeroFecha(inicio time.Time) uint32 {
	t := inicio.In(zonaColombia)
	return uint32(t.Year()*10000 + int(t.Month())*100 + t.Day())
}

// Hash arma una huella SHA-256 determinista de las lecturas: una linea por
// lectura con sus campos en orden fijo, ordenadas por contenido para no
// depender del orden de entrada. Sin lecturas, devuelve cadena vacia.
func Hash(lecturas []modelos.Lectura) string {
	if len(lecturas) == 0 {
		return ""
	}
	lineas := make([]string, len(lecturas))
	for i, l := range lecturas {
		lineas[i] = strings.Join([]string{
			strconv.FormatInt(l.Timestamp.UTC().Unix(), 10),
			f(l.TempExterior), f(l.TempInterior), f(l.HumedadExterior),
			f(l.HumedadInterior), f(l.VelocidadViento), f(l.RafagaViento),
			e(l.DireccionViento), f(l.LluviaHora), f(l.LluviaDia), f(l.LluviaEvento),
			f(l.PresionRelativa), f(l.PresionAbsoluta), e(l.UVIndex),
			f(l.RadiacionSolar), f(l.PuntoRocio), f(l.SensacionTermica),
		}, "|")
	}
	sort.Strings(lineas)
	suma := sha256.Sum256([]byte(strings.Join(lineas, "\n")))
	return hex.EncodeToString(suma[:])
}

func f(v *float64) string {
	if v == nil {
		return ""
	}
	return strconv.FormatFloat(*v, 'f', -1, 64)
}

func e(v *int) string {
	if v == nil {
		return ""
	}
	return strconv.Itoa(*v)
}
