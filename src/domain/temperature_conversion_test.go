package domain_test

import (
	"github.com/mikejeuga/temperature-converter/specifications"
	"github.com/mikejeuga/temperature-converter/src/domain"
	"testing"
)

func TestConvertCtoF(t *testing.T) {
	converter := domain.Converter(domain.ConvertCtoF)

	specifications.ConvertCtoFCriteria(t, converter)
}
