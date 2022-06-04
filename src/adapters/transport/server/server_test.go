package server_test

import (
	"github.com/alecthomas/assert/v2"
	"github.com/mikejeuga/temperature-converter/src/adapters/transport/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerisHealthy(t *testing.T) {
	newServer := server.NewServer()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	newServer.Handler.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
	assert.Equal(t, resp.Body.String(), "Server is Healthy, the temperature seems perfect!")
}

func Test(t *testing.T) {
	t.Parallel()
	newServer := server.NewServer()
	for _, tc := range []struct {
		description         string
		request             *http.Request
		response            *httptest.ResponseRecorder
		expectedTemperature string
	}{
		{
			description:         "Convert Celsius to Fahrenheit",
			request:             httptest.NewRequest(http.MethodGet, "/to-fahrenheit/5", nil),
			response:            httptest.NewRecorder(),
			expectedTemperature: "41",
		},
		{
			description:         "Convert Fahrenheit to Celsius",
			request:             httptest.NewRequest(http.MethodGet, "/to-celsius/41", nil),
			response:            httptest.NewRecorder(),
			expectedTemperature: "5",
		},
	} {
		t.Run(tc.description, func(t *testing.T) {
			newServer.Handler.ServeHTTP(tc.response, tc.request)
			assert.Equal(t, tc.response.Code, http.StatusOK)
			assert.Equal(t, tc.response.Body.String(), tc.expectedTemperature)
		})
	}
}
