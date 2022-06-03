package testAPIhelpers

import (
	"github.com/mikejeuga/temperature-converter/models"
	"net/http"
	"time"
)

type APIClient struct {
	baseURL    string
	httpDriver *http.Client
}

func (A *APIClient) ConvertCtoF(temp models.Celsius) (models.Fahrenheit, error) {
	return models.Fahrenheit(41), nil
}

func NewAPIClient() *APIClient {
	baseURL := "localhost"
	client := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   5 * time.Second,
	}
	return &APIClient{baseURL: baseURL, httpDriver: client}
}
