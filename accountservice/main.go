package main

import (
	"fmt"
	"github.com/ludovicose/go_microservice/accountservice/services"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	services.StartWebServer("6767")
}
