package UDPServer

//var addrs []*net.UDPConn
//var addrLock sync.RWMutex

//func RunUdpSrever()  {
//	//Basic variables
//	port := ":5050"
//	protocol := "udp"
//
//	addrs = make([]*net.UDPConn,0)
//
//	//Build the address
//	udpAddr, err := net.ResolveUDPAddr(protocol, port)
//	if err != nil {
//		fmt.Println("Wrong Address")
//		return
//	}
//
//	//Output
//	fmt.Println("Coded by Amir Hosein Kashani\nReading " + protocol + " from " + udpAddr.String())
//
//	//Create the connection
//	udpConn, err := net.ListenUDP(protocol, udpAddr)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	//Keep calling this function
//	for {
//		go display(udpConn)
//	}
//}
//
//
//func display(conn *net.UDPConn) {
//	var buf [4096]byte
//	n,addr,err := conn.ReadFromUDP(buf[0:])
//	if err != nil {
//		fmt.Println("Error Reading")
//		return
//	} else {
//		fmt.Println(hex.EncodeToString(buf[0:n]))
//		fmt.Println("Package Done")
//	}
//	conn.WriteToUDP([]byte("dadash resid"),addr)
//	fmt.Println(addr.String())
//
//	if checkIsNew(conn) {
//		addrLock.Lock()
//		addrs = append(addrs, conn)
//		addrLock.Unlock()
//	}
//
//}
//
//func checkIsNew(conn *net.UDPConn)bool  {
//	addrLock.RLock()
//	defer addrLock.RUnlock()
//	for i := 0; i < len(addrs) ; i++ {
//		if addrs[i] == conn {
//			fmt.Println("repeated address")
//			return false
//		}
//	}
//	return true
//}
//
//func BroadCastToAll(s string)  {
//	addrLock.RLock()
//	defer addrLock.RUnlock()
//	for i := 0; i < len(addrs); i++ {
//		fmt.Fprintf(addrs[i], s)
//	}
//
//}
