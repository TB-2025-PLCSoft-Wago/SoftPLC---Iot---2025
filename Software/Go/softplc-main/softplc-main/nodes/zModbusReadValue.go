package nodes

import (
	"fmt"
	"github.com/goburrow/modbus"
	"strings"
	"time"
)

type ModbusReadValueNode struct {
	id                 int
	nodeType           string
	input              []InputHandle
	output             []OutputHandle
	parameterValueData []string
	client             modbus.Client
	handler            *modbus.TCPClientHandler
	connectionIsInit   bool
	outputFlag         bool
	lastValuesReceived []string
	functionCode       int
}

var modbusReadValueDescription = nodeDescription{
	AccordionName: "Communication",
	PrimaryType:   "LogicalNode",
	Type_:         "ConfigurableNodeModbusReadValue",
	Display:       "Modbus Read Value (0x04)",
	Label:         "Modbus Read Value",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "bool", Name: "xEnable"},
		{DataType: "value", Name: "UnitID"},
		{DataType: "value", Name: "Addresses"},
		{DataType: "value", Name: "Quantity"},
	},
	Output: []dataTypeNameStruct{
		{DataType: "bool", Name: "xDone"},
		{DataType: "value", Name: "ValuesReceived"},
	},
	ParameterNameData: []string{"host", "port"},
}

func init() {
	RegisterNodeCreator("ConfigurableNodeModbusReadValue", func() (Node, error) {
		return &ModbusReadValueNode{
			id:       -1,
			nodeType: "",
		}, nil
	}, modbusReadValueDescription)
}

func (n *ModbusReadValueNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
	n.parameterValueData = parameterValueData_
	n.connectionIsInit = false
}

func (n *ModbusReadValueNode) initConnection(unitID byte) error {
	if n.connectionIsInit {
		return nil
	}

	//n.handler = modbus.NewTCPClientHandler(fmt.Sprintf("%s:%s", n.parameterValueData[0], n.parameterValueData[1]))
	address := fmt.Sprintf("%s:%s", n.parameterValueData[0], n.parameterValueData[1])
	handler := modbus.NewTCPClientHandler(address)
	handler.Timeout = 2 * time.Second
	handler.SlaveId = unitID

	err := handler.Connect()
	if err != nil {
		return err
	}

	n.handler = handler
	n.client = modbus.NewClient(n.handler)
	n.connectionIsInit = true
	return nil
}

func (n *ModbusReadValueNode) ProcessLogic() {
	go func() {
		if n.input == nil || len(n.input) < 4 {
			n.output[0].Output = "0"
			return
		}
		if n.input[0].Input == nil {
			n.output[0].Output = "0"
			return
		}

		if *n.input[0].Input != "1" {
			n.output[0].Output = "0"
			return
		}
		unitID := byte(0)
		if n.input[1].Input != nil {
			unitID = byte(atoiDefault(*n.input[1].Input, 0))
		}

		addresses := []string{"0"}
		if n.input[2].Input != nil {
			addresses = strings.Split(*n.input[2].Input, " ,, ")
		}
		quantities := []string{"1"}
		if n.input[3].Input != nil {
			quantities = strings.Split(*n.input[3].Input, " ,, ")
		}

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic in modbusReadValue:", r)
				n.connectionIsInit = false
				if n.handler != nil {
					n.handler.Close()
				}
				n.handler = nil
				n.client = nil
				n.output[0].Output = "0"
				n.output[1].Output = "error: panic"
			}
		}()

		if !n.connectionIsInit {
			if err := n.initConnection(unitID); err != nil {
				fmt.Println("ModbusReadValue connection error:", err)
				n.output[0].Output = "0"
				return
			}
		}
		if n.handler == nil || n.client == nil || !n.connectionIsInit {
			fmt.Println("Handler not ready")
			n.output[0].Output = "0"
			return
		}

		results := []string{}

		for i, addrStr := range addresses {
			address := uint16(atoiDefault(addrStr, 0))
			var quantity uint16
			if len(quantities) > i {
				quantity = uint16(atoiDefault(quantities[i], 1))
			} else {
				quantity = 1
			}
			var res []byte
			var err error

			res, err = n.client.ReadInputRegisters(address, quantity)

			if err != nil {
				fmt.Println("ModbusReadValue read error:", err)
				n.connectionIsInit = false
				if n.handler != nil {
					n.handler.Close()
				}
				n.handler = nil
				n.client = nil
				results = append(results, "error: "+err.Error())
				continue
			} else {
				val := fmt.Sprintf("%d", bytesToUint16(res))
				results = append(results, val)
			}
		}

		n.lastValuesReceived = results
		n.outputFlag = true
		n.output[0].Output = "1"
		n.output[1].Output = strings.Join(n.lastValuesReceived, " ,, ")
	}()
}

func (n *ModbusReadValueNode) GetNodeType() string { return n.nodeType }
func (n *ModbusReadValueNode) GetId() int          { return n.id }
func (n *ModbusReadValueNode) GetOutput(outName string) *OutputHandle {
	for i := range n.output {
		if n.output[i].Name == outName {
			return &n.output[i]
		}
	}
	return nil
}
func (n *ModbusReadValueNode) GetInput() []InputHandle { return n.input }

func (n *ModbusReadValueNode) DestroyToBuildAgain() {
	if n.handler != nil {
		n.handler.Close()
	}
	n.handler = nil
	n.client = nil
	n.connectionIsInit = false
	n.outputFlag = false
	n.lastValuesReceived = nil
}
