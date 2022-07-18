package TcpGpsServer

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"locationServer/SensorConnections"
	"locationServer/UDPServer"
	"log"
	"math"
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
			log.Println("Error Accepting +-*=")
			continue
		}
		// Handle connections in a new goroutine.
		go handleRequest2(conn)
	}
}

//func handleRequest(conn net.Conn) {
//	// Make a buffer to hold incoming data.
//	buf := make([]byte, 1024)
//	// Read the incoming connection into the buffer.
//	_, err := conn.Read(buf)
//	if err != nil {
//		fmt.Println("Error reading:", err.Error())
//	}
//
//	SensorConnections.Handle(buf)
//	// Send a response back to person contacting us.
//	conn.Write([]byte(string(buf) + "Message received."))
//	// Close the connection when you're done with it.
//	conn.Close()
//}

func handleRequest2(conn net.Conn) {
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		Handle([]byte(netData))
		fmt.Print("-> ", string(netData))

		responseAuth := "24244C473500010101697320746F6B656E000000000000000000000000000000000000000000000000000000000000000000000000000000000000237E0D0A"
		decodedByteArray, err := hex.DecodeString(responseAuth)
		if err != nil {
			fmt.Println("Unable to convert hex to byte. ", err)
		}
		conn.Write(decodedByteArray)
	}

}

func Handle(inputBytes []byte) {

	name := "gpsSensor"
	x, y, z, qos := gpsSensorParser(inputBytes, string(inputBytes))
	println("***", name, "****")

	//big log
	log.Println("**_** : " + string(inputBytes))
	log.Println(inputBytes)
	log.Println("input bytes\n---")

	if x == " " || x == "" {
		return
	}
	newSensorData := SensorConnections.Sensor{name, x, y, z, qos}
	SensorConnections.Sensors[name] = newSensorData

	//log in to file
	myLogJson, _ := json.Marshal(newSensorData)
	fmt.Println("--- ", string(myLogJson))
	log.Println(string(myLogJson))

	//msg := fmt.Sprintf("%s %s %s %s %d", name, x, y, z, qos)
	msg := string(myLogJson)
	fmt.Println("--- sended msg as string : >> " + msg)

	UDPServer.Publish(msg)
}

//gpsSensorParser return Latitude Longitude Elevation
func gpsSensorParser(decodedByteArray []byte, str string) (string, string, string, byte) {

	if strings.Contains(str, "GP") {
		if len(decodedByteArray) < 26 {
			// we cant resolve this msg
			return "", "", "", 0
		}
		latIndex := 6
		lat := float64frombytes(decodedByteArray[latIndex : latIndex+8])

		longIndex := 14
		long := float64frombytes(decodedByteArray[longIndex : longIndex+8])

		heightIndex := 22
		height := float32frombytes(decodedByteArray[heightIndex : heightIndex+4])
		fmt.Println()

		var qualiyOfService = decodedByteArray[26]

		return convertFloat64ToString(lat), convertFloat64ToString(long), convertFloat32ToString(height), qualiyOfService

	} else {
		return "", "", "", 0
	}
}

func convertFloat64ToString(f float64) string {
	return fmt.Sprintf("%.12f", f)
}

func convertFloat32ToString(f float32) string {
	return fmt.Sprintf("%.6f", f)
}

func float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func float32frombytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}
