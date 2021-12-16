
package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

func main() {

	conn, err := net.Dial("udp", "127.0.0.1:1053")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn,"%s","connect")
	fmt.Fprintf(conn,"%s","connect")

	//n, err := bufio.NewReader(conn).Read(p)
	//if err == nil {
	//	fmt.Printf("%s\n", p[0:n])
	//} else {
	//	fmt.Printf("Some error %v\n", err)
	//}

	go listen(conn)

	var m sync.WaitGroup
	m.Add(1)
	m.Wait()

	conn.Close()
}

func listen(conn net.Conn) {
	for  {
		p :=  make([]byte, 4096)
		n, err := bufio.NewReader(conn).Read(p)
		if err == nil {
			fmt.Printf("%s\n", p[0:n])
		} else {
			fmt.Printf("Some error %v\n", err)
		}
	}
}