package app

import (
	"github.com/mikejeuga/temperature-converter/src/internal/adapters/command"
	"github.com/mikejeuga/temperature-converter/src/internal/adapters/transport/server"
	"net/http"
)

type AppWebb struct {
	Server *http.Server
}

type AppCli struct {
	Cli *command.CLI
}

func NewAppCli() *AppCli {
	cli := command.NewCLI()
	return &AppCli{Cli: cli}
}

func NewApp() *AppWebb {
	newServer := server.NewServer()
	return &AppWebb{
		Server: newServer,
	}
}
