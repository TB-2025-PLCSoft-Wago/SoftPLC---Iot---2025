// Package inputUpdate This package is used to update the inputs from the HAL with a http request.
package inputUpdate

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type InputsOutputs struct {
	FriendlyName string
	Value        string
	Service      string
	SubService   string
	id           string
}

var mutex sync.Mutex
var InputsOutputsState []InputsOutputs

// var client mqtt.Client
var urlsID []string
var postID string //monitoring list ID
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
	for i, input := range InputsOutputsState {
		if input.id == splitTopic[0] && input.Service == splitTopic[1] {
			for key, value := range body {
				if key == input.SubService {
					valFloat := value.(map[string]interface{})["value"].(float64)
					valStr := strconv.FormatFloat(valFloat, 'f', -1, 64)
					InputsOutputsState[i].Value = valStr
				}
			}
		}
	}

}

/*
func sub(topic string) {
	token := client.Subscribe(topic, 1, messagePubHandler)
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("Failed to subscribe to topic: %v\n", token.Error())
		return
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)
}*/

func configureDO() {
	username := "admin"
	password := "wago"
	url := "https://192.168.37.134/wda/parameters/0-0-io-channelcompositions-4-channels"
	// Create the JSON of body
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"id":   "0-0-io-channelcompositions-4-channels",
			"type": "parameters",
			"attributes": map[string]interface{}{
				"value": []int{
					9, 10, 11, 12, 13, 14, 15, 16,
				},
			},
		},
	}

	// JSON serialization
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/vnd.api+json")

	client := createHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		fmt.Println("Error configure output DO")
	}
	resp.Body.Close()
}

func configureControleMode(nb uint8) {
	username := "admin"
	password := "wago"
	url := "https://192.168.37.134/wda/parameters/0-0-io-iocheckaccessmode"
	// Create the JSON of body
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"id":   "0-0-io-iocheckaccessmode",
			"type": "parameters",
			"attributes": map[string]interface{}{
				"value": nb,
			},
		},
	}

	// JSON serialization
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/vnd.api+json")

	/*
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}*/
	client := createHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		fmt.Println("Error configure controle mode")
	}
	resp.Body.Close()
}

func CreateMonitoringLists() {
	username := "admin"
	password := "wago"
	url := "https://192.168.37.134/wda/monitoring-lists"
	// Create the JSON of body
	var parameters []map[string]interface{}

	for _, id := range urlsID {
		parameters = append(parameters, map[string]interface{}{
			"id":   id,
			"type": "parameters",
		})
	}
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"type": "monitoring-lists",
			"attributes": map[string]interface{}{
				"timeout": 600,
			},
			"relationships": map[string]interface{}{
				"parameters": map[string]interface{}{
					"data": parameters,
				},
			},
		},
	}

	// JSON serialization
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/vnd.api+json")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	//client := createHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	postID, err = getMonitoringListID(resp.Body)
	if resp.StatusCode != 201 {
		fmt.Println("Error createMonitoringLists not 201")
		panic(resp.StatusCode)
	}
	resp.Body.Close()

}
func getMonitoringListID(respBody io.Reader) (string, error) {
	var responseData map[string]interface{}

	err := json.NewDecoder(respBody).Decode(&responseData)
	if err != nil {
		return "", fmt.Errorf("JSON decoding failed: %w", err)
	}

	// data -> id
	if data, ok := responseData["data"].(map[string]interface{}); ok {
		if id, ok := data["id"].(string); ok {
			return id, nil
		}
	}

	return "", fmt.Errorf("ID not found in response")
}
func initClient() {
	/*
		var broker = "192.168.37.134"
		var port = 1884
		opts := mqtt.NewClientOptions()
		opts.AddBroker(fmt.Sprintf("ws://%s:%d/ws", broker, port))
		opts.SetClientID("go-simple")
		opts.OnConnect = connectHandler
		opts.SetDefaultPublishHandler(messagePubHandler)
		client = mqtt.NewClient(opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}*/
	generateURLsId()
	configureControleMode(0)
	configureDO()
	configureControleMode(2)
	CreateMonitoringLists()
}

