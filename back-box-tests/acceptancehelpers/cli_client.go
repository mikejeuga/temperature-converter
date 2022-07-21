package acceptancehelpers

import (
	"fmt"
	"github.com/mikejeuga/temperature-converter/models"
	"os/exec"
	"strconv"
)

type TestCliCLient struct {
	fileName string
}

func NewTestCliCLient(fileName string) *TestCliCLient {
	return &TestCliCLient{
		fileName: fileName,
	}
}

func (c TestCliCLient) ConvertCtoF(temp models.Celsius) (models.Fahrenheit, error) {
	output, err := c.goRun()
	if err != nil {
		return 0, err
	}

	f := string(output)
	floatFahrenheit, err := strconv.ParseFloat(f, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting string to float, %v", err)
	}

	return models.Fahrenheit(floatFahrenheit), nil
}

func (c TestCliCLient) ConvertFtoC(temp models.Fahrenheit) (models.Celsius, error) {
	//TODO implement me
	panic("implement me")
}

func (c TestCliCLient) goRun() ([]byte, error) {
	cmd := exec.Command("go", "run", c.fileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error running the fileName built: %v", err)
	}
	return output, nil
}
