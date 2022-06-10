//+go:build acceptance

package acceptance

import (
	"github.com/mikejeuga/temperature-converter/back-box-tests/acceptancehelpers"
	"github.com/mikejeuga/temperature-converter/specifications"
	"testing"
)

func TestTemperatureConversionAPI(t *testing.T) {

	apiClient := acceptancehelpers.NewAPIClient()
	spec := specifications.NewTemperatureConverterSpec(apiClient)

	spec.ConvertCelsiusToFahrenheit(t)

	spec.ConvertFahrenheitToCelsius(t)

}
