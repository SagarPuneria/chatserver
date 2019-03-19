package util

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	ip "chatserver/ipvalidation"
)

// Message struct stores CurrentVersion
type Message struct {
	CurrentVersion string
}

// GetCurrentVersion retrieve current version
func GetCurrentVersion(cn net.Conn) (message Message, err error) {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred at ", RecoverExceptionDetails(FunctionName()), " and recovered in GetCurrentVersion function, Error Info: ", errD)
		}
	}()
	// retriving data from buffer and storing in a decoder object
	decoderObject := gob.NewDecoder(cn)

	// decodes buffer and unmarshals it into Message struct
	err = decoderObject.Decode(&message)
	if err != nil {
		fmt.Println("decode error:", err)
	}
	return
}

// SendCurrentVersion send current version
func SendCurrentVersion(cn net.Conn, message Message) (err error) {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred at ", RecoverExceptionDetails(FunctionName()), " and recovered in SendCurrentVersion function, Error Info: ", errD)
		}
	}()
	// create a encoder object
	encoderObject := gob.NewEncoder(cn)

	// encode buffer and marshal it into a encoderObject
	err = encoderObject.Encode(message)
	if err != nil {
		fmt.Println("Encode error:", err)
	}
	return
}

// Random Function genrate a random number
func Random(min, max int) string {
	rand.Seed(time.Now().Unix())
	return strconv.Itoa(rand.Intn(max-min) + min)
}

// readInputValue which Scans and read a line from Stdin(Console) - return Console input value.
func readInputValue(scanner *bufio.Scanner) string {
	scanner.Scan()
	inputValue := scanner.Text()
	if strings.TrimSpace(inputValue) == "" {
		inputValue = "8080"
	}
	return inputValue
}

// GetAddress function return address
func GetAddress() string {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred at ", RecoverExceptionDetails(FunctionName()), " and recovered in GetAddress function, Error Info: ", errD)
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Server IP address you want to connect:")
	ipString := readInputValue(scanner)
	if ok := ip.CheckIP(ipString); !ok {
		log.Fatal("Invalid IP address")
	}
	fmt.Println("Enter Server Port address you want to connect:")
	portString := readInputValue(scanner)
	if ok := ip.CheckPort(portString); !ok {
		log.Fatal("Invalid Port address")
	}
	if strings.ContainsAny(ipString, ":") { //for IPv6 address
		ipString = "[" + ipString + "]"
	}
	return ipString + ":" + portString
}

// FunctionName - should return name of calling function
func FunctionName() string {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception Occurred and Recovered in FunctionName(), Error Info: ", errD)
		}
	}()
	pc, _, _, _ := runtime.Caller(1)
	funcName := strings.TrimSuffix(runtime.FuncForPC(pc).Name(), ".func1") // This is for defer function
	funcName = strings.TrimSuffix(funcName, ".1")                          // This is for go runtine function
	return funcName
}

// RecoverExceptionDetails will take one formal parameter as function name - should return exeception detail formated as: packageName.functionName:lineNumber
// Each format detail appended  with "<<" if it is multiple stack frames(LIFO). For example: packageName.functionName:lineNumber << packageName.functionName:lineNumber
func RecoverExceptionDetails(strfuncName string) string {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception Occurred and Recovered in RecoverExceptionDetails(), Error Info: ", errD)
		}
	}()
	var output string
	flag := false
	for skip := 1; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		strfunctionName := runtime.FuncForPC(pc).Name()
		if strings.Contains(file, "/runtime/") && strings.Contains(strfunctionName, "runtime.") {
			flag = true
			continue
		}
		if flag && strings.HasSuffix(file, ".go") {
			output += strfunctionName + ":" + strconv.Itoa(line) + " << "
			if strfuncName == strfunctionName {
				output = strings.TrimSuffix(output, " << ")
				break
			}
		}
	}
	return output
}
