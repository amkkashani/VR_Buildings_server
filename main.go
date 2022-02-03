package main

import (
	"locationServer/RestAPI"
	"locationServer/SensorConnections"
	"locationServer/TcpGpsServer"
	"locationServer/UDPServer"
	"sync"
)

func main() {

	// open 1053 for udp receiver :)
	// open 1054 for udp unity app  :)
	go SensorConnections.SensorReceiver()
	go TcpGpsServer.RunGpsTcpServer() // on port 3333
	go RestAPI.RunRestApi()           //old api
	go UDPServer.Server3()            // unity server

	//never exit :)
	var m sync.WaitGroup
	m.Add(1)
	m.Wait()
}
