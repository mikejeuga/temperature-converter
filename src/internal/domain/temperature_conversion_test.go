//+go:build unit

package domain_test

import (
	"github.com/mikejeuga/temperature-converter/specifications"
	"github.com/mikejeuga/temperature-converter/src/internal/domain"
	"testing"
)

func TestConvertCtoF(t *testing.T) {
	converter := domain.NewConverter(domain.ConvertCtoF, domain.ConvertFtoC)
	spec := specifications.NewTemperatureConverterSpec(converter)

	spec.ConvertCelsiusToFahrenheit(t)
	spec.ConvertFahrenheitToCelsius(t)
}
