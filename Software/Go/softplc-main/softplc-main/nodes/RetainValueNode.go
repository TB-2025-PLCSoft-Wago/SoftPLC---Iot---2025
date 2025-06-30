package nodes

// RetainValueNode : Provide single string from multiple
type RetainValueNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
}

var retainValueDescription = nodeDescription{
	AccordionName: "Handling value",
	PrimaryType:   "LogicalNode",
	Type_:         "RetainValueNode",
	Display:       "Retain Value Node",
	Label:         "Retain value",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "bool", Name: "pass"},
		{DataType: "value", Name: "str"},
	},
	Output: []dataTypeNameStruct{{DataType: "value", Name: "str"}},
}

func init() {
	RegisterNodeCreator("RetainValueNode", func() (Node, error) {
		return &RetainValueNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, retainValueDescription)
}

func (n *RetainValueNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = ""
		return
	}
	if *n.input[0].Input == "1" {
		n.output[0].Output = *n.input[1].Input
		return
	} else {
		n.output[0].Output = ""
	}
}

func (n *RetainValueNode) GetNodeType() string {
	return n.nodeType
}

func (n *RetainValueNode) GetId() int {
	return n.id
}

func (n *RetainValueNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *RetainValueNode) GetInput() []InputHandle {
	return n.input
}

func (n *RetainValueNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
func (n *RetainValueNode) DestroyToBuildAgain() {

}
