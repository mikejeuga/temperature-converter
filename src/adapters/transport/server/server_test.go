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

func TestConvertCtoF(t *testing.T) {
	newServer := server.NewServer()
	req := httptest.NewRequest(http.MethodGet, "/tofahrenheit/5", nil)
	resp := httptest.NewRecorder()

	newServer.Handler.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)

}
