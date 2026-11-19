package estacion

import (
	"math"
	"strconv"
	"strings"
)

func aFloat(s string) *float64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil
	}
	return &v
}

func redondear(v float64, dec int) float64 {
	f := math.Pow(10, float64(dec))
	return math.Round(v*f) / f
}

func FahrenheitACelsius(raw string) *float64 {
	f := aFloat(raw)
	if f == nil {
		return nil
	}
	v := redondear((*f-32)*5/9, 1)
	return &v
}

func MphAKmh(raw string) *float64 {
	f := aFloat(raw)
	if f == nil {
		return nil
	}
	v := redondear(*f*1.60934, 1)
	return &v
}

func PulgadasAMm(raw string) *float64 {
	f := aFloat(raw)
	if f == nil {
		return nil
	}
	v := redondear(*f*25.4, 2)
	return &v
}

func InhgAHpa(raw string) *float64 {
	f := aFloat(raw)
	if f == nil {
		return nil
	}
	v := redondear(*f*33.8639, 1)
	return &v
}

func AEntero(raw string) *int {
	f := aFloat(raw)
	if f == nil {
		return nil
	}
	v := int(math.Round(*f))
	return &v
}

// AplicarCompensacionTermica corrige el autocalentamiento del sensor por sol
// directo sobre su carcasa; el viento enfria por conveccion y reduce el error.
func AplicarCompensacionTermica(tempC, radiacion, vientoKmh *float64) *float64 {
	if tempC == nil || radiacion == nil || vientoKmh == nil {
		return tempC
	}
	temp, rad, viento := *tempC, *radiacion, *vientoKmh

	var base float64
	switch {
	case rad > 1200:
		base = 8.0
	case rad > 900:
		base = 6.0
	case rad > 600:
		base = 4.0
	case rad > 300:
		base = 2.0
	default:
		base = 0.0
	}
	if base == 0.0 {
		v := redondear(temp, 1)
		return &v
	}

	var factor float64
	switch {
	case viento >= 10:
		factor = 0.4
	case viento >= 5:
		factor = 0.6
	case viento >= 3:
		factor = 0.8
	default:
		factor = 1.0
	}

	v := redondear(temp-base*factor, 1)
	return &v
}

func presionVaporSaturado(tempC float64) float64 {
	return 6.1094 * math.Exp(17.625*tempC/(tempC+243.04))
}

// CorregirHumedadPorTemp recalcula la humedad relativa que corresponde al
// vapor de agua medido, evaluado a la temperatura ya corregida.
func CorregirHumedadPorTemp(humedadRaw, tempRaw, tempCorregida *float64) *float64 {
	if humedadRaw == nil || tempRaw == nil || tempCorregida == nil {
		return humedadRaw
	}
	h, tRaw, tCorr := *humedadRaw, *tempRaw, *tempCorregida
	if tRaw == tCorr {
		v := redondear(h, 1)
		return &v
	}
	vapor := (h / 100.0) * presionVaporSaturado(tRaw)
	hCorr := (vapor / presionVaporSaturado(tCorr)) * 100.0
	v := redondear(math.Max(0.0, math.Min(100.0, hCorr)), 1)
	return &v
}

func AbsolutaARelativa(hpa *float64, altitudM float64) *float64 {
	if hpa == nil {
		return nil
	}
	v := redondear(*hpa/math.Pow(1-(0.0065*altitudM/288.15), 5.255), 1)
	return &v
}
