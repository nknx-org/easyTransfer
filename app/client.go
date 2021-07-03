package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	nkn "github.com/nknorg/nkn-sdk-go"
	"github.com/nknx-org/easyTransfer/model"
)

//The nkn client
var client *nkn.Client

//InitializeClient Initiates the surge client and instantiates connection with the NKN network
func InitializeClient() {
	var err error

	account := InitializeAccount()
	client, err = nkn.NewClient(account, "", &nkn.ClientConfig{
		ConnectRetries: 1000,
	})
	if err != nil {
		//pushError(err.Error(), "do you have an active internet connection?")
	}

	<-client.OnConnect.C
}

//StopClient Stops the surge client and cleans up
func StopClient() {
	client.Close()
}

// GetFileSize .
func GetFileSize(path string) int64 {

	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error on get filesize", err)
	}
	// get the size
	return fi.Size()
}

// HashFile generates hash for file given filepath
func ReadFile(path string) model.FileData {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	fileName := filepath.Base(path)

	fileData := model.FileData{
		Name: fileName,
		Data: b,
	}

	return fileData
}

func SendFile(path string, destination string) {
	//Read data from disk
	fileData := ReadFile(path)

	//Marshal to bytes
	bytes, err := json.Marshal(fileData)

	if err != nil {
		fmt.Println(err)
	}

	onMsg, err := client.SendBinary(nkn.NewStringArray(destination), bytes, nkn.GetDefaultMessageConfig())

	if err != nil {
		fmt.Println(err)
	}

	<-onMsg.C
	fmt.Println("File sent!")
}
