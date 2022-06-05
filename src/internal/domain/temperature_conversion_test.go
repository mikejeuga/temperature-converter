//+go:build unit

package domain_test

import (
	"github.com/mikejeuga/temperature-converter/specifications"
	"github.com/mikejeuga/temperature-converter/src/internal/domain"
	"testing"
)

func TestConvertCtoF(t *testing.T) {
	converter := domain.NewConverter(domain.ConvertCtoF, domain.ConvertFtoC)
	specifications.ConvertCtoFCriteria(t, converter)
	specifications.ConvertFtoCCriteria(t, converter)
}
