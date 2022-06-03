package acceptance

import (
	"github.com/mikejeuga/temperature-converter/back-box-tests/testAPIhelpers"
	"github.com/mikejeuga/temperature-converter/specifications"
	"testing"
)

func TestTemperatureConversionAPI(t *testing.T) {
	apiClient := testAPIhelpers.NewAPIClient()

	t.Run("Convert Celsius to Fahrenheit", func(t *testing.T) {
		specifications.ConvertCtoFCriteria(t, apiClient)
	})

}
