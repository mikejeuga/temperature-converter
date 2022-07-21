package command

import (
	"flag"
	"fmt"
	"github.com/mikejeuga/temperature-converter/models"
	"github.com/mikejeuga/temperature-converter/src/internal/domain"
	"io"
	"os"
	"strconv"
)

type CLI struct {
	tempConverter domain.Converter
}

func NewCLI() *CLI {
	converter := domain.NewConverter(domain.ConvertCtoF, domain.ConvertFtoC)
	return &CLI{tempConverter: converter}
}

func (c CLI) Convert() {
	s := flag.String("ConverterCtoF", "toF", "Flag for converting Celsius to Fahrenheit")
	flag.Parse()

	c.ConvertCtoF(os.Stdout, os.Stdin)

	fmt.Println(s)
}

func (c CLI) ConvertCtoF(out io.Writer, in io.Reader) error {

	byteTemp, err := io.ReadAll(in)
	if err != nil {
		return fmt.Errorf("there is no in put given, %v", err)
	}

	fmt.Println(byteTemp)

	tempFloat, err := strconv.ParseFloat(string(byteTemp), 64)
	if err != nil {
		return fmt.Errorf("error converting to float, %v", err)
	}

	fahrenheit, err := c.tempConverter.ConvertCtoF(models.Celsius(int(tempFloat)))
	if err != nil {
		return fmt.Errorf("error converting to Fahrenheit, %v", err)
	}

	fmt.Fprintf(out, fmt.Sprintf("%f", fahrenheit))

	return nil
}

func (c CLI) ConvertFtoC(out io.Writer, in io.Reader) {

}
