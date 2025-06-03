package nodes

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"time"
)

var clientID uint16 = 0 //Create unique ID for each connection (0..65535)

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
	clientID           uint16
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
	Output:            []dataTypeNameStruct{{DataType: "bool", Name: "xDone"}, {DataType: "value", Name: "msg"}},
	ParameterNameData: []string{"broker", "port", "user", "password"},
}

func (n *MqttNode) messageHandler() mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s to clientID %d\n", msg.Payload(), msg.Topic(), n.clientID)
		if !n.outputFlag {
			n.lastPayload = []string{}
		}
		n.lastPayload = append(n.lastPayload, fmt.Sprint(msg.Payload()))
		n.outputFlag = true
		if msg.Retained() {
			client.Publish(msg.Topic(), 0, true, "") //when is a message retain
		}

	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
	fmt.Println()
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
	if n.connectionIsInit {
		if n.input == nil {
			n.output[0].Output = "0"
			return
		}
		if *n.input[0].Input == "1" {
			topicToSend := strings.Split(*n.input[1].Input, " ,, ")
			msgToSend := strings.Split(*n.input[2].Input, " ,, ")
			publish(n.client, topicToSend, msgToSend)

			//topicToReceive := strings.Split(*n.input[3].Input, " ,, ")
			//sub(n.client, topicToReceive)
			//publish(client, "topic/test2", "Bonjour")
			//n.client.Disconnect(250)
		}
		//A message has been read
		if n.outputFlag {
			n.output[0].Output = "1"
			n.output[1].Output = strings.Join(n.lastPayload, " ,, ")
			n.outputFlag = false
			return
		} else {
			n.output[0].Output = "0"
		}
	} else {
		//initConnection(n, "broker.hivemq.com", 1883, "go_mqtt_client")
		initConnection(n)
		//topicToReceive := strings.Split(*n.input[3].Input, " ,, ")
		//sub(n.client, topicToReceive)

	}
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
func initConnection(n *MqttNode) { //, broker_ string, port_ int, ClientID string
	n.connectionIsInit = true
	var broker = n.parameterValueData[0] //"broker.hivemq.com"
	var port = n.parameterValueData[1]   //1883
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
	opts.OnConnectionLost = connectLostHandler

	opts.AutoReconnect = true
	opts.ConnectRetry = true
	opts.ConnectRetryInterval = 2 * time.Second

	n.client = mqtt.NewClient(opts)
	if token := n.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topicToReceive := strings.Split(*n.input[3].Input, " ,, ")
	sub(n.client, topicToReceive)
}
func sub(client mqtt.Client, topic []string) {
	for _, topicTemp := range topic {
		if topicTemp != "nothing to send" {
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
		if topicTemp != "nothing to send" {
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
		if n.input != nil && len(n.input) > 3 && n.input[3].Input != nil {
			topics := strings.Split(*n.input[3].Input, " ,, ")
			for _, topic := range topics {
				if topic != "nothing to send" {
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

}
