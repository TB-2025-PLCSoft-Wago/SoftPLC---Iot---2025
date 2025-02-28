// Package inputUpdate This package is used to update the inputs from the HAL with a http request.
package inputUpdate

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type InputsOutputs struct {
	FriendlyName string
	Value        float64
	Service      string
	SubService   string
	id           string
}

var mutex sync.Mutex
var InputsOutputsState []InputsOutputs
var client mqtt.Client

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	//fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	splitTopic := strings.Split(msg.Topic(), "/")
	var payload map[string]interface{}
	err := json.Unmarshal(msg.Payload(), &payload)
	if err != nil {
		fmt.Println(err)
	}
	body := payload["dataPoints"].(map[string]interface{})
	mutex.Lock()
	for i, input := range InputsOutputsState {
		if input.id == splitTopic[0] && input.Service == splitTopic[1] {
			for key, value := range body {
				if key == input.SubService {
					InputsOutputsState[i].Value = value.(map[string]interface{})["value"].(float64)
				}
			}
		}
	}
	mutex.Unlock()
}

func sub(topic string) {
	token := client.Subscribe(topic, 1, messagePubHandler)
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("Failed to subscribe to topic: %v\n", token.Error())
		return
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)
}

func initClient() {
	var broker = "192.168.1.175"
	var port = 1884
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("ws://%s:%d/ws", broker, port))
	opts.SetClientID("go-simple")
	opts.OnConnect = connectHandler
	opts.SetDefaultPublishHandler(messagePubHandler)
	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
func InitInputs() {
	mutex.Lock()
	initClient()
	var result map[string]interface{}
	var tabResult []map[string]interface{}
	// Get the inputs from the HAL (DI, DO, AI, AO, Temp)
	resp, err := http.Get("http://192.168.1.175:8888/api/v1/hal/io")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	for key, value := range result {
		for i, v := range value.([]interface{}) {
			var val float64
			if v == true {
				val = 1
			} else if v == false {
				val = 0
			} else {
				val = v.(float64)
			}
			InputsOutputsState = append(InputsOutputsState, InputsOutputs{
				FriendlyName: "",
				Value:        val,
				Service:      strings.ToUpper(key) + strconv.Itoa(i+1),
				SubService:   "",
				id:           "",
			})
		}
	}

	//Get the appliances from the HAL
	resp, err = http.Get("http://192.168.1.175:8888/api/v1/appliance/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &tabResult)
	if err != nil {
		panic(err)
	}
	var friendlyName []string
	var id []string
	var services [][]string
	for _, v := range tabResult {
		var servicesOfThisId []string
		friendlyName = append(friendlyName, v["friendlyName"].(string))
		id = append(id, v["id"].(string))
		for _, s := range v["services"].([]interface{}) {
			servicesOfThisId = append(servicesOfThisId, s.(string))
		}
		services = append(services, servicesOfThisId)
	}
	for i, actualId := range id {
		for _, v := range services[i] {
			res, err := http.Get("http://192.168.1.175:8888/api/v1/appliance/" + actualId + "/" + v)
			if err != nil {
				panic(err)
			}
			body, err = io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			result = nil
			err = json.Unmarshal(body, &result)
			if err != nil {
				panic(err)
			}
			res.Body.Close()
			if len(result["dataPoints"].(map[string]interface{})) != 0 {
				for key, value := range result["dataPoints"].(map[string]interface{}) {
					InputsOutputsState = append(InputsOutputsState, InputsOutputs{
						FriendlyName: friendlyName[i],
						Value:        value.(map[string]interface{})["value"].(float64),
						Service:      v,
						SubService:   key,
						id:           actualId,
					})
				}
			}
		}
	}
	for i, actualId := range id {
		for _, v := range services[i] {
			if strings.Contains(v, "sgr") {
				sub(actualId + "/" + v + "/@UPDATE")
			}
		}
	}
	mutex.Unlock()
	fmt.Println("IO init finished")
}

func UpdateInputs() {
	mutex.Lock()
	var result map[string]interface{}
	resp, err := http.Get("http://192.168.1.175:8888/api/v1/hal/io")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	for key, value := range result {
		for i, v := range value.([]interface{}) {
			for j, input := range InputsOutputsState {
				if input.Service == strings.ToUpper(key)+strconv.Itoa(i+1) {
					if v == true {
						InputsOutputsState[j].Value = 1
					} else if v == false {
						InputsOutputsState[j].Value = 0
					} else {
						InputsOutputsState[j].Value = v.(float64)
					}
				}
			}
		}
	}

	/*for i, input := range InputsOutputsState {
		result = nil
		if input.FriendlyName != "" {
			resp, err := http.Get("http://192.168.1.175:8888/api/v1/appliance/" + input.id + "/" + input.Service)
			if err != nil {
				panic(err)
			}
			body, err = io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(body, &result)
			if err != nil {
				panic(err)
			}
			resp.Body.Close()
			for key, value := range result["dataPoints"].(map[string]interface{}) {
				if key == input.SubService {
					InputsOutputsState[i].Value = value.(map[string]interface{})["value"].(float64)
				}
			}
		}
	}*/
	mutex.Unlock()
}
