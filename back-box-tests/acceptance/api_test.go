//+go:build acceptance

package acceptance

import (
	"github.com/mikejeuga/temperature-converter/back-box-tests/acceptancehelpers"
	"github.com/mikejeuga/temperature-converter/specifications"
	"testing"
)

func TestTemperatureConversionAPI(t *testing.T) {

	apiClient := acceptancehelpers.NewAPIClient()

	t.Run("Convert Celsius to Fahrenheit", func(t *testing.T) {
		specifications.ConvertCelsiusToFahrenheit(t, apiClient)
	})

	t.Run("Convert Fahrenheit to Celsius", func(t *testing.T) {
		specifications.ConvertFahrenheitToCelsius(t, apiClient)
	})

}
