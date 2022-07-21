package domain

import (
	"github.com/mikejeuga/temperature-converter/models"
)

type ConverterCtoF func(temp models.Celsius) (models.Fahrenheit, error)
type ConverterFtoC func(temp models.Fahrenheit) (models.Celsius, error)

type Converter struct {
	CelsiusConverter    ConverterCtoF
	FahrenheitConverter ConverterFtoC
}

func NewConverter(celsiusConverter ConverterCtoF, fahrenheitConverter ConverterFtoC) Converter {
	return Converter{CelsiusConverter: celsiusConverter, FahrenheitConverter: fahrenheitConverter}
}

func (c Converter) ConvertCtoF(temp models.Celsius) (models.Fahrenheit, error) {
	return c.CelsiusConverter(temp)
}

func (c Converter) ConvertFtoC(temp models.Fahrenheit) (models.Celsius, error) {
	return c.FahrenheitConverter(temp)
}

func ConvertCtoF(temp models.Celsius) (models.Fahrenheit, error) {
	result := (float64(temp) * 1.8) + 32
	return models.Fahrenheit(result), nil
}

func ConvertFtoC(temp models.Fahrenheit) (models.Celsius, error) {
	result := (float64(temp) - 32) / 1.8
	return models.Celsius(result), nil
}
