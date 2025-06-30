package nodes

// ConcatNode : Provide single string from multiple
type ConcatNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
}

var concatDescription = nodeDescription{
	AccordionName: "Handling value",
	PrimaryType:   "LogicalNode",
	Type_:         "ConcatNode",
	Display:       "Concat Node",
	Label:         "Concat",
	Stretchable:   true,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "value", Name: "str"},
	},
	Output: []dataTypeNameStruct{{DataType: "value", Name: "str"}},
}

func init() {
	RegisterNodeCreator("ConcatNode", func() (Node, error) {
		return &ConcatNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, concatDescription)
}

func (n *ConcatNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = ""
		return
	}
	n.output[0].Output = ""
	var temp string
	for _, in := range n.input {
		temp = temp + *in.Input
	}

	n.output[0].Output = temp
}

func (n *ConcatNode) GetNodeType() string {
	return n.nodeType
}

func (n *ConcatNode) GetId() int {
	return n.id
}

func (n *ConcatNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *ConcatNode) GetInput() []InputHandle {
	return n.input
}

func (n *ConcatNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
func (n *ConcatNode) DestroyToBuildAgain() {

}
