package main

import (
	"locationServer/RestAPI"
	"locationServer/UDPServer"
	"sync"
)

func main() {
	go UDPServer.Server3()
	go RestAPI.RunRestApi()
	var m sync.WaitGroup
	m.Add(1)
	m.Wait()
}