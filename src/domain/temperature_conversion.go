package domain

import "github.com/mikejeuga/temperature-converter/models"

type Converter func(temp models.Celsius) (models.Fahrenheit, error)

func (c Converter) ConvertCtoF(temp models.Celsius) (models.Fahrenheit, error) {
	return c(temp)
}

func ConvertCtoF(temp models.Celsius) (models.Fahrenheit, error) {
	result := (float64(temp) * 1.8) + 32
	return models.Fahrenheit(result), nil
}
