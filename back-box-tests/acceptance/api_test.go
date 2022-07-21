//+go:build acceptance

package acceptance

import (
	"github.com/mikejeuga/temperature-converter/back-box-tests/acceptancehelpers/cli"
	"github.com/mikejeuga/temperature-converter/back-box-tests/acceptancehelpers/web"
	"github.com/mikejeuga/temperature-converter/specifications"
	"testing"
)

func TestTemperatureConversionAPI(t *testing.T) {

	t.Run("Acceptance test with the HTTP driver", func(t *testing.T) {

		apiClient := web.NewAPIClient()
		spec := specifications.NewTemperatureConverterSpec(apiClient)

		spec.ConvertCelsiusToFahrenheit(t)

		spec.ConvertFahrenheitToCelsius(t)
	})

	t.Run("Acceptance test with the CLI driver", func(t *testing.T) {
		cliClient := cli.NewClient()

		spec := specifications.NewTemperatureConverterSpec(cliClient)
		spec.ConvertCelsiusToFahrenheit(t)

	})

}
