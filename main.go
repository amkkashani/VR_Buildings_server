package main

import (
	"locationServer/SensorConnections"
	"locationServer/UDPServer"
	"sync"
)

func main() {
	// open 1053 for udp receiver :)
	// open 1054 for udp unity app  :)
	go SensorConnections.SensorReceiver()
	go UDPServer.Server3() // unity server
	//go RestAPI.RunRestApi() //old api

	//never exit :)
	var m sync.WaitGroup
	m.Add(1)
	m.Wait()
}
