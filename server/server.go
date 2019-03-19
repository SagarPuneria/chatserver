package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	ut "chatserver/util"
)

func main() {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred at ", ut.RecoverExceptionDetails(ut.FunctionName()), " and recovered in main function, Error Info: ", errD)
		}
	}()
	address := ut.GetAddress()
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Error in net.Listen and Error Info:", err)
	}
	defer ln.Close()
	fmt.Println("Server Started...")
	serverCV := ut.Random(6, 10)
	serverPV := ut.Random(1, 5)
	fmt.Println("Server current version", serverCV)
	fmt.Println("Server previous version", serverPV)
	scanner := bufio.NewScanner(os.Stdin)
	limit := 10000
	fmt.Println("Enter number of clients you want to connect(default is 10000)")
	scanner.Scan()
	if strings.TrimSpace(scanner.Text()) != "" {
		if val, err := strconv.Atoi(strings.TrimSpace(scanner.Text())); err == nil {
			limit = val
		}
	}
	var wg sync.WaitGroup
	for i := 0; i < limit; i++ {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error in ln.Accept and Error Info:", err)
			conn.Close()
			continue
		}
		fmt.Println("Connection Accepted")
		wg.Add(1)
		go handleConnection(serverCV, conn, wg)
	}
	wg.Wait()
}

func handleConnection(serverCV string, con net.Conn, wg sync.WaitGroup) {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred at ", ut.RecoverExceptionDetails(ut.FunctionName()), " and recovered in handleConnection function, Error Info: ", errD)
		}
	}()
	defer wg.Done()
	defer con.Close()
	message, err := ut.GetCurrentVersion(con)
	if err != nil {
		return
	}
	if message.CurrentVersion != serverCV {
		message.CurrentVersion = serverCV
		err = ut.SendCurrentVersion(con, message)
		if err != nil {
			return
		}
	}
}
