package SensorConnections

import (
	"encoding/json"
	"fmt"
	"locationServer/AIConnection"
	"locationServer/StringParser"
	"locationServer/UDPServer"
	"log"
	"net"
)

var Sensors map[string]Sensor
var LOG_DIR = "logs/logs.txt"
var AI_IS_ACTIVE = true

func SensorReceiver() {
	// listen to incoming udp packets
	pc, err := net.ListenPacket("udp", ":1053")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
		go serve(pc, addr, buf[:n])
	}

}

func serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	msg := string(buf)
	fmt.Println(msg, " **** ", len(msg))
	handle(buf)
	pc.WriteTo(buf, addr)
}

func handle(inputBytes []byte) {
	name, x, y, z := StringParser.StringParser(string(inputBytes))
	newSensorData := Sensor{name, x, y, z}
	Sensors[name] = newSensorData

	//log in to file
	myLogJson, _ := json.Marshal(newSensorData)
	fmt.Println("--- ", string(myLogJson))
	log.Println(string(myLogJson))

	msg := fmt.Sprintf("%s %d %d %d", name, x, y, z)

	if AI_IS_ACTIVE {
		AIConnection.AddSensorData(newSensorData.Name, newSensorData.X, newSensorData.Y, newSensorData.Z)
	}

	//UDPServer.BroadCastToAll(msg)
	UDPServer.Publish(msg)
}

type Sensor struct {
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Z    int    `json:"z"`
}
