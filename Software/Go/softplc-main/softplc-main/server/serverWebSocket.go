package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

type Appliance struct {
	Name   string  `json:"name"`
	Inputs []Input `json:"inputs"`
}

type AppliancesData struct {
	Type       string      `json:"type"`
	Appliances []Appliance `json:"appliances"`
}

var appliancesJSON = AppliancesData{
	Type:       "appliances",
	Appliances: []Appliance{},
}

type InputState struct {
	IRCode int
	Value  string
}

var irCode int
var InputsStateWeb []InputState

type Input struct {
	Text      string `json:"text"`
	IRCode    int    `json:"irCode"`
	TextInput bool   `json:"textInput,omitempty"`
}

/*
	type Output struct {
		defaultValue  string
		overrides     bool
	}
*/
type ApplianceUpdate struct {
	Type      string   `json:"type"` // always "update"
	Appliance string   `json:"appliance"`
	Outputs   []Output `json:"outputs"`
}

var AllOutputsByAppliance = map[string][]Output{}

type OutputState struct {
	ID    int
	Value interface{}
}

var OutputsStateWeb []OutputState

type Output struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	ApplianceName string      `json:"applianceName"`
	Type          string      `json:"type"` // "bool", "string", "float"
	Value         interface{} `json:"value"`
}

var outputChange bool = false

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	clients   = make(map[*websocket.Conn]bool) // active connections
	clientsMu sync.Mutex
)

// Manages WebSocket connections
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// Adding the connection in the list
	clientsMu.Lock()
	clients[conn] = true
	clientsMu.Unlock()

	fmt.Println("🟢 client connected")

	/*
			appliancesJSON := `{
		    "type": "appliances",
		    "appliances": [
					{
						"name": "TV",
						"inputs": [
							{ "text": "Power", "irCode": 123456 },
							{ "text": "Volume Up", "irCode": 789012 }
						]
					},
					{
						"name": "HiFi",
						"inputs": [
							{ "text": "Bass down", "irCode": 345678 }
						]
					},
					{
						"name": "Rien",
						"inputs": [
							{ "text": "Rien en txt", "irCode": 12431122 },
							{ "text": "Volume", "irCode": 789012, "textInput": true }
						]
					}
				]
			}`*/
	jsonBytes, err := json.Marshal(appliancesJSON)
	if err != nil {
		log.Println("Error server webSocket Marshal JSON :", err)
		return
	}
	conn.WriteMessage(websocket.TextMessage, jsonBytes)
	initOutput()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)

			// Delete the connection if it is dead
			clientsMu.Lock()
			delete(clients, conn)
			clientsMu.Unlock()

			break
		}

		//fmt.Println("Received:", string(msg))
		handleIncomingMessage(conn, msg)

		//conn.WriteMessage(websocket.TextMessage, []byte("Echo: "+string(msg))) // Reply only to the one who sent
		//initOutput()
	}
}

// start server WebSocket
func CreateWebSocket() {
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("🌐 WebSocket server on :8890")
	err := http.ListenAndServe(":8890", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

// Sends a message to all connected clients
func SendToWebSocket(msg string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			fmt.Println("Send error:", err)
			conn.Close()
			delete(clients, conn)
		}
	}
}

