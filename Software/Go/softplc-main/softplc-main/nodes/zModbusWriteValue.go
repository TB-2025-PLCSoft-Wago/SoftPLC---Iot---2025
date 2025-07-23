package nodes

import (
	"SoftPLC/server"
	"fmt"
	"github.com/goburrow/modbus"
	"strconv"
	"strings"
	"time"
)

type ModbusWriteValueNode struct {
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
	//functionCode 		int
}

var modbusWriteValueDescription = nodeDescription{
	AccordionName: "Communication",
	PrimaryType:   "LogicalNode",
	Type_:         "ConfigurableNodeModbusWriteValue",
	Display:       "Modbus Write Value (0x06)",
	Label:         "Modbus Write Value",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "bool", Name: "xEnable"},
		{DataType: "value", Name: "UnitID"},
		{DataType: "value", Name: "Addresses"},
		{DataType: "value", Name: "NewValues"},
	},
	Output: []dataTypeNameStruct{
		{DataType: "bool", Name: "xDone"},
		{DataType: "value", Name: "ValuesReceived"},
	},
	ParameterNameData: []string{"host", "port"},
}

func init() {
	RegisterNodeCreator("ConfigurableNodeModbusWriteValue", func() (Node, error) {
		return &ModbusWriteValueNode{
			id:       -1,
			nodeType: "",
		}, nil
	}, modbusWriteValueDescription)
}

func (n *ModbusWriteValueNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
	n.parameterValueData = parameterValueData_
	n.connectionIsInit = false
}

func (n *ModbusWriteValueNode) initConnection(unitID byte) error {
	if n.connectionIsInit {
		return nil
	}

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

func (n *ModbusWriteValueNode) ProcessLogic() {
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

		newValue := []string{"0"}
		if n.input[3].Input != nil {
			newValue = strings.Split(*n.input[3].Input, " ,, ")
		}
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic in ModbusWriteValue:", r)
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
		var err error
		if !n.connectionIsInit {
			if err = n.initConnection(unitID); err != nil {
				fmt.Println("ModbusWriteValue connection error:", err)
				n.output[0].Output = "0"
				n.output[1].Output = "connection error"
				//return
			}
		}
		if n.handler == nil || n.client == nil || !n.connectionIsInit {
			fmt.Println("Handler not ready")
			n.output[0].Output = "0"
			n.output[1].Output = "Handler not ready"
			return
		}
		/* prepare to send */
		results := []string{}
		if len(addresses) == len(newValue) {
			server.SendToWebSocket("modbus value : send the corresponding value to each respective address")
			for i, addrStr := range addresses {
				var res []byte
				address := uint16(atoiDefault(addrStr, 0))

				// Convert the first element of the slice to uint16
				var newValueConvert uint16
				if value, err2 := strconv.ParseUint(newValue[i], 10, 16); err2 == nil {
					newValueConvert = uint16(value)
				}
				//send
				res, err = n.client.WriteSingleRegister(address, newValueConvert)
				if err != nil {
					fmt.Println("ModbusWriteValue write error:", err)
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
			server.SendToWebSocket("modbus value : send the corresponding value to each respective address and then send the remaining ValuesReceived starting from address " + addresses[len(addresses)-1])
			//send the corresponding value to each respective address
			for i, addrStr := range addresses {
				var res []byte

				address := uint16(atoiDefault(addrStr, 0))

				// Convert the element of the slice to uint16
				var newValueConvert uint16
				if value, err2 := strconv.ParseUint(newValue[i], 10, 16); err2 == nil {
					newValueConvert = uint16(value)
				}
				//send
				res, err = n.client.WriteSingleRegister(address, newValueConvert)
				if err != nil {
					fmt.Println("ModbusWriteValue write error:", err)
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
			//send the remaining ValuesReceived starting from address addresses[len(addresses)-1]
			dataBytes := stringsToBytes(newValue[len(newValue)-len(addresses):])
			address := uint16(atoiDefault(addresses[len(addresses)-1], 0) + 1)
			quantity := uint16(len(newValue) - len(addresses))
			var res []byte
			res, err = n.client.WriteMultipleRegisters(address, quantity, dataBytes)
			if err != nil {
				fmt.Println("ModbusWriteValue write error:", err)
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
			server.SendToWebSocket("modbus value : send the corresponding value to each respective address and then send the last value " + newValue[len(newValue)-1] + "to each address")
			for i, addrStr := range addresses {
				var res []byte

				address := uint16(atoiDefault(addrStr, 0))

				// Convert the first element of the slice to uint16
				var newValueConvert uint16
				if len(newValue) > i {
					if value, err2 := strconv.ParseUint(newValue[i], 10, 16); err2 == nil {
						newValueConvert = uint16(value)
					}
				} else {
					if value, err2 := strconv.ParseUint(newValue[len(newValue)-1], 10, 16); err2 == nil {
						newValueConvert = uint16(value)
					}
				}

				//send
				res, err = n.client.WriteSingleRegister(address, newValueConvert)
				if err != nil {
					fmt.Println("ModbusWriteValue write error:", err)
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
		n.lastValuesReceived = results
		n.outputFlag = true
		if err == nil {
			n.output[0].Output = "1"
		} else {
			n.output[0].Output = "0"
		}
		n.output[1].Output = strings.Join(n.lastValuesReceived, " ,, ")
	}()
}

func (n *ModbusWriteValueNode) GetNodeType() string { return n.nodeType }
func (n *ModbusWriteValueNode) GetId() int          { return n.id }
func (n *ModbusWriteValueNode) GetOutput(outName string) *OutputHandle {
	for i := range n.output {
		if n.output[i].Name == outName {
			return &n.output[i]
		}
	}
	return nil
}
func (n *ModbusWriteValueNode) GetInput() []InputHandle { return n.input }

func (n *ModbusWriteValueNode) DestroyToBuildAgain() {
	if n.handler != nil {
		n.handler.Close()
	}
	n.handler = nil
	n.client = nil
	n.connectionIsInit = false
	n.outputFlag = false
	n.lastValuesReceived = nil
}
func stringsToBytes(strings []string) []byte {
	var result []byte
	for _, str := range strings {
		result = append(result, []byte(str)...)
	}
	return result
}
