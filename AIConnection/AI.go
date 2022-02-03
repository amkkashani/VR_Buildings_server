package AIConnection

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var sensors map[string][]string
var sensorsLock sync.RWMutex

var sensorsResults map[string]int
var resultLock sync.RWMutex

const (
	PredictionSize = 15
	pythonUrl      = "http://localhost:9000"
)

func AddSensorData(sensorName string, x, y, z string) {
	if sensors == nil {
		sensors = make(map[string][]string)
	}
	newVector := []string{x, y, z}
	if _, find := sensors[sensorName]; find {
		bigVector := append(sensors[sensorName], newVector...)

		if len(bigVector) > PredictionSize {
			bigVector = bigVector[len(bigVector)-PredictionSize-1 : len(bigVector)-1]
		}

		if len(bigVector) == 15 {
			PredictByPython(bigVector)
		}
		sensors[sensorName] = bigVector
	} else {
		sensors[sensorName] = newVector
	}
}

func PredictByPython(vector []string) {
	MyStr := ""
	fmt.Println(vector)
	for i, _ := range vector {
		MyStr = fmt.Sprintf("%s %d", MyStr, vector[i])
	}
	fmt.Println(MyStr)
	reqBody := bytes.NewBuffer([]byte(MyStr))
	resp, err := http.Post(pythonUrl, "text/html; charset=utf-8", reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
}
