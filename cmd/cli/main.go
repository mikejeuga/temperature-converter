package main

import (
	"fmt"
	"os"
)

func main() {
	//App := app.NewAppCli()
	args := os.Args[1:]

	//App.Cli.ConvertCtoF(os.Stdout, os.Stdin)
	fmt.Println(args)
}
