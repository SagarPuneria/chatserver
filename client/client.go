package main

import (
	"fmt"
	"io"
	"log"
	"net"

	ut "chatserver/util"
)

func main() {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred at ", ut.RecoverExceptionDetails(ut.FunctionName()), " and recovered in main function, Error Info: ", errD)
		}
	}()
	address := ut.GetAddress()
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("Error in net.Dial and Error Info:", err)
	}
	fmt.Println("Connected to Server.")
	defer conn.Close()
	clientCV := ut.Random(6, 10)
	clientPV := ut.Random(1, 5)
	fmt.Println("Client current version", clientCV)
	fmt.Println("Client previous version", clientPV)

	message := ut.Message{CurrentVersion: clientCV}
	err = ut.SendCurrentVersion(conn, message)
	if err != nil {
		log.Fatal("Error in SendCurrentVersion, Error Info:", err)
	}
	message, err = ut.GetCurrentVersion(conn)
	if err != nil && err != io.EOF {
		log.Fatal("Error in GetCurrentVersion and Error Info:", err)
	} else if err == nil {
		clientPV = clientCV
		clientCV = message.CurrentVersion
		fmt.Println("Updating current version received from server")
		fmt.Println("Client current version", clientCV)
		fmt.Println("Client previous version", clientPV)
	} else {
		fmt.Println("No update require")
	}
}
