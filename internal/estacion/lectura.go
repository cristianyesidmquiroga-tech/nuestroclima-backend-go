package estacion

import (
	"net/url"
	"strings"

	"github.com/cristianyesidmquiroga-tech/nuestroclima-backend-go/internal/modelos"
)

const altitudRegion = 1800.0

func getField(datos url.Values, nombres ...string) string {
	for _, n := range nombres {
		if v := strings.TrimSpace(datos.Get(n)); v != "" {
			return v
		}
	}
	return ""
}

// ConstruirLectura arma una lectura calibrada a partir de los datos crudos que
// envia la estacion (mismas conversiones y compensacion que el sistema actual).
func ConstruirLectura(datos url.Values) *modelos.Lectura {
	radiacion := aFloat(datos.Get("solarradiation"))
	viento := MphAKmh(datos.Get("windspeedmph"))
	tempRaw := FahrenheitACelsius(datos.Get("tempf"))
	tempCalibrada := AplicarCompensacionTermica(tempRaw, radiacion, viento)

	var offset float64
	if tempRaw != nil && tempCalibrada != nil {
		offset = *tempRaw - *tempCalibrada
	}
	var sensacion *float64
	if s := FahrenheitACelsius(getField(datos, "feelslikef", "feelslike")); s != nil {
		v := redondear(*s-offset, 1)
		sensacion = &v
	}

	presionAbs := InhgAHpa(datos.Get("baromabsin"))
	presionRel := InhgAHpa(datos.Get("baromrelin"))
	if presionRel == nil {
		presionRel = AbsolutaARelativa(presionAbs, altitudRegion)
	}

	return &modelos.Lectura{
		TempExterior:     tempCalibrada,
		TempInterior:     FahrenheitACelsius(getField(datos, "tempinf", "indoortempf")),
		HumedadExterior:  CorregirHumedadPorTemp(aFloat(datos.Get("humidity")), tempRaw, tempCalibrada),
		HumedadInterior:  aFloat(getField(datos, "humidityin", "indoorhumidity")),
		VelocidadViento:  viento,
		RafagaViento:     MphAKmh(datos.Get("windgustmph")),
		DireccionViento:  AEntero(datos.Get("winddir")),
		LluviaHora:       PulgadasAMm(datos.Get("rainratein")),
		LluviaDia:        PulgadasAMm(datos.Get("dailyrainin")),
		LluviaEvento:     PulgadasAMm(datos.Get("eventrainin")),
		PresionRelativa:  presionRel,
		PresionAbsoluta:  presionAbs,
		UVIndex:          AEntero(datos.Get("uv")),
		RadiacionSolar:   radiacion,
		PuntoRocio:       FahrenheitACelsius(getField(datos, "dewptf", "dewpoint")),
		SensacionTermica: sensacion,
	}
}
