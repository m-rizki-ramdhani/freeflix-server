package main

import (
	"freeflix/config"
	"freeflix/server"
	"sync"
)

func main() {
	appConfig := config.LoadConfig()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		server.StartGRPCServer(appConfig.GRPC_PORT)
	}()

	go func() {
		defer wg.Done()
		server.StartRESTServer(appConfig.REST_PORT)
	}()

	wg.Wait()
}
