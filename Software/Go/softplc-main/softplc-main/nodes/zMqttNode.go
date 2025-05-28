package nodes

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"time"
)

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
}

var mqttDescription = nodeDescription{
	AccordionName: "Communication",
	PrimaryType:   "LogicalNode",
	Type_:         "MqttNode",
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
	Output: []dataTypeNameStruct{{DataType: "bool", Name: "xDone"}, {DataType: "value", Name: "msg"}},
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func (n *MqttNode) messageHandler() mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
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
	RegisterNodeCreator("MqttNode", func() (Node, error) {
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

			topicToReceive := strings.Split(*n.input[3].Input, " ,, ")
			sub(n.client, topicToReceive)
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
		initConnection(n, "broker.hivemq.com", 1883, "go_mqtt_client")
		topicToReceive := strings.Split(*n.input[3].Input, " ,, ")
		sub(n.client, topicToReceive)

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
func initConnection(n *MqttNode, broker_ string, port_ int, ClientID string) {
	n.connectionIsInit = true
	var broker = "broker.hivemq.com"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	//opts.SetUsername("emqx")
	//opts.SetPassword("public")
	opts.SetDefaultPublishHandler(n.messageHandler())
	//opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	opts.AutoReconnect = true
	opts.ConnectRetry = true
	opts.ConnectRetryInterval = 2 * time.Second

	n.client = mqtt.NewClient(opts)
	if token := n.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
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
				token.Wait()
			}(topicTemp, message[i])
			/*token := client.Publish(topicTemp, 0, false, message[i])
			token.Wait()
			time.Sleep(time.Second)*/
		}
	}
}
