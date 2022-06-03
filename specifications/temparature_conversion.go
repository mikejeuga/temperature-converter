package specifications

import (
	"github.com/alecthomas/assert/v2"
	"github.com/mikejeuga/temperature-converter/models"
	"testing"
)

type TemperatureConverter interface {
	ConvertCtoF(temp models.Celsius) (models.Fahrenheit, error)
}

func ConvertCtoFCriteria(t *testing.T, converter TemperatureConverter) {
	t.Run("Given a converter receives the temperature in Celsius:", func(t *testing.T) {
		t.Run("It gives back the temperature in Fahrenheit.", func(t *testing.T) {
			celsius := models.Celsius(5)
			expectedFahrenheit := models.Fahrenheit(41)

			fahrenheit, err := converter.ConvertCtoF(celsius)
			assert.NoError(t, err)

			assert.Equal(t, fahrenheit, expectedFahrenheit)
		})
	})
}