func addToState(irCode_ int, typeOf string) {
	var inputOrOutputToAdd InputState
	switch typeOf {
	case "number":
		inputOrOutputToAdd.Value = "0"
	case "string":
		inputOrOutputToAdd.Value = ""
	case "bool":
		inputOrOutputToAdd.Value = "0"
	default:
		fmt.Println("Unrecognized type of typeOf : ", typeOf)
		return
	}
	inputOrOutputToAdd.IRCode = irCode_
	InputsStateWeb = append(InputsStateWeb, inputOrOutputToAdd)
}
func initOutput() {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	// For each device in AllOutputsByAppliance
	for applianceName, outputs := range AllOutputsByAppliance {
		// Update the output values from OutputsStateWeb
		for i := range outputs {
			for _, state := range OutputsStateWeb {
				if outputs[i].ID == state.ID {
					switch outputs[i].Type {
					case "number":
						//output.Value, _ = strconv.ParseFloat(newValue, 64)
					case "value":
						outputs[i].Value = state.Value
					case "bool":
						if state.Value == "1" {
							outputs[i].Value = true
						} else {
							outputs[i].Value = false
						}
					default:
						fmt.Println("Unrecognized type of typeOf : ", outputs[i].Type)
						break
					}
					break
				}
			}
		}

		// Build and send ApplianceUpdate with all outputs
		applianceUpdate := ApplianceUpdate{
			Type:      "update",
			Appliance: applianceName,
			Outputs:   outputs,
		}

		jsonData, err := json.Marshal(applianceUpdate)
		if err != nil {
			log.Println("❌ JSON marshal error in initOutput:", err)
			continue
		}

		for conn := range clients {
			err = conn.WriteMessage(websocket.TextMessage, jsonData)
			if err != nil {
				fmt.Println("❌ Send error:", err)
				conn.Close()
				delete(clients, conn)
			}
		}
	}
}

/*
func initOutput() {
	appliances := []ApplianceUpdate{
		{
			Type:      "update",
			Appliance: "TV",
			Outputs: []Output{
				{ID: 1, Name: "Power State", ApplianceName: "TV", Type: "bool", Value: true},
				{ID: 2, Name: "Current Channel", ApplianceName: "TV", Type: "string", Value: "Netflix"},
				{ID: 3, Name: "Volume", ApplianceName: "TV", Type: "float", Value: 17.5},
			},
		},
		{
			Type:      "update",
			Appliance: "HiFi",
			Outputs: []Output{
				{ID: 1, Name: "Power State", ApplianceName: "HiFi", Type: "bool", Value: false},
				{ID: 2, Name: "Bass Level", ApplianceName: "HiFi", Type: "float", Value: 3.2},
				{ID: 3, Name: "Source", ApplianceName: "HiFi", Type: "string", Value: "Bluetooth"},
			},
		},

		{
			Type:      "update",
			Appliance: "home",
			Outputs: []Output{
				{ID: 1, Name: "Power State of home", ApplianceName: "home", Type: "bool", Value: false},
				{ID: 2, Name: "home Level", ApplianceName: "home", Type: "float", Value: 3.2},
				{ID: 3, Name: "Source home", ApplianceName: "home", Type: "string", Value: "Bluetooth"},
			},
		},
	}

	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		for _, appliance := range appliances {
			jsonData, err := json.Marshal(appliance)
			if err != nil {
				fmt.Println("JSON marshal error:", err)
				continue
			}
			err = conn.WriteMessage(websocket.TextMessage, jsonData)
			if err != nil {
				fmt.Println("Send error:", err)
				conn.Close()
				delete(clients, conn)
			}
		}
	}
}*/

