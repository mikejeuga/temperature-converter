//+go:build acceptance

package acceptance

import (
	"github.com/mikejeuga/temperature-converter/back-box-tests/acceptancehelpers/cli"
	"github.com/mikejeuga/temperature-converter/back-box-tests/acceptancehelpers/web"
	"github.com/mikejeuga/temperature-converter/specifications"
	"log"
	"os"
	"testing"
)

func TestTemperatureConversionAPI(t *testing.T) {

	t.Run("Acceptance test with the HTTP driver", func(t *testing.T) {

		apiClient := web.NewAPIClient()
		spec := specifications.NewTemperatureConverterSpec(apiClient)

		spec.ConvertCelsiusToFahrenheit(t)
		spec.ConvertFahrenheitToCelsius(t)
		spec.ConvertCelsiusToFahrenheit(t)

		spec.ConvertFahrenheitToCelsius(t)
	})

}

func TestMain(m *testing.M) {

	err := os.Chdir(os.Getenv("TEST_DIR"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestCLITemperatureConverterAPI(t *testing.T) {

	fileName := "main.go"
	testClientCLI := cli.NewTestCliCLient(fileName)

	spec := specifications.NewTemperatureConverterSpec(testClientCLI)

	spec.ConvertCelsiusToFahrenheit(t)
}
