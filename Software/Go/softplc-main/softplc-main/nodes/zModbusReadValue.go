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
	lastValues         []string
	functionCode       int
}

var modbusReadValueDescription = nodeDescription{
	AccordionName: "Communication",
	PrimaryType:   "LogicalNode",
	Type_:         "ConfigurableNodeModbusReadValue",
	Display:       "Modbus Read Value Node",
	Label:         "Modbus Read Value",
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
		{DataType: "value", Name: "Values"},
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

	/*
		n.functionCode = 4
		if n.input[2].Input != nil {
			n.functionCode = atoiDefault(*n.input[2].Input, 2)
		}*/
	addresses := []string{"0"}
	if n.input[2].Input != nil {
		addresses = strings.Split(*n.input[2].Input, " ,, ")
	}
	quantites := []string{"1"}
	if n.input[3].Input != nil {
		quantites = strings.Split(*n.input[3].Input, " ,, ")
	}
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
		quantity := uint16(atoiDefault(quantites[i], 1))
		var res []byte
		var err error
		//bits := make([]bool, quantity)

		res, err = n.client.ReadInputRegisters(address, quantity)
		//Convert to Value
		/*
			for i2 := 0; uint16(i2) < quantity; i2++ {
				byteIndex := i2 / 8
				bitIndex := uint(i2 % 8)
				bits[i2] = (res[byteIndex] & (1 << bitIndex)) != 0
			}
			for _, bit := range bits {
				if bit {
					results = append(results, "1")
				} else {
					results = append(results, "0")
				}
			}*/

		if err != nil {
			fmt.Println("ModbusReadValue read error:", err)
			n.connectionIsInit = false
			n.handler.Close()
			n.handler = nil
			n.client = nil
			results = append(results, "error: "+err.Error())
			continue
		} else {
			val := fmt.Sprintf("%d", bytesToUint16(res))
			results = append(results, val)
		}
	}

	n.lastValues = results
	n.outputFlag = true
	n.output[0].Output = "1"
	n.output[1].Output = strings.Join(n.lastValues, " ,, ")
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
	n.lastValues = nil
}
