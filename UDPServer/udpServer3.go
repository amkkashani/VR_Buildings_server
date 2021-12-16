package UDPServer

import (
	"fmt"
	"log"
	"net"
	"sync"
)

var clients []*Client
var clientMubtex sync.RWMutex

func Server3() {

	// listen to incoming udp packets
	pc, err := net.ListenPacket("udp", ":1053")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	go publisher()

	for {
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
		go serve(pc, addr, buf[:n])
	}

}

func publisher() {
	for {
		var str string
		fmt.Scanf("%s", &str)
		fmt.Println(str)
		Publish(str)
	}
}

func serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	// 0 - 1: ID
	// 2: QR(1): Opcode(4)
	msg := string(buf)
	fmt.Println(msg, " **** ", len(msg))
	if msg == "connect" {

		if !checkIsRepeated(pc, addr) {
			addToClients(creatClient(addr, pc))
			pc.WriteTo([]byte("accept"), addr)
		}else{
			pc.WriteTo([]byte("refreshed"), addr)
		}

		return
	}
	pc.WriteTo(buf, addr)
}

func checkIsRepeated(pc net.PacketConn, addr net.Addr) bool {
	for i := 0; i < len(clients); i++ {
		if  clients[i].pc == pc &&  clients[i].addr.String() == addr.String() {
			return true
			fmt.Println("repeat")
		}
	}
	return false
}

func Publish(msg string) {
	for i := 0; i < len(clients); i++ {
		clients[i].sendToClient(msg)
	}
}

type Client struct {
	addr net.Addr
	pc   net.PacketConn
}

func creatClient(addr net.Addr, pc net.PacketConn) *Client {
	newCline := Client{addr, pc}
	return &newCline
}

func addToClients(c *Client) {
	clientMubtex.Lock()
	clients = append(clients, c)
	clientMubtex.Unlock()
}

func (c *Client) sendToClient(str string) {
	c.pc.WriteTo([]byte(str), c.addr)
}