func handleIncomingMessage(conn *websocket.Conn, msg []byte) {

	var data map[string]interface{}
	if err := json.Unmarshal(msg, &data); err != nil {
		// is not a JSON
		fmt.Println("Received (raw):", string(msg))
		conn.WriteMessage(websocket.TextMessage, []byte("Echo: "+string(msg)))
		return
	}

	// identify msg
	switch {
	case data["type"] == "irCommand" && data["irCode"] != nil:
		//toggleInput
		for i := range InputsStateWeb {
			if irCodeFloat, ok := data["irCode"].(float64); ok {
				if InputsStateWeb[i].IRCode == int(irCodeFloat) {
					if InputsStateWeb[i].Value == "0" {
						InputsStateWeb[i].Value = "1"
					} else {
						InputsStateWeb[i].Value = "0"
					}
				}
			}

		}

		//conn.WriteMessage(websocket.TextMessage, []byte("IR order received"))

	case data["irCode"] != nil && data["value"] != nil:
		fmt.Println("✍️ User input value:", data)
		for i := range InputsStateWeb {
			if irCodeFloat, ok := data["irCode"].(float64); ok {
				if InputsStateWeb[i].IRCode == int(irCodeFloat) {

					if valueStr, ok := data["value"].(string); ok {
						InputsStateWeb[i].Value = valueStr
					} else {
						fmt.Println("❌ data[\"value\"]")
					}

				}
			}

		}
		//conn.WriteMessage(websocket.TextMessage, []byte("value saved"))

	case data["type"] == "update":
		fmt.Println("🔄 update state (output):", data)
	case data["type"] == "edge_clicked":
		fmt.Println("🖱️ edge clicked:", data)

		if data["tool"] == "DisplayConnectionDebug" {
			fmt.Println("🖱️ edge clicked DisplayConnectionDebug")
			//delete if exist else append
			source, ok1 := data["source"].(string)
			sourceHandle, ok2 := data["sourceHandle"].(string)
			if ok1 && ok2 {
				for _, toDebugItem := range toDebugList {
					if source == toDebugItem.id && sourceHandle == toDebugItem.sourceHandle {
						removeFromDebugList(source, sourceHandle)
						return
					}
				}

				toDebugList = append(toDebugList, debugList{
					id:           source,
					sourceHandle: sourceHandle,
				})
			} else {
				fmt.Println("Failed source or sourceHandle to string")
			}
		}

	default:
		fmt.Println("📦 Unrecognized JSON message:", data)
		//conn.WriteMessage(websocket.TextMessage, []byte("Format not recognized"))
	}
}

func AddInputToAppliance(applianceName string, nameSignal string, typeOf string) string {
	irCode++
	var textInput bool
	switch typeOf {
	case "number":
		textInput = true
	case "string":
		textInput = true
	case "bool":
		textInput = false
	default:
		fmt.Println("Unrecognized type of input : ", nameSignal)

	}
	addToState(irCode, typeOf)
	// look for appliance
	for i, appliance := range appliancesJSON.Appliances {
		if appliance.Name == applianceName {
			// add input
			input := Input{
				Text:   nameSignal,
				IRCode: irCode,
			}
			if textInput {
				input.TextInput = true
			}
			appliancesJSON.Appliances[i].Inputs = append(appliancesJSON.Appliances[i].Inputs, input)
			return strconv.Itoa(irCode)
		}
	}

	// create Appliance if not find
	input := Input{
		Text:   nameSignal,
		IRCode: irCode,
	}
	if textInput {
		input.TextInput = true
	}
	newAppliance := Appliance{
		Name:   applianceName,
		Inputs: []Input{input},
	}
	appliancesJSON.Appliances = append(appliancesJSON.Appliances, newAppliance)
	return strconv.Itoa(irCode)
}

func AddOutputToAppliance(applianceName string, outputName string, typeOf string, value interface{}) string {
	newID := len(OutputsStateWeb) + 1

	output := Output{
		ID:            newID,
		Name:          outputName,
		ApplianceName: applianceName,
		Type:          typeOf,
		Value:         value,
	}

	// Add to OutputsStateWeb
	OutputsStateWeb = append(OutputsStateWeb, OutputState{ID: newID, Value: value})

	// Add to the map
	AllOutputsByAppliance[applianceName] = append(AllOutputsByAppliance[applianceName], output)

	// send WebSocket
	applianceUpdate := ApplianceUpdate{
		Type:      "update",
		Appliance: applianceName,
		Outputs:   []Output{output},
	}
	jsonData, err := json.Marshal(applianceUpdate)
	if err == nil {
		SendToWebSocket(string(jsonData))
	}
	return strconv.Itoa(newID)
}

// Utility to retrieve all outputs
func initAllOutputs() []ApplianceUpdate {
	var updates []ApplianceUpdate
	for applianceName, outputs := range AllOutputsByAppliance {
		updates = append(updates, ApplianceUpdate{
			Type:      "update",
			Appliance: applianceName,
			Outputs:   outputs,
		})
	}
	return updates
}

