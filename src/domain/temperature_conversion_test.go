package domain_test

import (
	"github.com/alecthomas/assert/v2"
	"github.com/mikejeuga/temperature-converter/models"
	"github.com/mikejeuga/temperature-converter/src/domain"
	"testing"
)

func TestConvertCtoF(t *testing.T) {
	celsius := models.Celsius(5)
	expected := models.Fahrenheit(41)

	fahrenheit, err := domain.ConvertCtoF(celsius)
	assert.NoError(t, err)

	assert.Equal(t, fahrenheit, expected)
}
