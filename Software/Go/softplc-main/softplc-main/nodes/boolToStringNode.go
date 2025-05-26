package nodes

// BToSNode that wor like a logical OR gate
type BToSNode struct {
	id                 int
	nodeType           string
	input              []InputHandle
	output             []OutputHandle
	parameterValueData []string
}

var boolToStringDescription = nodeDescription{
	AccordionName: "Communication",
	PrimaryType:   "LogicalNode",
	Type_:         "BToSNode",
	Display:       "bool To String Node",
	Label:         "bool to string",
	Stretchable:   true,
	Services: []servicesStruct{
		{FriendlyName: "testFN", NameServices: []string{"testNS1", "testNS2"}},
	},
	SubServices: []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "bool", Name: "x"},
	},
	Output: []dataTypeNameStruct{{DataType: "bool", Name: "Output"}},
}

func init() {
	RegisterNodeCreator("BToSNode", func() (Node, error) {
		return &BToSNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, boolToStringDescription)
}

func (n *BToSNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = "0"
		return
	}
	for _, in := range n.input {
		if *in.Input == "1" && (n.parameterValueData[0] == n.parameterValueData[1]) {
			n.output[0].Output = "1"
			return
		}
	}
	n.output[0].Output = "0"
	/*
		if *n.input[0].Input > *n.input[1].Input {
			n.output[0].Output = 1
		} else {
			n.output[0].Output = 0
		}*/
}

func (n *BToSNode) GetNodeType() string {
	return n.nodeType
}

func (n *BToSNode) GetId() int {
	return n.id
}

func (n *BToSNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *BToSNode) GetInput() []InputHandle {
	return n.input
}

func (n *BToSNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
	n.parameterValueData = parameterValueData_
}