func UpdateOutputValueByID(outputID int, newValue interface{}) {
	//Update OutputsStateWeb
	for i := range OutputsStateWeb {
		if OutputsStateWeb[i].ID == outputID {
			OutputsStateWeb[i].Value = newValue
			outputChange = true
			return
		}
	}
	log.Printf("⚠️ No OutputState with ID %d found\n", outputID)
}

/*
func UpdateOutputValueByID(outputID int, newValue interface{}) {
	//Update OutputsStateWeb
	for i := range OutputsStateWeb {
		if OutputsStateWeb[i].ID == outputID {
			OutputsStateWeb[i].Value = newValue

			// Find the appliance from the outputID
			for applianceName, outputs := range AllOutputsByAppliance {
				for j := range outputs {
					if outputs[j].ID == outputID {
						/*
							// Update the value in the local copy
							switch outputs[j].Type {
							case "number":
								//output.Value, _ = strconv.ParseFloat(newValue, 64)
							case "value":
								outputs[j].Value = newValue
							case "bool":
								if newValue == "1" {
									outputs[j].Value = true
								} else {
									outputs[j].Value = false
								}
							default:
								fmt.Println("Unrecognized type of typeOf : ", outputs[j].Type)
								continue
							}
*/
/*
						// Update AllOutputsByAppliance
						AllOutputsByAppliance[applianceName][j] = outputs[j]

						// Update the values of all outputs of this device
						for k := range AllOutputsByAppliance[applianceName] {
							for _, state := range OutputsStateWeb {
								if AllOutputsByAppliance[applianceName][k].ID == state.ID {
									switch AllOutputsByAppliance[applianceName][k].Type {
									case "number":
										//output.Value, _ = strconv.ParseFloat(newValue, 64)
									case "value":
										AllOutputsByAppliance[applianceName][k].Value = state.Value
									case "bool":
										if state.Value == "1" {
											AllOutputsByAppliance[applianceName][k].Value = true
										} else {
											AllOutputsByAppliance[applianceName][k].Value = false
										}
									default:
										fmt.Println("Unrecognized type of typeOf : ", AllOutputsByAppliance[applianceName][k].Type)
										continue
									}
								}
							}
						}

						// Send the entire output array for this device
						applianceUpdate := ApplianceUpdate{
							Type:      "update",
							Appliance: applianceName,
							Outputs:   AllOutputsByAppliance[applianceName],
						}

						jsonData, err := json.Marshal(applianceUpdate)
						if err != nil {
							log.Println("❌ JSON marshal error in UpdateOutputValueByID:", err)
							return
						}

						clientsMu.Lock()
						defer clientsMu.Unlock()
						for conn := range clients {
							err = conn.WriteMessage(websocket.TextMessage, jsonData)
							if err != nil {
								fmt.Println("Send error:", err)
								conn.Close()
								delete(clients, conn)
							}
						}
						return
					}
				}
			}

			log.Printf("⚠️ Complete output not found for ID %d\n", outputID)
			return
		}
	}

	log.Printf("⚠️ No OutputState with ID %d found\n", outputID)
}*/

func UpdateOutputValueView() {
	if outputChange {
		outputChange = false
		// Find the appliance from the outputID
		for applianceName, outputs := range AllOutputsByAppliance {
			for j := range outputs {
				// Update AllOutputsByAppliance
				AllOutputsByAppliance[applianceName][j] = outputs[j]

				// Update the values of all outputs of this device
				for k := range AllOutputsByAppliance[applianceName] {
					for _, state := range OutputsStateWeb {
						if AllOutputsByAppliance[applianceName][k].ID == state.ID {
							switch AllOutputsByAppliance[applianceName][k].Type {
							case "number":
								//output.Value, _ = strconv.ParseFloat(newValue, 64)
							case "value":
								AllOutputsByAppliance[applianceName][k].Value = state.Value
							case "bool":
								if state.Value == "1" {
									AllOutputsByAppliance[applianceName][k].Value = true
								} else {
									AllOutputsByAppliance[applianceName][k].Value = false
								}
							default:
								fmt.Println("Unrecognized type of typeOf : ", AllOutputsByAppliance[applianceName][k].Type)
								continue
							}
						}
					}
				}

				// Send the entire output array for this device
				applianceUpdate := ApplianceUpdate{
					Type:      "update",
					Appliance: applianceName,
					Outputs:   AllOutputsByAppliance[applianceName],
				}

				jsonData, err := json.Marshal(applianceUpdate)
				if err != nil {
					log.Println("❌ JSON marshal error in UpdateOutputValueByID:", err)
					return
				}

				clientsMu.Lock()
				for conn := range clients {
					err = conn.WriteMessage(websocket.TextMessage, jsonData)
					if err != nil {
						fmt.Println("Send error:", err)
						conn.Close()
						delete(clients, conn)
					}
				}
				clientsMu.Unlock()

			}
		}
	}

}

