package acceptance

import (
	"github.com/mikejeuga/temperature-converter/back-box-tests/acceptancehelpers"
	"github.com/mikejeuga/temperature-converter/specifications"
	"testing"
)

func TestTemperatureConversionAPI(t *testing.T) {
	t.Skip("trying the actions")
	apiClient := acceptancehelpers.NewAPIClient()

	t.Run("Convert Celsius to Fahrenheit", func(t *testing.T) {
		specifications.ConvertCtoFCriteria(t, apiClient)
	})

	t.Run("Convert Fahrenheit to Celsius", func(t *testing.T) {
		specifications.ConvertFtoCCriteria(t, apiClient)
	})

}
