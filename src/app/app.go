package app

import (
	"github.com/mikejeuga/temperature-converter/src/internal/adapters/transport/server"
	"net/http"
)

type App struct {
	Server *http.Server
}

func NewApp() *App {
	server := server.NewServer()

	return &App{Server: server}
}



