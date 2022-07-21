package main

import (
	"fmt"
	"github.com/mikejeuga/temperature-converter/src/app"
	"log"
)

func main() {
	fmt.Println("Check the temperature!")

	App := app.NewApp()

	err := App.Server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
