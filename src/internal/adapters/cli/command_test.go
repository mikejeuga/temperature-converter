package cli_test

import (
	"bytes"
	"github.com/matryer/is"
	"github.com/mikejeuga/temperature-converter/src/internal/adapters/cli"
	"testing"
)

func TestCLI(t *testing.T) {
	is := is.New(t)
	newCLI := cli.NewCLI()

	buffer := &bytes.Buffer{}
	value := []byte(`5`)

	data := bytes.NewReader(value)
	newCLI.ConvertCtoF(buffer, data)

	want := `41`
	is.Equal(buffer.String(), want)

}
