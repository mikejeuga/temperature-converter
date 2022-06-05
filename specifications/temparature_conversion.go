package specifications

import (
	"github.com/matryer/is"
	"github.com/mikejeuga/temperature-converter/models"
	"testing"
)

type TemperatureConverter interface {
	ConvertCtoF(temp models.Celsius) (models.Fahrenheit, error)
	ConvertFtoC(temp models.Fahrenheit) (models.Celsius, error)
}

func ConvertCelsiusToFahrenheit(t *testing.T, converter TemperatureConverter) {
	t.Helper()
	is := is.New(t)
	t.Run("Given a converter receives the temperature in Celsius:", func(t *testing.T) {
		t.Run("It gives back the temperature in Fahrenheit.", func(t *testing.T) {
			celsius := models.Celsius(5)
			expectedFahrenheit := models.Fahrenheit(41)

			fahrenheit, err := converter.ConvertCtoF(celsius)
			is.NoErr(err)

			is.Equal(fahrenheit, expectedFahrenheit)
		})
	})
}

func ConvertFahrenheitToCelsius(t *testing.T, converter TemperatureConverter) {
	t.Helper()
	is := is.New(t)
	t.Run("Given a converter receives the temperature in Fahrenheit:", func(t *testing.T) {
		t.Run("It gives back the temperature in Celsius.", func(t *testing.T) {
			fahrenheit := models.Fahrenheit(41)
			expectedCelsius := models.Celsius(5)

			celsius, err := converter.ConvertFtoC(fahrenheit)
			is.NoErr(err)

			is.Equal(celsius, expectedCelsius)
		})
	})
}
