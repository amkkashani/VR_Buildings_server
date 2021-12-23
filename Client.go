package main

import (
	"bufio"
	"fmt"
	"net"
)

//func main() {
//	fmt.Println("client test")
//	conn, err := net.Dial("udp", "127.0.0.1:1053")
//	if err != nil {
//		fmt.Printf("Some error %v", err)
//		return
//	}
//	fmt.Fprintf(conn, "%s", "connect")
//	fmt.Fprintf(conn, "%s", "connect")
//	fmt.Fprintf(conn, "%s",
//		"80.923: 0xâ€¦00C9 location data: position: x=367 y=-117 z=121 q=49; distances: 21C0 "+
//			"distance=Distance{length=296, quality=100}, 08B9 distance=Distance{length=315, quality=100},"+
//			" 14E6 distance=Distance{length=1882, quality=100}, 1522 distance=Distance{length=2144, quality=100}")
//
//	//n, err := bufio.NewReader(conn).Read(p)
//	//if err == nil {
//	//	fmt.Printf("%s\n", p[0:n])
//	//} else {
//	//	fmt.Printf("Some error %v\n", err)
//	//}
//
//	//go listen(conn)
//
//	var m sync.WaitGroup
//	m.Add(1)
//	m.Wait()
//
//	conn.Close()
//}

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
