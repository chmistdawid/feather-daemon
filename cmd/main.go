package main

import (
	"fmt"
	"log"
	"time"

	"github.com/samekigor/feather-deamon/internal/config"
	"github.com/samekigor/feather-deamon/internal/logger"
	"github.com/samekigor/feather-deamon/internal/socket"
)

func main() {
	config.LoadConfigFile()
	logger.SetupLogging()
	defer logger.CloseLogging()
	fmt.Println("Starting GRPC server")
	go socket.StartGRPCServer()

	go socket.HandleSignals()
	fmt.Println("GRPC server started")
	for {
		log.Println("Running...")
		time.Sleep(1 * time.Second)
	}

}
