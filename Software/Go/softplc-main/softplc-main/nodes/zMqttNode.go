package nodes

import (
	"SoftPLC/server"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"time"
)

var clientID uint16 = 0 //CreateMqtt unique ID for each connection (0..65535)
var serveurIsInit bool = false

// MqttNode that wor like a logical
type MqttNode struct {
	id                 int
	nodeType           string
	input              []InputHandle
	output             []OutputHandle
	parameterValueData []string
	client             mqtt.Client
	connectionIsInit   bool
	msgHandler         mqtt.MessageHandler
	outputFlag         bool
	lastPayload        []string
	lastTopic          []string
	clientID           uint16
	topicToReceive     []string
	topicToReceiveSave []string
}

var mqttDescription = nodeDescription{
	AccordionName: "Communication",
	PrimaryType:   "LogicalNode",
	Type_:         "ConfigurableNode",
	Display:       "mqtt Node",
	Label:         "MQTT",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "bool", Name: "xEnable"},
		{DataType: "value", Name: "topicToSend"},
		{DataType: "value", Name: "msgToSend"},
		{DataType: "value", Name: "topicToReceive"},
	},
	Output:            []dataTypeNameStruct{{DataType: "bool", Name: "xReceive"}, {DataType: "value", Name: "msgLastReceived"}},
	ParameterNameData: []string{"broker", "port", "user", "password"},
}

func (n *MqttNode) messageHandler() mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s to clientID %d\n", msg.Payload(), msg.Topic(), n.clientID)
		if !n.outputFlag {
			n.lastPayload = []string{}
			n.lastTopic = []string{}
		}
		n.lastPayload = append(n.lastPayload, string(msg.Payload()))
		n.lastTopic = append(n.lastTopic, string(msg.Topic()))
		n.outputFlag = true
		if msg.Retained() {
			client.Publish(msg.Topic(), 0, true, "") //when is a message retain
		}

	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

func makeConnectLostHandler(n *MqttNode) mqtt.ConnectionLostHandler {
	return func(client mqtt.Client, err error) {
		fmt.Printf("Connect lost for client %d: %v\n", n.clientID, err)
		time.Sleep(2 * time.Second)
		if n == nil {
			fmt.Println("Reconnect handler: MqttNode is nil, aborting reconnect")
			return
		}
		if n.client == nil {
			fmt.Println("Reconnect handler: client is nil, aborting reconnect")
			return
		}
		// Manual reconnection attempt if AutoReconnect is not sufficient
		if token := n.client.Connect(); token.Wait() && token.Error() != nil {
			fmt.Printf("Reconnection failed for client %d: %v\n", n.clientID, token.Error())
		} else {
			fmt.Printf("Reconnected successfully for client %d\n", n.clientID)
			n.connectionIsInit = false
		}
	}
}

