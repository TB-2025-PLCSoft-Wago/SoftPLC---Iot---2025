package nodes

import (
	"SoftPLC/server"
	"fmt"
	"github.com/goburrow/modbus"
	"strconv"
	"strings"
	"time"
)

type ModbusWriteBoolNode struct {
	id                 int
	nodeType           string
	input              []InputHandle
	output             []OutputHandle
	parameterValueData []string
	client             modbus.Client
	handler            *modbus.TCPClientHandler
	connectionIsInit   bool
	outputFlag         bool
	lastValues         []string
	//functionCode 		int
}

var modbusWriteBoolDescription = nodeDescription{
	AccordionName: "Communication",
	PrimaryType:   "LogicalNode",
	Type_:         "ConfigurableNodeModbusWriteBool",
	Display:       "Modbus Write Bool (0x15)",
	Label:         "Modbus Write Bool",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "bool", Name: "xEnable"},
		{DataType: "value", Name: "UnitID"},
		{DataType: "value", Name: "Addresses"},
		//{DataType: "value", Name: "Quantity"},
		{DataType: "value", Name: "NewValue"},
	},
	Output: []dataTypeNameStruct{
		{DataType: "bool", Name: "xDone"},
		{DataType: "value", Name: "Values"},
	},
	ParameterNameData: []string{"host", "port"},
}

func init() {
	RegisterNodeCreator("ConfigurableNodeModbusWriteBool", func() (Node, error) {
		return &ModbusWriteBoolNode{
			id:       -1,
			nodeType: "",
		}, nil
	}, modbusWriteBoolDescription)
}

func (n *ModbusWriteBoolNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
	n.parameterValueData = parameterValueData_
	n.connectionIsInit = false
}

