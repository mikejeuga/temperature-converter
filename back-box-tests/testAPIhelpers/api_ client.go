package testAPIhelpers

import (
	"fmt"
	"github.com/mikejeuga/temperature-converter/models"
	"io"
	"net/http"
	"strconv"
	"time"
)

type APIClient struct {
	baseURL    string
	httpDriver *http.Client
}

func (c *APIClient) ConvertCtoF(temp models.Celsius) (models.Fahrenheit, error) {
	url := c.baseURL + "/tofahrenheit/" + fmt.Sprintf("%v", temp)

	resp, err := c.httpDriver.Get(url)
	if err != nil {
		return 0, fmt.Errorf("problem reaching out %s, %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected response code %d but expected %d from %s", resp.StatusCode, http.StatusOK, url)
	}

	fahrenheit, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	temperatureResponse, err := strconv.ParseFloat(string(fahrenheit), 64)
	if err != nil {
		return 0, err
	}

	return models.Fahrenheit(temperatureResponse), nil
}

func NewAPIClient() *APIClient {
	baseURL := "https://localhost:8069"
	client := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   5 * time.Second,
	}
	return &APIClient{baseURL: baseURL, httpDriver: client}
}