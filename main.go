package main

import (
	"locationServer/RestAPI"
	"locationServer/SensorConnections"
	"locationServer/TcpGpsServer"
	"locationServer/UDPServer"
	"sync"
)

func main() {

	// open 1053 for udp receiver :)
	// open 1054 for udp unity app  :)
	go SensorConnections.SensorReceiver()
	go TcpGpsServer.RunGpsTcpServer() // on port 3333
	go RestAPI.RunRestApi()           //old api
	go UDPServer.Server3()            // unity server 1053

	//never exit :)
	var m sync.WaitGroup
	m.Add(1)
	m.Wait()
}

//
//func test()  {
//	var fff float64 = 15514.5588484844
//	xx := fmt.Sprintf("%.15f", fff)
//	fmt.Println(xx)
//
//	var x float64
//	x = 40.716999
//	var buf [8]byte
//	binary.BigEndian.PutUint64(buf[:], math.Float64bits(x))
//	fmt.Println(buf)
//
//	responseAuth := "242447502B00310484FF22575C40ED96058F46FB364035DE2242016891ED3C32303231303932323039313035326400000023E00D0A"
//	decodedByteArray, err := hex.DecodeString(responseAuth)
//	if err != nil {
//		fmt.Println("Unable to convert hex to byte. ", err)
//	}
//	fmt.Println(decodedByteArray)
//	fmt.Println(string(decodedByteArray))
//	f := []byte{36, 71, 80, 43, 0, 49, 4, 132}
//	println(Float64frombytes(f))
//
//	startPoint := 6
//	fmt.Println(Float64frombytes(decodedByteArray[startPoint : startPoint+8]))
//
//	second := 14
//	fmt.Println(Float64frombytes(decodedByteArray[second : second+8]))
//
//	height := 22
//	fmt.Println(Float32frombytes(decodedByteArray[height : height+4]))
//
//
//	//sample bytes
//	t1 := []byte {36,36, 71, 80, 43 ,0 ,218, 49, 235, 76, 50, 184, 73, 64, 99, 17, 253, 185, 44, 216, 65, 64, 238, 236, 147, 68, 1 ,141, 151, 14, 63, 50 ,48 ,50 ,50 ,48 ,50 ,49 ,54 ,49 ,57 ,50, 54 ,48 ,51 ,75, 0, 0, 0 ,35 ,140 ,13, 10}
//	pivot := 6
//	fmt.Println(Float64frombytes(t1[pivot : pivot+8]))
//
//	p2 := 22
//	fmt.Println(Float32frombytes(t1[p2 : p2+ 4]))
//	x1 := fmt.Sprintf("%0.8f" , Float32frombytes(t1[p2 : p2+ 4]))
//	fmt.Println(x1)
//
//	x2 := fmt.Sprintf("%0.8f" , Float64frombytes(t1[pivot : pivot+8]))
//	fmt.Println(x2)
//}
//
//func Float64frombytes(bytes []byte) float64 {
//	bits := binary.LittleEndian.Uint64(bytes)
//	float := math.Float64frombits(bits)
//	return float
//}
//
//func Float32frombytes(bytes []byte) float32 {
//	bits := binary.LittleEndian.Uint32(bytes)
//	float := math.Float32frombits(bits)
//	return float
//}
