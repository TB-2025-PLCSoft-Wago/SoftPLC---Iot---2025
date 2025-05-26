package nodes

// EQNode that wor like a logical equal
type EQNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
}

var eqDescription = nodeDescription{
	AccordionName: "Comparator",
	PrimaryType:   "LogicalNode",
	Type_:         "EQNode",
	Display:       "EQ Node",
	Label:         "x = y",
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
	RegisterNodeCreator("EQNode", func() (Node, error) {
		return &EQNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, eqDescription)
}

func (n *EQNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = "0"
		return
	}

	if *n.input[0].Input == *n.input[1].Input {
		n.output[0].Output = "1"
	} else {
		n.output[0].Output = "0"
	}
}

func (n *EQNode) GetNodeType() string {
	return n.nodeType
}

func (n *EQNode) GetId() int {
	return n.id
}

func (n *EQNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *EQNode) GetInput() []InputHandle {
	return n.input
}

func (n *EQNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
