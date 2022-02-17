package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

func main5() {
	fmt.Println("client test")
	conn, err := net.Dial("tcp", "185.110.189.249:3333")
	//conn, err := net.Dial("udp", "127.0.0.1:15493")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	//fmt.Fprintf(conn, "%s", "connect")
	//fmt.Fprintf(conn, "%s", "connect")
	//fmt.Fprintf(conn, "%s",
	//	"80.923: 0xâ€¦a4542 location data: position: x=70000 y=2000 z=1000 q=49; distances: 21C0 "+
	//		"distance=Distance{length=296, quality=100}, 08B9 distance=Distance{length=315, quality=100},"+
	//		" 14E6 distance=Distance{length=1882, quality=100}, 1522 distance=Distance{length=2144, quality=100}")

	t1 := []byte{36, 36, 71, 80, 43, 0, 218, 49, 235, 76, 50, 184, 73, 64, 99, 17, 253, 185, 44, 216, 65, 64, 238, 236, 147, 68, 1, 141, 151, 14, 63, 50, 48, 50, 50, 48, 50, 49, 54, 49, 57, 50, 54, 48, 51, 75, 0, 0, 0, 35, 140, 13, 10}
	conn.Write(t1)

	//n, err := bufio.NewReader(conn).Read(p)
	//if err == nil {
	//	fmt.Printf("%s\n", p[0:n])
	//} else {
	//	fmt.Printf("Some error %v\n", err)
	//}

	//go listen(conn)

	var m sync.WaitGroup
	m.Add(1)
	m.Wait()

	conn.Close()
}

func listen(conn net.Conn) {
	for {
		p := make([]byte, 4096)
		n, err := bufio.NewReader(conn).Read(p)
		if err == nil {
			fmt.Printf("%s\n", p[0:n])
		} else {
			fmt.Printf("Some error %v\n", err)
		}
	}
}
