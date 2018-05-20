package main

import (
	"fmt"
	"github.com/ludovicose/go_microservice/accountservice/services"
	"github.com/ludovicose/go_microservice/accountservice/dbclient"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	services.StartWebServer("6767")
}

func initializeBoltClient() {
	services.DBClient = &dbclient.BoltClient{}
	services.DBClient.OpenBoldDB()
	services.DBClient.Seed()
}
