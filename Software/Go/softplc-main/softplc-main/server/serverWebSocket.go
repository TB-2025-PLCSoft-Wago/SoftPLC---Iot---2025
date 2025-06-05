package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type ApplianceUpdate struct {
	Type      string   `json:"type"` // always "update"
	Appliance string   `json:"appliance"`
	Outputs   []Output `json:"outputs"`
}
type Input struct {
	applianceName string
	id            int
	name          string
	defaultValue  string
}

/*
	type Output struct {
		defaultValue  string
		overrides     bool
	}
*/
type Output struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	ApplianceName string      `json:"applianceName"`
	Type          string      `json:"type"` // "bool", "string", "float"
	Value         interface{} `json:"value"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	clients   = make(map[*websocket.Conn]bool) // Connexions actives
	clientsMu sync.Mutex                       // Mutex pour accès concurrent
)

// Gère les connexions WebSocket
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// Ajout de la connexion dans la liste
	clientsMu.Lock()
	clients[conn] = true
	clientsMu.Unlock()

	fmt.Println("🟢 Client connecté")
	appliancesJSON := `{
    "type": "appliances",
    "appliances": [
			{
				"name": "TV",
				"buttons": [
					{ "text": "Power", "irCode": 123456 },
					{ "text": "Volume Up", "irCode": 789012 }
				]
			},
			{
				"name": "HiFi",
				"buttons": [
					{ "text": "Bass down", "irCode": 345678 }
				]
			}
		]
	}`
	conn.WriteMessage(websocket.TextMessage, []byte(appliancesJSON))

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)

			// Supprimer la connexion si elle est morte
			clientsMu.Lock()
			delete(clients, conn)
			clientsMu.Unlock()

			break
		}

		fmt.Println("Received:", string(msg))
		// Répond uniquement à celui qui a envoyé
		conn.WriteMessage(websocket.TextMessage, []byte("Echo: "+string(msg)))
		initOutput()
	}
}

// Lance le serveur WebSocket
func CreateWebSocket() {
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("🌐 WebSocket server on :8890")
	err := http.ListenAndServe(":8890", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

// Envoie un message à tous les clients connectés
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
func AddAppliance(name string, input []Input, output []Output) {
	//to do
}

func initOutput() {
	appliancesOutput := ApplianceUpdate{
		Type:      "update",
		Appliance: "TV",
		Outputs: []Output{
			{ID: 1, Name: "Power State", ApplianceName: "TV", Type: "bool", Value: true},
			{ID: 2, Name: "Current Channel", ApplianceName: "TV", Type: "string", Value: "Netflix"},
			{ID: 3, Name: "Volume", ApplianceName: "TV", Type: "float", Value: 17.5},
		},
	}

	// 🔄 Marshal en JSON
	jsonData, err := json.Marshal(appliancesOutput)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return
	}

	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		err := conn.WriteMessage(websocket.TextMessage, jsonData)
		if err != nil {
			fmt.Println("Send error:", err)
			conn.Close()
			delete(clients, conn)
		}
	}
}
