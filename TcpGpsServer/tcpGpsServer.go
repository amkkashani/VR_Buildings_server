package TcpGpsServer

import (
	"bufio"
	"fmt"
	"locationServer/SensorConnections"
	"net"
	"os"
	"strings"
)

const (
	CONN_HOST = "127.0.0.1"
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
		conn.Write([]byte(string(netData) + "Message received."))
	}
}
