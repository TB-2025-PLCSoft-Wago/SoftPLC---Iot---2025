package nodes

// SRValueNode R1 is priority
type SRValueNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
}

var srValueDescription = nodeDescription{
	AccordionName: "Handling value",
	PrimaryType:   "LogicalNode",
	Type_:         "SRValueNode",
	Display:       "SRValue Node",
	Label:         "SRValue",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "bool", Name: "S"},
		{DataType: "bool", Name: "R1"},
		{DataType: "value", Name: "valToS"},
	},
	Output: []dataTypeNameStruct{{DataType: "bool", Name: "Q"}, {DataType: "value", Name: "valOut"}},
}

func init() {
	RegisterNodeCreator("SRValueNode", func() (Node, error) {
		return &SRValueNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, srValueDescription)
}

func (n *SRValueNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = "0"
		n.output[1].Output = ""
		return
	}
	if n.input[0].Input == nil {
		n.output[0].Output = "0"
		n.output[1].Output = ""
		return
	}
	if n.input[1].Input != nil {
		if *n.input[1].Input == "1" {
			n.output[0].Output = "0"
		} else if *n.input[0].Input == "1" {
			n.output[0].Output = "1"
		}

		if n.input[2].Input == nil {
			n.output[1].Output = ""
			return
		} else {
			if *n.input[1].Input == "1" {
				n.output[1].Output = ""
			} else if *n.input[0].Input == "1" {
				n.output[1].Output = *n.input[2].Input
			}
		}
	} else {
		if *n.input[0].Input == "1" {
			if n.input[2].Input == nil {
				n.output[1].Output = ""
				return
			} else {
				n.output[1].Output = *n.input[2].Input
			}
		}
	}

}

func (n *SRValueNode) GetNodeType() string {
	return n.nodeType
}

func (n *SRValueNode) GetId() int {
	return n.id
}

func (n *SRValueNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *SRValueNode) GetInput() []InputHandle {
	return n.input
}

func (n *SRValueNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
func (n *SRValueNode) DestroyToBuildAgain() {
	n.output[0].Output = "0"
	n.output[1].Output = ""

}
