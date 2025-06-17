package nodes

import (
	"fmt"
	"strconv"
	"strings"
)

// CounterNode R1 is priority
type CounterNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
}

const (
	STEP  = 0
	UP    = 1
	DOWN  = 2
	RESET = 3
)

func countDecimals(f float64) int {
	s := strconv.FormatFloat(f, 'f', -1, 64)
	parts := strings.Split(s, ".")
	if len(parts) == 2 {
		return len(parts[1])
	}
	return 0
}

var counterDescription = nodeDescription{
	AccordionName: "Logical gate",
	PrimaryType:   "LogicalNode",
	Type_:         "CounterNode",
	Display:       "Counter Node",
	Label:         "Counter",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "value", Name: "step"},
		{DataType: "bool", Name: "+"}, //up ⬆️
		{DataType: "bool", Name: "-"}, //down ⬇️
		{DataType: "bool", Name: "R"},
	},
	Output: []dataTypeNameStruct{{DataType: "value", Name: "result"}},
}

func init() {
	RegisterNodeCreator("CounterNode", func() (Node, error) {
		return &CounterNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, counterDescription)
}

func (n *CounterNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = "0"
		return
	}
	step, err := strconv.ParseFloat(*n.input[STEP].Input, 64)
	if *n.input[STEP].Input == "" {
		step = 1
		err = nil
	}
	result, err2 := strconv.ParseFloat(n.output[0].Output, 64)
	if err != nil || err2 != nil {
		fmt.Println("error of conversion :", err, " error 2 :", err2)
		n.output[0].Output = "0"
		return
	}
	if *n.input[RESET].Input == "1" {
		n.output[0].Output = "0"
	}

	if *n.input[UP].Input == "1" && *n.input[DOWN].Input == "1" {
		return
	}
	if *n.input[UP].Input == "1" {
		decimalPlaces := countDecimals(step)
		n.output[0].Output = strconv.FormatFloat(result+step, 'f', decimalPlaces, 64)
	}

	if *n.input[DOWN].Input == "1" {
		decimalPlaces := countDecimals(step)
		n.output[0].Output = strconv.FormatFloat(result-step, 'f', decimalPlaces, 64)
	}
}

func (n *CounterNode) GetNodeType() string {
	return n.nodeType
}

func (n *CounterNode) GetId() int {
	return n.id
}

func (n *CounterNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *CounterNode) GetInput() []InputHandle {
	return n.input
}

func (n *CounterNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
func (n *CounterNode) DestroyToBuildAgain() {

}
