package main

import (
	"fmt"
	"os"

	"github.com/nknx-org/easyTransfer/app"
)

func main() {
	//Prevalidate args
	if len(os.Args) < 3 {
		fmt.Println("Two parameters expected (path, destAddr)")
		os.Exit(1)
	}

	//Take only given args
	args := os.Args[1:]
	filePath := args[0]
	fileDestination := args[1]

	err := app.InitializeClient()
	if err != nil {
		os.Exit(1)
	}

	//Work
	err = app.SendFile(filePath, fileDestination)
	if err != nil {
		os.Exit(1)
	}

	app.StopClient()

	os.Exit(0)
}
