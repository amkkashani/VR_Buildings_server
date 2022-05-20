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
var AI_IS_ACTIVE = false

func SensorReceiver() {
	Sensors = make(map[string]Sensor)
	// listen to incoming udp packets
	pc, err := net.ListenPacket("udp", ":15493")
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
	println(msg)
	Handle(buf)
	pc.WriteTo(buf, addr)
}

func Handle(inputBytes []byte) {
	name, x, y, z := StringParser.StringParser(string(inputBytes))
	println("***", name, "****")

	//big log
	log.Println("**_** : " + string(inputBytes))

	if name == " " || name == "" {
		return
	}
	newSensorData := Sensor{name, x, y, z, 1}
	Sensors[name] = newSensorData

	//log in to file
	myLogJson, _ := json.Marshal(newSensorData)
	fmt.Println("--- ", string(myLogJson))
	log.Println(string(myLogJson))

	msg := fmt.Sprintf("%s %s %s %s", name, x, y, z)
	fmt.Println(msg)

	if AI_IS_ACTIVE {
		AIConnection.AddSensorData(newSensorData.Name, newSensorData.X, newSensorData.Y, newSensorData.Z)
	}

	UDPServer.Publish(msg)
}

type Sensor struct {
	Name string `json:"name"`
	X    string `json:"x"`
	Y    string `json:"y"`
	Z    string `json:"z"`
	Qos  byte   `json:"qos"`
}
