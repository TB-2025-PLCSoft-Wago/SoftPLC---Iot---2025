package nodes

// SRNode R1 is priority
type SRNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
}

var srDescription = nodeDescription{
	AccordionName: "Logical gate",
	PrimaryType:   "LogicalNode",
	Type_:         "SRNode",
	Display:       "SR Node",
	Label:         "SR",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "bool", Name: "S"},
		{DataType: "bool", Name: "R1"},
	},
	Output: []dataTypeNameStruct{{DataType: "bool", Name: "Q"}},
}

func init() {
	RegisterNodeCreator("SRNode", func() (Node, error) {
		return &SRNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, srDescription)
}

func (n *SRNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = "0"
		return
	}
	if n.input[0].Input == nil || n.input[1].Input == nil {
		n.output[0].Output = "0"
		return
	}

	if *n.input[1].Input == "1" {
		n.output[0].Output = "0"
	} else if *n.input[0].Input == "1" {
		n.output[0].Output = "1"
	}
}

func (n *SRNode) GetNodeType() string {
	return n.nodeType
}

func (n *SRNode) GetId() int {
	return n.id
}

func (n *SRNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *SRNode) GetInput() []InputHandle {
	return n.input
}

func (n *SRNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
func (n *SRNode) DestroyToBuildAgain() {

}
