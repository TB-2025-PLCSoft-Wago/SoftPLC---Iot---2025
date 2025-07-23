package nodes

import (
	"fmt"
	"github.com/goburrow/modbus"
	"strings"
	"time"
)

type ModbusReadBoolNode struct {
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
	unitIDRetain       byte
}

var modbusDescription = nodeDescription{
	AccordionName: "Communication",
	PrimaryType:   "LogicalNode",
	Type_:         "ConfigurableNodeModbusReadBool",
	Display:       "Modbus Read Bool (0x02)",
	Label:         "Modbus Read Bool",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "bool", Name: "xEnable"},
		{DataType: "value", Name: "UnitID"},
		//{DataType: "value", Name: "FunctionCode"},
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
	RegisterNodeCreator("ConfigurableNodeModbusReadBool", func() (Node, error) {
		return &ModbusReadBoolNode{
			id:       -1,
			nodeType: "",
		}, nil
	}, modbusDescription)
}

func (n *ModbusReadBoolNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
	n.parameterValueData = parameterValueData_
	n.connectionIsInit = false
}

func (n *ModbusReadBoolNode) initConnection(unitID byte) error {
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

func (n *ModbusReadBoolNode) ProcessLogic() {
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

		n.functionCode = 2
		/*	if n.input[2].Input != nil {
			n.functionCode = atoiDefault(*n.input[2].Input, 2)
		}*/
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
				fmt.Println("Recovered from panic in ModbusReadBool:", r)
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
				fmt.Println("ModbusReadBool connection error:", err)
				n.output[0].Output = "0"
				n.output[1].Output = "connection error"
				return
			}
			n.unitIDRetain = unitID
		}
		if n.handler == nil || n.client == nil || !n.connectionIsInit {
			fmt.Println("Handler not ready")
			n.output[0].Output = "0"
			n.output[1].Output = "Handler not ready"
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
			bits := make([]bool, quantity)

			switch n.functionCode {
			case 1:
				res, err = n.client.ReadCoils(address, quantity)
			case 2:
				res, err = n.client.ReadDiscreteInputs(address, quantity*8)
				//Convert to Bool
				for i2 := 0; uint16(i2) < quantity; i2++ {
					byteIndex := i2 / 8
					bitIndex := uint(i2 % 8)
					if len(res) > byteIndex {
						bits[i2] = (res[byteIndex] & (1 << bitIndex)) != 0
					} else {
						bits[i2] = false
					}

				}
				for _, bit := range bits {
					if bit {
						results = append(results, "1")
					} else {
						results = append(results, "0")
					}
				}
			case 3:
				res, err = n.client.ReadHoldingRegisters(address, quantity)

			case 4:
				res, err = n.client.ReadInputRegisters(address, quantity)
			default:
				res = []byte{}
				err = fmt.Errorf("Unsupported function code %d", n.functionCode)
			}

			if err != nil {
				fmt.Println("ModbusReadBool read error:", err)
				n.connectionIsInit = false
				if n.handler != nil {
					n.handler.Close()
				}
				n.handler = nil
				n.client = nil
				results = append(results, "error: "+err.Error())
				continue
			} else {
				//val := fmt.Sprintf("%d", bytesToUint16(res))
				//results = append(results, val)
			}
		}

		n.lastValuesReceived = results
		n.outputFlag = true
		n.output[0].Output = "1"
		n.output[1].Output = strings.Join(n.lastValuesReceived, " ,, ")
	}()
}

func atoiDefault(s string, def int) int {
	var i int
	_, err := fmt.Sscanf(s, "%d", &i)
	if err != nil {
		return def
	}
	return i
}

func bytesToUint16(data []byte) uint16 {
	if len(data) < 2 {
		return 0
	}
	return uint16(data[0])<<8 | uint16(data[1])
}

func (n *ModbusReadBoolNode) GetNodeType() string { return n.nodeType }
func (n *ModbusReadBoolNode) GetId() int          { return n.id }
func (n *ModbusReadBoolNode) GetOutput(outName string) *OutputHandle {
	for i := range n.output {
		if n.output[i].Name == outName {
			return &n.output[i]
		}
	}
	return nil
}
func (n *ModbusReadBoolNode) GetInput() []InputHandle { return n.input }

func (n *ModbusReadBoolNode) DestroyToBuildAgain() {
	if n.handler != nil {
		n.handler.Close()
	}
	n.handler = nil
	n.client = nil
	n.connectionIsInit = false
	n.outputFlag = false
	n.lastValuesReceived = nil
}
