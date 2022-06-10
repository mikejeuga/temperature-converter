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

type TemperatureConverterSpec struct {
	converter TemperatureConverter
}

func NewTemperatureConverterSpec(converter TemperatureConverter) *TemperatureConverterSpec {
	return &TemperatureConverterSpec{converter: converter}
}

func (th *TemperatureConverterSpec) ConvertCelsiusToFahrenheit(t *testing.T) {
	t.Helper()
	is := is.New(t)
	t.Run("Given a converter receives the temperature in Celsius:", func(t *testing.T) {
		t.Run("It gives back the temperature in Fahrenheit.", func(t *testing.T) {
			celsius := models.Celsius(5)
			expectedFahrenheit := models.Fahrenheit(41)

			fahrenheit, err := th.converter.ConvertCtoF(celsius)
			is.NoErr(err)

			is.Equal(fahrenheit, expectedFahrenheit)
		})
	})
}

func (th *TemperatureConverterSpec) ConvertFahrenheitToCelsius(t *testing.T) {
	t.Helper()
	is := is.New(t)
	t.Run("Given a converter receives the temperature in Fahrenheit:", func(t *testing.T) {
		t.Run("It gives back the temperature in Celsius.", func(t *testing.T) {
			fahrenheit := models.Fahrenheit(41)
			expectedCelsius := models.Celsius(5)

			celsius, err := th.converter.ConvertFtoC(fahrenheit)
			is.NoErr(err)

			is.Equal(celsius, expectedCelsius)
		})
	})
}
