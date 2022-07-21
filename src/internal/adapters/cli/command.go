package cli

import (
	"fmt"
	"github.com/mikejeuga/temperature-converter/models"
	"github.com/mikejeuga/temperature-converter/src/internal/domain"
	"io"
	"strconv"
)

type CLI struct {
	tempConverter domain.Converter
}

func (c CLI) ConvertCtoF(w io.Writer, r io.Reader) (models.Fahrenheit, error) {
	readAll, err := io.ReadAll(r)
	if err != nil {
		fmt.Errorf("error reading data from r, %v", err)
	}

	input := string(readAll)

	float, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Errorf("converting string to float, %v", err)
	}

	celsius := models.Celsius(float)
	fahrenheit, err := c.tempConverter.ConvertCtoF(celsius)
	if err != nil {
		fmt.Errorf("error converting Celsius to Fahrenheit, %v", err)
	}

	return fahrenheit, nil
}

func (c CLI) ConvertFtoC(temp models.Fahrenheit) (models.Celsius, error) {
	//TODO implement me
	panic("implement me")
}

func NewCLI(tempConverter domain.Converter) *CLI {
	return &CLI{tempConverter: tempConverter}
}