func init() {
	RegisterNodeCreator("ConfigurableNode", func() (Node, error) {
		return &MqttNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, mqttDescription)
}

func (n *MqttNode) ProcessLogic() {
	go func() {
		if n.connectionIsInit {
			if n.input == nil {
				n.output[0].Output = "0"
				n.output[1].Output = ""
				return
			}
			if n.input[0].Input == nil {
				n.output[0].Output = "0"
				n.output[1].Output = ""
				return
			}
			if *n.input[0].Input == "1" {
				var topicToSend, msgToSend []string
				if n.input[1].Input != nil {
					topicToSend = strings.Split(*n.input[1].Input, " ,, ")
				} else {
					topicToSend = []string{}
				}
				if n.input[2].Input != nil {
					msgToSend = strings.Split(*n.input[2].Input, " ,, ")
				} else {
					msgToSend = []string{}
				}

				publish(n.client, topicToSend, msgToSend)
			}
			var topicToReceive []string
			if n.input[3].Input != nil {
				topicToReceive = strings.Split(*n.input[3].Input, " ,, ")
			} else {
				topicToReceive = make([]string, 0)
			}
			//A message has been read
			if n.outputFlag {
				n.output[0].Output = "1"

				var newMsg []string
			OuterLoop:
				for _, topicTempToReceive := range topicToReceive {
					for i, topicTemp := range n.lastTopic {
						if topicTempToReceive == topicTemp {
							newMsg = append(newMsg, n.lastPayload[i])
							continue OuterLoop
						}
					}
					newMsg = append(newMsg, "")
					//token := n.client.Unsubscribe(topicTemp)
					//token.Wait()
				}
				n.output[1].Output = strings.Join(newMsg, " ,, ")
				//n.output[1].Output = strings.Join(n.lastPayload, " ,, ")
				n.outputFlag = false
			} else {
				n.output[0].Output = "0"
			}

			//Handle : Subscibe and Unsubscribe
			flagSave := false
		OuterLoop2:
			for _, tempReceive := range topicToReceive {
				for _, tempOldReceive := range n.topicToReceiveSave {
					if tempOldReceive == tempReceive {
						continue OuterLoop2
					}
				}
				if tempReceive != "null" && tempReceive != "empty" {
					if n.connectionIsInit && n.client != nil && n.client.IsConnected() {
						sub(n.client, []string{tempReceive})
						flagSave = true
					}
				}
			}

		OuterLoop3:
			for _, tempOldReceive := range n.topicToReceiveSave {
				for _, tempReceive := range topicToReceive {
					if tempOldReceive == tempReceive {
						continue OuterLoop3
					}
				}
				if tempOldReceive != "null" && tempOldReceive != "empty" {
					if n.connectionIsInit && n.client != nil && n.client.IsConnected() {
						token := n.client.Unsubscribe(tempOldReceive)
						token.Wait()
						flagSave = true
					}
				}
			}

			if flagSave {
				n.topicToReceiveSave = topicToReceive
			}
		} else {
			//initConnection(n, "broker.hivemq.com", 1883, "go_mqtt_client")
			initConnection(n)
			n.output[1].Output = ""

		}
	}()
}

func (n *MqttNode) GetNodeType() string {
	return n.nodeType
}

func (n *MqttNode) GetId() int {
	return n.id
}

func (n *MqttNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *MqttNode) GetInput() []InputHandle {
	return n.input
}

func (n *MqttNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
	n.parameterValueData = parameterValueData_

	n.connectionIsInit = false
}
func initConnection(n *MqttNode) {
	if !serveurIsInit {
		if n.parameterValueData[0] != "" {
			go server.CreateMqtt()
		}

		serveurIsInit = true
	}

	n.connectionIsInit = true
	var broker = n.parameterValueData[0] //"broker.hivemq.com"
	var port = n.parameterValueData[1]   //1883
	if port == "" {
		port = "1883"
	}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", broker, port))
	//opts.SetClientID((fmt.Sprintf("CC100_mqtt_client_%d", n.id)))
	opts.SetClientID((fmt.Sprintf("CC100_mqtt_c%d", clientID)))
	n.clientID = clientID
	clientID = clientID + 1
	opts.SetUsername(n.parameterValueData[2]) //emqx
	opts.SetPassword(n.parameterValueData[3]) //public
	opts.SetDefaultPublishHandler(n.messageHandler())
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = makeConnectLostHandler(n)

	opts.AutoReconnect = true
	opts.ConnectRetry = true
	opts.ConnectRetryInterval = 2 * time.Second

	n.client = mqtt.NewClient(opts)
	if token := n.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if n.input[3].Input != nil {
		n.topicToReceive = strings.Split(*n.input[3].Input, " ,, ")
	} else {
		n.topicToReceive = make([]string, 0)
	}

	n.topicToReceiveSave = n.topicToReceive
	sub(n.client, n.topicToReceive)
}
func sub(client mqtt.Client, topic []string) {
	for _, topicTemp := range topic {
		if topicTemp != "null" && topicTemp != "empty" {
			token := client.Subscribe(topicTemp, 1, nil)
			token.Wait()
			//fmt.Printf("Subscribed to topic %s", topicTemp)
			//fmt.Println()
		}
	}

	//fmt.Printf("mqtt id %s", topic)
}

func publish(client mqtt.Client, topic []string, message []string) {
	for i, topicTemp := range topic {
		if topicTemp != "null" && topicTemp != "empty" {
			go func(topicStr, msgStr string) {
				token := client.Publish(topicStr, 0, false, msgStr)
				fmt.Printf("publish message: %s to topic: %s \n", msgStr, topicStr)
				token.Wait()
			}(topicTemp, message[i])
		}
	}
}

func (n *MqttNode) DestroyToBuildAgain() {
	// Check if a connection is initialized
	if n.connectionIsInit && n.client != nil && n.client.IsConnected() {
		// Attempt to unsubscribe from all known topics
		if n.topicToReceive != nil {
			for _, topic := range n.topicToReceive {
				if topic != "null" && topic != "" {
					token := n.client.Unsubscribe(topic)
					token.Wait()
				}
			}
		}

		// Disconnecting the client
		n.client.Disconnect(250)
	}

	// Cleaning of internal fields
	n.client = nil
	n.connectionIsInit = false
	n.outputFlag = false
	n.lastPayload = nil
	n.lastTopic = nil
	n.topicToReceive = nil

}
