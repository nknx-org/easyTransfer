package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/nknx-org/easyTransfer/app"
)

const DEFAULT_TIMEOUT = 15

func main() {
	//Prevalidate args
	if len(os.Args) < 3 {
		fmt.Println("Two parameters expected (path, destAddr)")
		os.Exit(1)
	}

	//Take only given args
	args := os.Args[1:]
	filePathArg := args[0]
	fileDestination := args[1]

	filePathArr := strings.Split(filePathArg, ";")

	//Take optional timeout param
	var timeOut = DEFAULT_TIMEOUT
	if len(args) > 2 {
		var atoiErr error
		timeOut, atoiErr = strconv.Atoi(args[2])
		if atoiErr != nil {
			fmt.Println("Timeout param was not of type integer")
			os.Exit(1)
		}
	}

	//Timeout
	go timeout(timeOut)

	err := app.InitializeClient()
	if err != nil {
		os.Exit(1)
	}

	//Work
	for i := 0; i < len(filePathArr); i++ {
		err = app.SendFile(filePathArr[i], fileDestination)
		if err != nil {
			os.Exit(1)
		}
	}

	app.StopClient()

	os.Exit(0)
}

func timeout(timeout int) {
	time.Sleep(time.Duration(timeout) * time.Second)
	fmt.Println("Application timeout after " + strconv.Itoa(timeout) + " seconds")
	os.Exit(1)
}
