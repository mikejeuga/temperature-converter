package main

import (
	"fmt"
	"github.com/mikejeuga/temperature-converter/src/app/web"
	"log"
)

func main() {
	fmt.Println("Check the temperature!")

	App := web.NewApp()

	err := App.Server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
