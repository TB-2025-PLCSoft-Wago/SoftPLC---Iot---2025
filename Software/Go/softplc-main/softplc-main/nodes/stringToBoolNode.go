package nodes

import (
	"fmt"
	"strconv"
	"strings"
)

// StringToBoolNode that wor like a logical
type StringToBoolNode struct {
	id                 int
	nodeType           string
	input              []InputHandle
	output             []OutputHandle
	parameterValueData []string
}

var msgToBoolDescription = nodeDescription{
	AccordionName: "Communication",
	PrimaryType:   "LogicalNode",
	Type_:         "StringToBoolNode",
	Display:       "string To bool Node",
	Label:         "string to bool",
	Stretchable:   true,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "value", Name: "str"},
	},
	Output: []dataTypeNameStruct{{DataType: "bool", Name: "x"}},
}

func init() {
	RegisterNodeCreator("StringToBoolNode", func() (Node, error) {
		return &StringToBoolNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, msgToBoolDescription)
}

func (n *StringToBoolNode) ProcessLogic() {
	if n.input == nil {
		for i, _ := range n.parameterValueData {
			n.output[i].Output = "0"
		}
		return
	}
	if n.input[0].Input == nil {
		for i, _ := range n.parameterValueData {
			n.output[i].Output = "0"
		}
		return
	}
	//reset
	for i, _ := range n.output {
		n.output[i].Output = "0"
	}
	inputOrder := strings.Split(*n.input[0].Input, " ,, ")
	/*
		if len(inputOrder) <= 1 {
			for i, _ := range n.parameterValueData {
				if *n.input[0].Input == n.parameterValueData[i] && n.parameterValueData[i] != "" {
					n.output[i].Output = "1"
				} else {
					n.output[i].Output = "0"
				}
			}
		}*/
	if !(len(inputOrder) <= 1) {

		if len(inputOrder) >= len(n.parameterValueData) {
			for i, _ := range n.parameterValueData {
				if inputOrder[i] == n.parameterValueData[i] && n.parameterValueData[i] != "" {
					n.output[i].Output = "1"
				}
			}
		} else {
			for i, _ := range inputOrder {
				if inputOrder[i] == n.parameterValueData[i] && n.parameterValueData[i] != "" {
					n.output[i].Output = "1"
				}
			}
		}
	}
	//input equal one parameter
	for i, _ := range n.parameterValueData {
		if *n.input[0].Input == n.parameterValueData[i] && n.parameterValueData[i] != "" {
			n.output[i].Output = "1"
		}
	}

}

func (n *StringToBoolNode) GetNodeType() string {
	return n.nodeType
}

func (n *StringToBoolNode) GetId() int {
	return n.id
}

func (n *StringToBoolNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *StringToBoolNode) GetInput() []InputHandle {
	return n.input
}

func (n *StringToBoolNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
	n.output = make([]OutputHandle, len(parameterValueData_))
	for i := range parameterValueData_ {
		n.output[i] = OutputHandle{
			Output:   strconv.Itoa(i),
			Name:     fmt.Sprintf("%s%d", output_[0].Name, i),
			DataType: output_[0].DataType,
		}
	}
	n.parameterValueData = parameterValueData_
}

func (n *StringToBoolNode) DestroyToBuildAgain() {

}
