package main

import (
	"fmt"
	"github.com/mikejeuga/temperature-converter/src/adapters/transport/server"
	"log"
)

func main() {
	fmt.Println("Check the temperature!")

	newServer := server.NewServer()

	err := newServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
