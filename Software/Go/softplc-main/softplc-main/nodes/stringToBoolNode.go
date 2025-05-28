package nodes

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
	Label:         "string to bool Node",
	Stretchable:   false,
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
		n.output[0].Output = "0"
		return
	}

	n.output[0].Output = "0"
	for i, in := range n.input {
		if *in.Input == n.parameterValueData[i] && n.parameterValueData[i] != "" {
			n.output[0].Output = "1"
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
	n.parameterValueData = parameterValueData_
}