func ResetAll() {
	irCode = 0
	InputsStateWeb = []InputState{}
	appliancesJSON = AppliancesData{
		Type:       "appliances",
		Appliances: []Appliance{},
	}

	OutputsStateWeb = []OutputState{}
	AllOutputsByAppliance = map[string][]Output{}
	LogicalsStateWeb = []DebugState{}
}

/** Debug Mode **/
var IsActiveDebug bool = false
var DebugGraphJson interface{}
var DebugGraphJsonCopy interface{}

type DebugState struct {
	ID           string
	Value        *string
	SourceHandle string
	Type_        string
}
type debugList struct {
	id           string
	sourceHandle string
}

var test bool
var LogicalsStateWeb []DebugState
var toDebugList []debugList // the list of all edges that were added

func AddDebugState(id string, valuePtr *string, sourceHandle string, type_ string) {
	if valuePtr != nil {
		LogicalsStateWeb = append(LogicalsStateWeb, DebugState{
			ID:           id,
			Value:        valuePtr,
			SourceHandle: sourceHandle,
			Type_:        type_,
		})
	}
}

func RemoveLabelsFromDebugGraph() {
	graphMap, ok := DebugGraphJson.(map[string]interface{})
	if !ok {
		return
	}

	edges, ok := graphMap["edges"].([]interface{})
	if !ok {
		return
	}

	for _, edge := range edges {
		if edgeMap, ok := edge.(map[string]interface{}); ok {
			if dataMap, ok := edgeMap["data"].(map[string]interface{}); ok {
				delete(dataMap, "label")
			}
			edgeMap["type"] = "step"
			delete(edgeMap, "label")

			styleMap, ok := edgeMap["style"].(map[string]interface{})
			if !ok {
				styleMap = make(map[string]interface{})
				edgeMap["style"] = styleMap
			}

			delete(styleMap, "strokeDasharray")
			delete(styleMap, "animation")
		}
	}
}

