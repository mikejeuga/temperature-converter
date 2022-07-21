package cli

import (
	"fmt"
	"github.com/mikejeuga/temperature-converter/models"
	"os/exec"
	"strconv"
)

type APIClient struct {
	Cmd *exec.Cmd
}

func NewClient() *APIClient {
	return &APIClient{}
}

func (A *APIClient) ConvertCtoF(temp models.Celsius) (models.Fahrenheit, error) {
	A.Cmd = exec.Command("toF", fmt.Sprintf("%f", temp))
	output, err := A.Cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("error, you did not receive a correct output, %v", err)
	}

	tempFloat, coversionErr := strconv.ParseFloat(string(output), 64)
	if coversionErr != nil {
		return 0, fmt.Errorf("error converting to float, %v", coversionErr)
	}

	return models.Fahrenheit(tempFloat), nil
}

func (A *APIClient) ConvertFtoC(temp models.Fahrenheit) (models.Celsius, error) {
	//TODO implement me
	panic("implement me")
}
