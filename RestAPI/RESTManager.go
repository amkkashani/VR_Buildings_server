package RestAPI

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"locationServer/AIConnection"
	"locationServer/StringParser"
	"locationServer/UDPServer"
	"log"
	"net/http"
	"os"
)

var Sensors map[string]Sensor
var LOG_DIR  = "logs/logs.txt"


func RunRestApi() {
	logFile, err := os.OpenFile(LOG_DIR, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil{
		panic(err)
	}
	log.SetOutput(logFile)

	Sensors = make(map[string]Sensor)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/Sensor", func(w http.ResponseWriter, r *http.Request){
		if r.Method == "POST" {
			inputBytes , _ :=  ioutil.ReadAll(r.Body)
			w.WriteHeader(200)
			fmt.Println(inputBytes)
			name , x ,y,z :=StringParser.StringParser(string(inputBytes))
			newSensorData := Sensor{name,x,y,z}
			Sensors[name] = newSensorData

			//log in to file
			myLogJson, _ := json.Marshal(Sensors)
			fmt.Println("--- " , string(myLogJson) )
			log.Println(string(myLogJson))

			fmt.Println(Sensors)
			msg := fmt.Sprintf("%s %d %d %d" , name , x , y ,z)
			AIConnection.AddSensorData(newSensorData.Name,newSensorData.X,newSensorData.Y,newSensorData.Z)
			//UDPServer.BroadCastToAll(msg)
			UDPServer.Publish(msg)
			fmt.Fprintf(w,"tamom shode , tasir gozar bood :)")

		}
		//if r.Method == "POST" {
		//	inputBytes , _ :=  ioutil.ReadAll(r.Body)
		//var req map[string]interface{}
		//
		//json.Unmarshal(inputBytes, &req)
		//
		//
		//if name, ok := req["name"]; ok {
		//	x , c1 := req["x"].(float64)
		//	y , c2 := req["y"].(float64)
		//	z , c3 := req["z"].(float64)
		//	if c1 && c2 && c3 == false {
		//		http.Error(w,"400 inputs are not compeleted", http.StatusBadRequest)
		//		println(name,x , y ,z)
		//	}
		//}else {
		//	http.Error(w,"400 inputs are not compeleted", http.StatusBadRequest)
		//}
		//}else if r.Method == "GET"{
		//	http.Error(w,"400 inputs are not compeleted", http.StatusBadRequest)
		//}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}

type Sensor struct {
	Name string `json:"name"`
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}


//func test()  {
//	empJson := `{
//        "id" : 'as;ldfja',
//        "name" : "Irshad",
//        "department" : "IT",
//        "designation" : "Product Manager"
//	}`
//
//	// Declared an empty interface
//	var req map[string]interface{}
//
//	// Unmarshal or Decode the JSON to the interface.
//	json.Unmarshal([]byte(empJson), &req)
//	val := req["id"]
//
//
//	//iAreaId := val.(int)
//	iAreaId, ok := val.(int) // Alt. non panicking version
//	if ok == false {
//		println("ok nashod")
//	}
//	fmt.Println("id : " , iAreaId + 4 )
//}