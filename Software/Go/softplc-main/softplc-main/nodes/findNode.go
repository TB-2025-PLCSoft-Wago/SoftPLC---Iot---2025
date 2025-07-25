// Package nodes This package contains the definition of the FindNode struct, which is a node that performs the NOT operation on 1 input signals.
package nodes

import (
	"SoftPLC/serverResponse"
	"strconv"
	"strings"
)

// FindNode struct for a one input Find node
type FindNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
}

var findDescription = nodeDescription{
	AccordionName: "Handling value",
	PrimaryType:   "LogicalNode",
	Type_:         "FindNode",
	Display:       "Find Node",
	Label:         "Find",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "value", Name: "strWhereToSearch"},
		{DataType: "value", Name: "strToSeek"},
		{DataType: "value", Name: "iStart"}, //The index position within itfString from where the search starts
	},
	Output: []dataTypeNameStruct{{DataType: "bool", Name: "isFind"}},
}

func init() {
	RegisterNodeCreator("FindNode", func() (Node, error) {
		return &FindNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, findDescription)
}

func (n *FindNode) ProcessLogic() {
	n.output[0].Output = "0"
	if n.input == nil {
		return
	}
	if n.input[0].Input == nil || n.input[1].Input == nil {
		return
	}
	strWhereToSearch := *n.input[0].Input
	strToSearch := *n.input[1].Input
	valueStr := "0"
	if n.input[2].Input != nil {
		valueStr = *n.input[2].Input
	}
	start, _ := strconv.ParseInt(valueStr, 10, 64)
	if start >= int64(len(strWhereToSearch)) {
		strWhereToSearch = ""
	} else {
		strWhereToSearch = strWhereToSearch[start:]
	}
	if strings.Contains(strWhereToSearch, strToSearch) {
		n.output[0].Output = "1"
	}
}
func (n *FindNode) GetNodeType() string {
	return n.nodeType
}

func (n *FindNode) GetId() int {
	return n.id
}

func (n *FindNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			valueStr := "0"
			if n.input[2].Input != nil {
				valueStr = *n.input[2].Input
			}
			_, err := strconv.ParseInt(valueStr, 10, 64)
			if err != nil {
				serverResponse.ResponseProcessGraph = "Find - input start : " + valueStr + ", wrong format"
			}
			return &n.output[i]
		}
	}
	return nil
}

func (n *FindNode) GetInput() []InputHandle {
	return n.input
}

func (n *FindNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
func (n *FindNode) DestroyToBuildAgain() {

}