/*
	func InitInputs() {
		mutex.Lock()
		initClient()
		var result map[string]interface{}
		var tabResult []map[string]interface{}
		// Get the inputs from the HAL (DI, DO, AI, AO, Temp)
		resp, err := http.Get("http://192.168.37.134:8888/api/v1/hal/io")
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
				var val string
				if v == true {
					val = "1"
				} else if v == false {
					val = "0"
				} else {
					valFloat := v.(float64)
					valStr := strconv.FormatFloat(valFloat, 'f', -1, 64)
					val = valStr
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
		resp, err = http.Get("http://192.168.37.134:8888/api/v1/appliance/")
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
				res, err := http.Get("http://192.168.37.134:8888/api/v1/appliance/" + actualId + "/" + v)
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
						valFloat := value.(map[string]interface{})["value"].(float64)
						valStr := strconv.FormatFloat(valFloat, 'f', -1, 64)

						InputsOutputsState = append(InputsOutputsState, InputsOutputs{
							FriendlyName: friendlyName[i],
							Value:        valStr,
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
*/

// Generates all necessary URLs for IO values
func generateURLsId() {
	baseURL := "0-0-io-channels-"
	// divalue (21 to 28)
	for i := 21; i <= 28; i++ {
		urlsID = append(urlsID, fmt.Sprintf("%s%d-divalue", baseURL, i))
	}
	// dovalue (9 to 16)
	for i := 9; i <= 16; i++ {
		urlsID = append(urlsID, fmt.Sprintf("%s%d-dovalue", baseURL, i))
	}
	// aivalue (31 to 32)
	for i := 31; i <= 32; i++ {
		urlsID = append(urlsID, fmt.Sprintf("%s%d-aivalue", baseURL, i))
	}
	// aovalue (19 to 20)
	for i := 19; i <= 20; i++ {
		urlsID = append(urlsID, fmt.Sprintf("%s%d-aovalue", baseURL, i))
	}
	// aivalue (29 to 30) -> TEMP (X13)
	for i := 29; i <= 30; i++ {
		urlsID = append(urlsID, fmt.Sprintf("%s%d-aivalue", baseURL, i))
	}
}

/* Get the inputs from the HAL (DI, DO, AI, AO, Temp) */
func InitInputs() {
	mutex.Lock()
	defer mutex.Unlock()

	initClient()

	username := "admin"
	password := "wago"

	client := createHTTPClient()
	results := fetchValues(client, username, password)

	/*fmt.Println("Fetched values:")
	for url, val := range results {
		fmt.Printf("%s → %v\n", url, val)
	}
	*/
	mapResultsToInputsOutputs(results)

	//Get the appliances from the HAL
	/*
		var result map[string]interface{}
		resp2, err := http.Get("http://192.168.37.134:8888/api/v1/appliance/")
		if err != nil {
			panic(err)
		}
		defer resp2.Body.Close()
		body, err := io.ReadAll(resp2.Body)
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
				res, err := http.Get("http://192.168.37.134:8888/api/v1/appliance/" + actualId + "/" + v)
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
						valFloat := value.(map[string]interface{})["value"].(float64)
						valStr := strconv.FormatFloat(valFloat, 'f', -1, 64)

						InputsOutputsState = append(InputsOutputsState, InputsOutputs{
							FriendlyName: friendlyName[i],
							Value:        valStr,
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
	*/
	//
	//mutex.Unlock()
	fmt.Println("IO init finished")
}
func UpdateInputs() {
	mutex.Lock()
	defer mutex.Unlock() //unlock if return or panic
	// Authentication
	username := "admin"
	password := "wago"
	//generateURLsId()

	client := createHTTPClient()
	results := fetchValues(client, username, password)
	/*
		fmt.Println("Fetched values:")
		for url, val := range results {
			fmt.Printf("%s → %v\n", url, val)
		}
	*/
	updateInputsOutputsState(results) //Create InputsOutputsState

	//mutex.Unlock()
}

/*
func UpdateInputs() {
	mutex.Lock()
	var result map[string]interface{}
	resp, err := http.Get("http://192.168.37.134:8888/api/v1/hal/io")
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
						InputsOutputsState[j].Value = "1"
					} else if v == false {
						InputsOutputsState[j].Value = "0"
					} else {
						valFloat := v.(float64)
						valStr := strconv.FormatFloat(valFloat, 'f', -1, 64)
						InputsOutputsState[j].Value = valStr
					}
				}
			}
		}
	}

	/*for i, input := range InputsOutputsState {
		result = nil
		if input.FriendlyName != "" {
			resp, err := http.Get("http://192.168.37.134:8888/api/v1/appliance/" + input.id + "/" + input.Service)
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
	}*/ /*
	mutex.Unlock()
}*/

