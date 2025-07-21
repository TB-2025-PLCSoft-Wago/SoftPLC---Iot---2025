package nodes

import (
	bus2 "SoftPLC/bus"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var serveurHttpIsInit bool = false

// receive msg
var (
	lastMessage = make(map[string]interface{})
	storage     = make(map[int]map[string]interface{})
	nextID      = 1
	mu          sync.Mutex
)

var (
	outputFlag     bool
	lastReceive    []string
	lastResourceId int

	//same for all server
	usernameServer string
	passwordServer string
	urlServer      string
)

var bus = bus2.NewEventBus()

type HttpServerNode struct {
	id                 int
	nodeType           string
	input              []InputHandle
	output             []OutputHandle
	parameterValueData []string
	client             *http.Client
	subBus             <-chan string
}

var httpServerDescription = nodeDescription{
	AccordionName: "Communication",
	PrimaryType:   "LogicalNode",
	Type_:         "ConfigurableNodeHttpServer",
	Display:       "HTTP Server Node",
	Label:         "HTTP Server",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "value", Name: "Parameters to receive"},
	},
	Output: []dataTypeNameStruct{
		{DataType: "bool", Name: "xDone"},
		{DataType: "value", Name: "Values received"},
		{DataType: "value", Name: "Resource ID"},
		{DataType: "value", Name: "Received URL path"},
	},
	ParameterNameData: []string{"url server", "user server", "password server"},
}

func init() {
	RegisterNodeCreator("ConfigurableNodeHttpServer", func() (Node, error) {
		return &HttpServerNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, httpServerDescription)

}

func (n *HttpServerNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
	n.parameterValueData = parameterValueData_
	n.client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	n.subBus = bus.Subscribe()
	/*
		go func() {
			for msg := range n.subBus {
				fmt.Println("subscriber ", &n, " has receive :", msg)
			}
		}()*/

}

// Flatten recursively flattens a nested JSON into flat key paths.
func flatten(data map[string]interface{}, prefix string, flat map[string]interface{}) {
	for key, value := range data {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "-" + key
		}
		switch v := value.(type) {
		case map[string]interface{}:
			flatten(v, fullKey, flat)
		default:
			flat[fullKey] = v
		}
	}
}

// http://localhost:8080/msgTest
func h1(w http.ResponseWriter, r *http.Request) {
	bus.Publish("Last path : " + r.URL.Path[1:])
}

func h2(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello from a HandleFunc #2!\n")
}
func h3(w http.ResponseWriter, r *http.Request) {
	// Check the method
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check the Authorization header
	auth := r.Header.Get("Authorization")
	if strings.Contains(auth, "Basic") {
		username, password, ok := r.BasicAuth()
		if !ok || username != usernameServer || password != passwordServer {
			http.Error(w, "Not allowed", http.StatusUnauthorized)
			return
		}
	} else if auth != "Bearer super-secret-token" {
		http.Error(w, "Not allowed", http.StatusUnauthorized)
		return
	}

	//Display the received headers (console)
	fmt.Println("=== Headers receive ===")
	for k, v := range r.Header {
		fmt.Printf("%s: %v\n", k, v)
	}

	// read body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Body reading error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var input map[string]interface{}
	if err := json.Unmarshal(body, &input); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}

	//storage message
	flat := make(map[string]interface{})
	flatten(input, "", flat)
	mu.Lock()
	lastMessage = flat
	mu.Unlock()
	bus.Publish("messagePut")

}
func flattenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Check the Authorization header
	auth := r.Header.Get("Authorization")
	if strings.Contains(auth, "Basic") {
		username, password, ok := r.BasicAuth()
		if !ok || username != usernameServer || password != passwordServer {
			http.Error(w, "Not allowed", http.StatusUnauthorized)
			return
		}
	}
	//read
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading the body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var input map[string]interface{}
	if err := json.Unmarshal(body, &input); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}

	//storage
	flat := make(map[string]interface{})
	flatten(input, "", flat)
	mu.Lock()
	id := nextID
	storage[id] = flat
	nextID++
	mu.Unlock()

	response := map[string]interface{}{
		"id":     id,
		"result": flat,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	lastResourceId = id
	bus.Publish("post")
}