func DebugMode() {

	// Cast to map
	graphMap := DebugGraphJson.(map[string]interface{})
	edges := graphMap["edges"].([]interface{})

	// go through the edges
	for _, edge := range edges {
		edgeMap := edge.(map[string]interface{})

		sourceVal, ok1 := edgeMap["source"]
		sourceHandleVal, ok2 := edgeMap["sourceHandle"]
		sourcetypeVal, ok3 := edgeMap["type"]

		if !ok1 || sourceVal == nil {
			fmt.Println("source is missing or nil in edge:", edgeMap)
			continue
		}
		if !ok2 || sourceHandleVal == nil {
			fmt.Println("sourceHandle is missing or nil in edge:", edgeMap)
			continue
		}
		if !ok3 || sourcetypeVal == nil {
			fmt.Println("type is missing or nil in edge:", edgeMap)
			continue
		}

		sourceID, ok1 := sourceVal.(string)
		sourceHandle, ok2 := sourceHandleVal.(string)
		_, ok3 = sourceHandleVal.(string)
		if !ok1 || !ok2 || !ok3 {
			fmt.Println("source or sourceHandle or sourcetype is not a string in edge:", edgeMap)
			continue
		}

		isInDebugList := false

		// go through the states
		for _, state := range LogicalsStateWeb {
			if state.ID == sourceID && (state.SourceHandle == sourceHandle || state.SourceHandle == "") {
				/* show when in Debug list */
				/*for _, toDebugItem := range toDebugList {
						if state.ID == toDebugItem.id && state.SourceHandle == toDebugItem.sourceHandle {
							isInDebugList = true
							edgeMap["type"] = "customDebugEdge"
							// "data"
							dataMap, ok := edgeMap["data"].(map[string]interface{})
							if !ok {
								dataMap = make(map[string]interface{})
								edgeMap["data"] = dataMap
							}

							dataMap["label"] = *state.Value

							//"style"
							styleMap, ok := edgeMap["style"].(map[string]interface{})
							if !ok {
								styleMap = make(map[string]interface{})
								edgeMap["style"] = styleMap
							}
							if state.Type_ == "bool" {
								if *state.Value == "1" {
									styleMap["stroke"] = "#00FF00"
								} else {
									styleMap["stroke"] = "#FF0000"
								}
							}

							styleMap["strokeDasharray"] = "8 4"
							styleMap["animation"] = "dash 1s linear infinite"
							break
						}
					}
				}

				if isInDebugList {
					break
				} else {
					// delete old label
					if dataMap, ok := edgeMap["data"].(map[string]interface{}); ok {
						delete(dataMap, "label")
					}

					// clean of the style
					styleMap, ok := edgeMap["style"].(map[string]interface{})
					if !ok {
						styleMap = make(map[string]interface{})
						edgeMap["style"] = styleMap
					}
					delete(styleMap, "strokeDasharray")
					delete(styleMap, "animation")
				}*/

				/* show when NOT in Debug list */
				isInDebugList = false
				for _, item := range toDebugList {
					if state.ID == item.id && state.SourceHandle == item.sourceHandle {
						isInDebugList = true
					}
				}
				if !isInDebugList {
					// "data"
					dataMap, ok := edgeMap["data"].(map[string]interface{})
					if !ok {
						dataMap = make(map[string]interface{})
						edgeMap["data"] = dataMap
					}
					dataMap["label"] = *state.Value

					//"style"
					styleMap, ok := edgeMap["style"].(map[string]interface{})
					if !ok {
						styleMap = make(map[string]interface{})
						edgeMap["style"] = styleMap
					}
					if state.Type_ == "bool" {
						if *state.Value == "1" {
							styleMap["stroke"] = "#00FF00"
						} else {
							styleMap["stroke"] = "#FF0000"
						}
					}

					styleMap["strokeDasharray"] = "8 4"
					styleMap["animation"] = "dash 1s linear infinite"
				} else {
					// delete old label
					if dataMap, ok := edgeMap["data"].(map[string]interface{}); ok {
						delete(dataMap, "label")
					}
					// clean of the style
					styleMap, ok := edgeMap["style"].(map[string]interface{})
					if !ok {
						styleMap = make(map[string]interface{})
						edgeMap["style"] = styleMap
					}
					delete(styleMap, "strokeDasharray")
					delete(styleMap, "animation")
				}
			}
		}
	}

	// Serialize
	modifiedJSON, err := json.MarshalIndent(graphMap, "", "  ")
	if err == nil {
		SendToWebSocket(string(modifiedJSON))
	}
}

func removeFromDebugList(id, sourceHandle string) {
	filtered := make([]debugList, 0, len(toDebugList))
	for _, item := range toDebugList {
		if !(item.id == id && item.sourceHandle == sourceHandle) {
			filtered = append(filtered, item)
		}
	}
	toDebugList = filtered
}

// Deep copy of an interface{}
func CopyInterface(src interface{}) interface{} {
	bytes, err := json.Marshal(src)
	if err != nil {
		log.Println("error during marshal :", err)
		return nil
	}

	var dst interface{}
	err = json.Unmarshal(bytes, &dst)
	if err != nil {
		log.Println("error during unmarshal :", err)
		return nil
	}

	return dst
}
