package nodes

// OrNode that wor like a logical OR gate
type OrNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
}

var orDescription = nodeDescription{
	AccordionName: "Logical gate",
	PrimaryType:   "LogicalNode",
	Type_:         "OrNode",
	Display:       "Or Node",
	Label:         ">=1",
	Stretchable:   true,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input:         []dataTypeNameStruct{{DataType: "bool", Name: "Input"}},
	Output:        []dataTypeNameStruct{{DataType: "bool", Name: "Output"}},
}

func init() {
	RegisterNodeCreator("OrNode", func() (Node, error) {
		return &OrNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, orDescription)
}

func (n *OrNode) ProcessLogic() {
	for _, in := range n.input {
		if *in.Input == "1" {
			n.output[0].Output = "1"
			return
		}
	}
	n.output[0].Output = "0"
}

func (n *OrNode) GetNodeType() string {
	return n.nodeType
}

func (n *OrNode) GetId() int {
	return n.id
}

func (n *OrNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *OrNode) GetInput() []InputHandle {
	return n.input
}

func (n *OrNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
func (n *OrNode) DestroyToBuildAgain() {

}