// GET /flatten/{id}
func getOrDeleteFlattenedHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "URL invalid", http.StatusBadRequest)
		return
	}
	// Check the Authorization header
	auth := r.Header.Get("Authorization")
	if strings.Contains(auth, "Basic") {
		username, password, ok := r.BasicAuth()
		if !ok || username != usernameServer || password != passwordServer {
			http.Error(w, "Not allowed", http.StatusUnauthorized)
			return
		}
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalid", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		data, ok := storage[id]
		mu.Unlock()

		if !ok {
			http.Error(w, "No JSON with this ID", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)

	case http.MethodDelete:
		mu.Lock()
		_, ok := storage[id]
		if ok {
			delete(storage, id)
		}
		mu.Unlock()

		if !ok {
			http.Error(w, "Resource not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent) // 204 No Content
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// PATCH /flatten/{id}/{parameter result}
func patchFlattenedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Check the Authorization header
	auth := r.Header.Get("Authorization")
	if strings.Contains(auth, "Basic") {
		username, password, ok := r.BasicAuth()
		if !ok || username != usernameServer || password != passwordServer {
			http.Error(w, "Not allowed", http.StatusUnauthorized)
			return
		}
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 4 {
		http.Error(w, "URL invalid", http.StatusBadRequest)
		return
	}

	idStr := parts[3]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalid", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading the body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	if !outputFlag {
		lastReceive = []string{}
	}
	flat := make(map[string]interface{})
	flatten(data, "", flat)
	for k, v := range flat {
		mu.Lock()
		if storage[id][k] == nil {
			http.Error(w, "No JSON with this ID", http.StatusNotFound)
			mu.Unlock()
			return
		}
		storage[id][k] = v
		mu.Unlock()
		s := fmt.Sprintf("%v", v)
		lastReceive = append(lastReceive, s)
		lastResourceId = id
		outputFlag = true
		bus.Publish("patch")
	}
}
func (n *HttpServerNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = "0"
		return
	}
	if n.parameterValueData[0] == "" {
		urlServer = ":8080"
	} else {
		urlServer = n.parameterValueData[0]
	}
	usernameServer = n.parameterValueData[1]
	passwordServer = n.parameterValueData[2]
	if !serveurHttpIsInit {
		go func() {
			http.HandleFunc("/", h1)
			//http.HandleFunc("/endpoint", h2)
			http.HandleFunc("/message", h3)
			http.HandleFunc("/flatten", flattenHandler)
			http.HandleFunc("/flatten/", getOrDeleteFlattenedHandler)
			http.HandleFunc("/parameters/flatten/", patchFlattenedHandler) //id/...
			err := http.ListenAndServe(urlServer, nil)
			if err != nil {
				fmt.Println("Error server HTTP :", err)
			}

		}()
		serveurHttpIsInit = true
	}
	//select avoids lock
	select {
	case msg := <-n.subBus:
		//A message has been read
		//fmt.Println("subscriber received :", msg)
		if msg == "patch" || msg == "post" {
			n.output[2].Output = strconv.Itoa(lastResourceId)
			n.output[0].Output = "1"
			paramToSend := strings.Split(*n.input[0].Input, " ,, ")
			var temp []string
			for i := 0; i < len(paramToSend); i++ {
				value := fmt.Sprintf("%v", storage[lastResourceId][paramToSend[i]])
				temp = append(temp, value)

			}

			n.output[1].Output = strings.Join(temp, " ,, ") //n.lastReceive
			outputFlag = false
			return
		} else if msg == "messagePut" {
			n.output[2].Output = "-1"
			n.output[0].Output = "1"
			paramToSend := strings.Split(*n.input[0].Input, " ,, ")
			var temp []string
			for i := 0; i < len(paramToSend); i++ {
				value := fmt.Sprintf("%v", lastMessage[paramToSend[i]])
				if value == "<nil>" {
					temp = append(temp, "null")
				} else {
					temp = append(temp, value)
				}

			}

			n.output[1].Output = strings.Join(temp, " ,, ") //n.lastReceive

		} else if strings.Contains(msg, "Last path : ") {
			n.output[0].Output = "1"
			n.output[3].Output = strings.TrimPrefix(msg, "Last path : ")
		}
	default:
		n.output[0].Output = "0"
		n.output[1].Output = ""
		n.output[2].Output = ""
		n.output[3].Output = ""
		return
	}
}

func (n *HttpServerNode) GetNodeType() string {
	return n.nodeType
}

func (n *HttpServerNode) GetId() int {
	return n.id
}

func (n *HttpServerNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *HttpServerNode) GetInput() []InputHandle {
	return n.input
}

func (n *HttpServerNode) DestroyToBuildAgain() {
	n.client = nil
	outputFlag = false
	lastReceive = nil

	//Bus
	bus.Reset()
	bus = bus2.NewEventBus()
	n.subBus = bus.Subscribe()
}
