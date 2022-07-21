package web

import (
	"github.com/mikejeuga/temperature-converter/src/internal/adapters/transport/server"
	"log"
	"net/http"
)

type App struct {
	Server *http.Server
}

func NewApp() *App {
	newServer := server.NewServer()

	err := newServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	return &App{Server: newServer}
}