func (n *ModbusWriteBoolNode) initConnection(unitID byte) error {
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

func (n *ModbusWriteBoolNode) ProcessLogic() {
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
		//n.functionCode = 15

		addresses := []string{"0"}
		if n.input[2].Input != nil {
			addresses = strings.Split(*n.input[2].Input, " ,, ")
		}
		/*
			quantites := []string{"1"}
			if n.input[3].Input != nil {
				quantites = strings.Split(*n.input[3].Input, " ,, ")
			}*/
		newValue := []string{"0"}
		if n.input[3].Input != nil {
			newValue = strings.Split(*n.input[3].Input, " ,, ")
		}
		var err error
		if !n.connectionIsInit {
			if err = n.initConnection(unitID); err != nil {
				fmt.Println("ModbusWriteBool connection error:", err)
				n.output[0].Output = "0"
				return
			}
		}
		if n.handler == nil || n.client == nil || !n.connectionIsInit {
			fmt.Println("Handler not ready")
			n.output[0].Output = "0"
			return
		}

		/* prepare to send */
		results := []string{}
		if len(addresses) == len(newValue) {
			server.SendToWebSocket("modbus write bool : send the corresponding value to each respective address")
			for i, addrStr := range addresses {
				var res []byte

				address := uint16(atoiDefault(addrStr, 0))

				// Convert the first element of the slice to uint16
				var newBoolConvert uint16
				if value, err2 := strconv.ParseUint(newValue[i], 10, 16); err2 == nil {
					newBoolConvert = uint16(value)
				}
				//send
				res, err = n.client.WriteSingleCoil(address, newBoolConvert)
				if err != nil {
					fmt.Println("ModbusWriteBool write error:", err)
					n.connectionIsInit = false
					if n.handler != nil {
						n.handler.Close()
					}
					n.handler = nil
					n.client = nil
					results = append(results, "error: "+err.Error())
				} else {
					val := fmt.Sprintf("%d", bytesToUint16(res))
					results = append(results, val)
				}
			}

		} else if len(addresses) < len(newValue) {
			server.SendToWebSocket("modbus write bool : send the corresponding value to each respective address and then send the remaining values starting from address " + addresses[len(addresses)-1])
			//send the corresponding value to each respective address
			for i, addrStr := range addresses {
				var res []byte
				address := uint16(atoiDefault(addrStr, 0))

				// Convert the first element of the slice to uint16
				var dataBytes1 []byte
				dataBytes1, err = stringBitsToByteArray([]string{newValue[i]})
				//send
				res, err = n.client.WriteMultipleCoils(address, 1, dataBytes1)
				if err != nil {
					fmt.Println("ModbusWriteBool write error:", err)
					n.connectionIsInit = false
					if n.handler != nil {
						n.handler.Close()
					}
					n.handler = nil
					n.client = nil
					results = append(results, "error: "+err.Error())
				} else {
					val := fmt.Sprintf("%d", bytesToUint16(res))
					results = append(results, val)
				}
			}
			//send the remaining values starting from address addresses[len(addresses)-1]
			var dataBytes []byte
			var res []byte
			dataBytes, err = stringBitsToByteArray(newValue[len(addresses):])
			address := uint16(atoiDefault(addresses[len(addresses)-1], 0) + 1)
			quantity := uint16(len(newValue) - len(addresses))
			if n.client != nil {
				res, err = n.client.WriteMultipleCoils(address, quantity, dataBytes)
			}
			if err != nil {
				fmt.Println("ModbusWriteBool write error:", err)
				n.connectionIsInit = false
				if n.handler != nil {
					n.handler.Close()
				}
				n.handler = nil
				n.client = nil
				results = append(results, "error: "+err.Error())
			} else {
				val := fmt.Sprintf("%d", bytesToUint16(res))
				results = append(results, val)
			}

		} else {
			server.SendToWebSocket("modbus write bool : send the corresponding value to each respective address and then send the last value " + newValue[len(newValue)-1] + "to each address")
			for i, addrStr := range addresses {
				var res []byte
				address := uint16(atoiDefault(addrStr, 0))

				// Convert the first element of the slice to uint16
				var newBoolConvert uint16
				if len(newValue) > i {
					if value, err2 := strconv.ParseUint(newValue[i], 10, 16); err2 == nil {
						newBoolConvert = uint16(value)
					}
				} else {
					if value, err2 := strconv.ParseUint(newValue[len(newValue)-1], 10, 16); err2 == nil {
						newBoolConvert = uint16(value)
					}
				}

				//send
				res, err = n.client.WriteSingleCoil(address, newBoolConvert)
				if err != nil {
					fmt.Println("ModbusWriteBool write error:", err)
					n.connectionIsInit = false
					if n.handler != nil {
						n.handler.Close()
					}
					n.handler = nil
					n.client = nil
					results = append(results, "error: "+err.Error())
				} else {
					val := fmt.Sprintf("%d", bytesToUint16(res))
					results = append(results, val)
				}
			}
		}

		n.lastValues = results
		n.outputFlag = true
		if err == nil {
			n.output[0].Output = "1"
		} else {
			n.output[0].Output = "0"
		}

		n.output[1].Output = strings.Join(n.lastValues, " ,, ")
	}()
}

func (n *ModbusWriteBoolNode) GetNodeType() string { return n.nodeType }
func (n *ModbusWriteBoolNode) GetId() int          { return n.id }
func (n *ModbusWriteBoolNode) GetOutput(outName string) *OutputHandle {
	for i := range n.output {
		if n.output[i].Name == outName {
			return &n.output[i]
		}
	}
	return nil
}
func (n *ModbusWriteBoolNode) GetInput() []InputHandle { return n.input }

func (n *ModbusWriteBoolNode) DestroyToBuildAgain() {
	if n.handler != nil {
		n.handler.Close()
	}
	n.handler = nil
	n.client = nil
	n.connectionIsInit = false
	n.outputFlag = false
	n.lastValues = nil
}

func packBitsToBytes(bits []bool) []byte {
	byteCount := (len(bits) + 7) / 8
	result := make([]byte, byteCount)

	for i, bit := range bits {
		if bit {
			byteIndex := i / 8
			bitIndex := uint(i % 8)
			result[byteIndex] |= 1 << bitIndex
		}
	}
	return result
}
func stringBitsToByteArray(strBits []string) ([]byte, error) {
	bools := make([]bool, len(strBits))
	for i, s := range strBits {
		if s == "1" {
			bools[i] = true
		} else if s == "0" {
			bools[i] = false
		} else {
			return nil, fmt.Errorf("invalid bit: %s", s)
		}
	}
	return packBitsToBytes(bools), nil
}
