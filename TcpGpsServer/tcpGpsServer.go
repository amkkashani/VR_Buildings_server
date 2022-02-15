package TcpGpsServer

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"locationServer/SensorConnections"
	"locationServer/StringParser"
	"locationServer/UDPServer"
	"log"
	"net"
	"os"
	"strings"
)

const (
	CONN_HOST = ""
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func RunGpsTcpServer() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest2(conn)
	}
}

func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	SensorConnections.Handle(buf)
	// Send a response back to person contacting us.
	conn.Write([]byte(string(buf) + "Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}

func handleRequest2(conn net.Conn) {
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		SensorConnections.Handle([]byte(netData))
		fmt.Print("-> ", string(netData))

		hexString := "24244C473500010101697320746F6B656E000000000000000000000000000000000000000000000000000000000000000000000000000000000000237E0D0A"
		decodedByteArray, err := hex.DecodeString(hexString)
		if err != nil {
			fmt.Println("Unable to convert hex to byte. ", err)
		}
		conn.Write(decodedByteArray)
	}

}

func Handle(inputBytes []byte) {
	name, x, y, z := StringParser.StringParser(string(inputBytes))
	println("***", name, "****")

	//big log
	log.Println("**_** : " + string(inputBytes))

	if name == " " || name == "" {
		return
	}
	newSensorData := SensorConnections.Sensor{name, x, y, z}
	SensorConnections.Sensors[name] = newSensorData

	//log in to file
	myLogJson, _ := json.Marshal(newSensorData)
	fmt.Println("--- ", string(myLogJson))
	log.Println(string(myLogJson))

	msg := fmt.Sprintf("%s %s %s %s", name, x, y, z)
	fmt.Println(msg)

	UDPServer.Publish(msg)
}
