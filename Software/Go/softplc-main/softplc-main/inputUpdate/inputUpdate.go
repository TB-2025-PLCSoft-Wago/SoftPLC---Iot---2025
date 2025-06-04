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
var urls []string

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
					valFloat := value.(map[string]interface{})["value"].(float64)
					valStr := strconv.FormatFloat(valFloat, 'f', -1, 64)
					InputsOutputsState[i].Value = valStr
				}
			}
		}
	}
	mutex.Unlock()
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
		fmt.Println("Error configure output DO")
	}
	resp.Body.Close()
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
	configureControleMode(0)
	configureDO()
	configureControleMode(2)
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
func generateURLs() {
	baseURL := "https://192.168.37.134/wda/parameters/0-0-io-channels-"
	// divalue (21 to 28)
	for i := 21; i <= 28; i++ {
		urls = append(urls, fmt.Sprintf("%s%d-divalue", baseURL, i))
	}
	// dovalue (9 to 16)
	for i := 9; i <= 16; i++ {
		urls = append(urls, fmt.Sprintf("%s%d-dovalue", baseURL, i))
	}
	// aivalue (31 to 32)
	for i := 31; i <= 32; i++ {
		urls = append(urls, fmt.Sprintf("%s%d-aivalue", baseURL, i))
	}
	// aovalue (19 to 20)
	for i := 19; i <= 20; i++ {
		urls = append(urls, fmt.Sprintf("%s%d-aovalue", baseURL, i))
	}
	// aivalue (29 to 30) -> TEMP (X13)
	for i := 29; i <= 30; i++ {
		urls = append(urls, fmt.Sprintf("%s%d-aivalue", baseURL, i))
	}
}

/* Get the inputs from the HAL (DI, DO, AI, AO, Temp) */
func InitInputs() {
	mutex.Lock()
	defer mutex.Unlock()

	initClient()

	username := "admin"
	password := "wago"

	generateURLs()

	client := createHTTPClient()
	results := fetchValues(client, username, password, urls)

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
	//mutex.Unlock()
	fmt.Println("IO init finished")
}
func UpdateInputs() {
	mutex.Lock()
	defer mutex.Unlock() //unlock if return or panic
	// Authentification
	username := "admin"
	password := "wago"
	//generateURLs()

	client := createHTTPClient()
	results := fetchValues(client, username, password, urls)
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
			MaxConnsPerHost:     0, // 0 = illimité
			MaxIdleConnsPerHost: 10,
		},
	}
}

// Fetches JSON values from given URLs
func fetchValues(client *http.Client, username, password string, urls []string) map[string]interface{} {
	var (
		results = make(map[string]interface{})
		wg      sync.WaitGroup
		mu      sync.Mutex // pour protéger `results`
	)
	sem := make(chan struct{}, 1) // max 4 requêtes en parallèle
	start := time.Now()           //test
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)

			sem <- struct{}{}        // acquérir un slot
			defer func() { <-sem }() // libérer le slot

			value, err := fetchValue(url, username, password, client)
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", url, err)
				return
			}

			mu.Lock()
			results[url] = value
			mu.Unlock()
		}(url)
		/*
			fmt.Println("Requesting:", url)

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println("Request creation error:", err)
				continue
			}

			req.SetBasicAuth(username, password)
			req.Header.Set("Content-Type", "application/vnd.api+json")

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("HTTP request error:", err)
				continue
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Response read error:", err)
				continue
			}

			value := parseResponse(body, url)
			results[url] = value
		*/
	}
	wg.Wait()

	duration := time.Since(start)
	fmt.Printf("Total time for fetchValues: %s\n", duration)
	return results
}
func fetchValue(url, username, password string, client *http.Client) (interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("request creation failed: %w", err)
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/vnd.api+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read failed: %w", err)
	}

	var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("json parsing failed: %w", err)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d error from %s: %s", resp.StatusCode, url, string(body))
	}
	// safely extract nested value
	if d, ok := data["data"].(map[string]interface{}); ok {
		if attr, ok := d["attributes"].(map[string]interface{}); ok {
			if value, ok := attr["value"]; ok {
				return value, nil
			}
		}
	}
	return nil, fmt.Errorf("value field missing")
}

// Parses JSON and extracts the "value" field
func parseResponse(body []byte, url string) interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("JSON parsing error:", err)
		return nil
	}

	if dataField, ok := data["data"].(map[string]interface{}); ok {
		if attrField, ok := dataField["attributes"].(map[string]interface{}); ok {
			if value, ok := attrField["value"]; ok {
				return value
			}
			fmt.Println("'value' field missing in 'attributes'")
		} else {
			fmt.Println("'attributes' field missing or malformed")
		}
	} else {
		fmt.Println("'data' field missing or malformed")
	}
	return nil
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
		k := strings.TrimSuffix(ioType, "value") // exemple : "aiValue" --> ai

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

// Formats value to string based on type
func formatValue(v interface{}) string {
	switch val := v.(type) {
	case bool:
		if val {
			return "1"
		}
		return "0"
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", val)
	}
}
