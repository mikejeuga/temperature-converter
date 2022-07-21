package cli_test

import (
	"bytes"
	"github.com/matryer/is"
	"github.com/mikejeuga/temperature-converter/src/internal/adapters/cli"
	"github.com/mikejeuga/temperature-converter/src/internal/domain"
	"testing"
)

func TestCLI(t *testing.T) {
	is := is.New(t)
	converter := domain.NewConverter(domain.DefaultC2F, domain.DefaultF2C)
	newCLI := cli.NewCLI(converter)

	buffer := bytes.Buffer{}
	value := []byte(`5`)

	data := bytes.NewReader(value)
	fahrenheit, err := newCLI.ConvertCtoF(&buffer, data)
	is.NoErr(err)

	want := `41`
	is.Equal(fahrenheit.String(), want)

}
