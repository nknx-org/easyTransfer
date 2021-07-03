package main

import (
	"fmt"
	"os"

	"github.com/nknx-org/easyTransfer/app"
)

func main() {

	//ARGS STUFF
	args := os.Args[1:]

	filePath := args[0]
	fileDestination := args[1]

	fmt.Println(filePath)
	fmt.Println(fileDestination)

	app.InitializeClient()

	//Work
	app.SendFile(filePath, fileDestination)

	app.StopClient()
}