// Creates HTTP client with disabled TLS verification
func createHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			MaxIdleConns:        100,
			IdleConnTimeout:     90 * time.Second,
			DisableCompression:  false,
			MaxConnsPerHost:     0, // 0 = illimitable
			MaxIdleConnsPerHost: 10,
		},
	}
}

// Fetches JSON values from given URLs
func fetchValues(client *http.Client, username string, password string) map[string]interface{} {
	results := make(map[string]interface{})
	//start := time.Now()

	url := fmt.Sprintf("https://192.168.37.134/wda/monitoring-lists/%s/parameters", postID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Request creation failed:", err)
		return results
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/vnd.api+json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return results
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return results
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("JSON parse error:", err)
		return results
	}

	data, ok := response["data"].([]interface{})
	if !ok {
		fmt.Println("Unexpected JSON structure: missing `data` array")
		return results
	}

	for _, item := range data {
		param, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		id, _ := param["id"].(string)
		attrs, ok := param["attributes"].(map[string]interface{})
		if !ok {
			continue
		}
		val := attrs["value"] //extracts the "value" field
		results[id] = val
	}

	//fmt.Printf("Fetched %d values in %s\n", len(results), time.Since(start))
	return results
}

// Builds InputsOutputsState from fetched results (used in InitInputs)
func mapResultsToInputsOutputs(results map[string]interface{}) {
	for key, v := range results {
		parts := strings.Split(key, "-")
		if len(parts) < 6 {
			fmt.Println("Unexpected URL format:", key)
			continue
		}

		idStr := parts[4]
		id, _ := strconv.Atoi(idStr)
		if id >= 9 && id <= 16 {
			id = id - 8 //do
		} else if id >= 21 && id <= 28 {
			id = id - 20 //di
		} else if id == 29 || id == 30 {
			id = id - 28 //ai (x13)
		} else if id == 31 || id == 32 {
			id = id - 28 //ai (x14)
		} else if id == 17 || id == 18 {
			id = id - 16 //ai (x6)
		} else if id == 19 || id == 20 {
			id = id - 18 //ao (x6)
		} else {
			fmt.Println("Undefined I/O : ", key)
			continue
		}

		ioType := parts[5]
		k := strings.TrimSuffix(ioType, "value") // example : "aiValue" --> ai

		var val string
		if v == true {
			val = "1"
		} else if v == false {
			val = "0"
		} else {
			valFloat := v.(float64)
			valStr := strconv.FormatFloat(valFloat, 'f', -1, 64)
			val = valStr
		}
		InputsOutputsState = append(InputsOutputsState, InputsOutputs{
			FriendlyName: "",
			Value:        val,
			Service:      strings.ToUpper(k) + strconv.Itoa(id),
			SubService:   "",
			id:           "",
		})
	}
}

// Updates InputsOutputsState in place (used in UpdateInputs)
func updateInputsOutputsState(results map[string]interface{}) {
	//fmt.Println("updateInputsOutputsState")
	for key, v := range results {
		parts := strings.Split(key, "-")
		if len(parts) < 6 {
			fmt.Println("Unexpected URL format:", key)
			continue
		}

		idStr := parts[4]
		ioType := parts[5]
		k := strings.TrimSuffix(ioType, "value") // donne "ai"

		id, _ := strconv.Atoi(idStr)
		if id >= 9 && id <= 16 {
			id = id - 8 //do
		} else if id >= 21 && id <= 28 {
			id = id - 20 //di (x12)
		} else if id == 29 || id == 30 {
			id = id - 28 //ai (x13)
		} else if id == 31 || id == 32 {
			id = id - 28 //ai (x14)
		} else if id == 17 || id == 18 {
			id = id - 16 //ai (x6)
		} else if id == 19 || id == 20 {
			id = id - 18 //ao (x6)
		} else {
			fmt.Println("Undefined I/O : ", key)
			continue
		}

		for j, input := range InputsOutputsState {
			if input.Service == strings.ToUpper(k)+strconv.Itoa(id) {
				if v == true {
					InputsOutputsState[j].Value = "1"
				} else if v == false {
					InputsOutputsState[j].Value = "0"
				} else {
					valFloat := v.(float64)
					valStr := strconv.FormatFloat(valFloat, 'f', -1, 64)
					InputsOutputsState[j].Value = valStr
				}
			}
		}
	}
}
