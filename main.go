package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"locationServer/RestAPI"
	"locationServer/SensorConnections"
	"locationServer/TcpGpsServer"
	"locationServer/UDPServer"
	"math"
	"sync"
)

func main() {
	var fff float64 = 15514.5588484844
	xx := fmt.Sprintf("%.15f", fff)
	fmt.Println(xx)

	var x float64
	x = 40.716999
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(x))
	fmt.Println(buf)

	responseAuth := "242447502B00310484FF22575C40ED96058F46FB364035DE2242016891ED3C32303231303932323039313035326400000023E00D0A"
	decodedByteArray, err := hex.DecodeString(responseAuth)
	if err != nil {
		fmt.Println("Unable to convert hex to byte. ", err)
	}
	fmt.Println(decodedByteArray)
	fmt.Println(string(decodedByteArray))
	f := []byte{36, 71, 80, 43, 0, 49, 4, 132}
	println(Float64frombytes(f))

	startPoint := 6
	fmt.Println(Float64frombytes(decodedByteArray[startPoint : startPoint+8]))

	second := 14
	fmt.Println(Float64frombytes(decodedByteArray[second : second+8]))

	height := 22
	fmt.Println(Float32frombytes(decodedByteArray[height : height+4]))

	// open 1053 for udp receiver :)
	// open 1054 for udp unity app  :)
	go SensorConnections.SensorReceiver()
	go TcpGpsServer.RunGpsTcpServer() // on port 3333
	go RestAPI.RunRestApi()           //old api
	go UDPServer.Server3()            // unity server

	//never exit :)
	var m sync.WaitGroup
	m.Add(1)
	m.Wait()
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float32frombytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}
