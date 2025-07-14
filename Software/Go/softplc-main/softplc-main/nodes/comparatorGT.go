package nodes

import "strconv"

// GTNode that wor like a logical greater than
type GTNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
}

var gtDescription = nodeDescription{
	AccordionName: "Comparator",
	PrimaryType:   "LogicalNode",
	Type_:         "GTNode",
	Display:       "GT Node",
	Label:         "x > y",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "value", Name: "x"},
		{DataType: "value", Name: "y"},
	},
	Output: []dataTypeNameStruct{{DataType: "bool", Name: "Output"}},
}

func init() {
	RegisterNodeCreator("GTNode", func() (Node, error) {
		return &GTNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, gtDescription)
}

func (n *GTNode) ProcessLogic() {
	if n.input == nil || n.input[0].Input == nil || n.input[1].Input == nil {
		n.output[0].Output = "0"
		return
	}

	val0 := *n.input[0].Input
	val1 := *n.input[1].Input

	f0, err0 := strconv.ParseFloat(val0, 64)
	f1, err1 := strconv.ParseFloat(val1, 64)

	if err0 == nil && err1 == nil {
		if f0 > f1 {
			n.output[0].Output = "1"
		} else {
			n.output[0].Output = "0"
		}
	} else {
		if val0 > val1 {
			n.output[0].Output = "1"
		} else {
			n.output[0].Output = "0"
		}
	}
}

func (n *GTNode) GetNodeType() string {
	return n.nodeType
}

func (n *GTNode) GetId() int {
	return n.id
}

func (n *GTNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *GTNode) GetInput() []InputHandle {
	return n.input
}

func (n *GTNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
func (n *GTNode) DestroyToBuildAgain() {

}
