package main

import (
	"locationServer/RestAPI"
	"locationServer/UDPServer"
	"sync"
)

func main() {
	// open 1053 for udp :)
	// open 8081 for tcp :)
	go UDPServer.Server3()
	go RestAPI.RunRestApi()
	var m sync.WaitGroup
	m.Add(1)
	m.Wait()
}