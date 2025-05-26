// Package nodes This package contains the definition of the NotNode struct, which is a node that performs the NOT operation on 1 input signals.
package nodes

// NotNode struct for a one input Not node
type NotNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
}

var notDescription = nodeDescription{
	AccordionName: "Logical gate",
	PrimaryType:   "LogicalNode",
	Type_:         "NotNode",
	Display:       "Not Node",
	Label:         "Not",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input:         []dataTypeNameStruct{{DataType: "bool", Name: "Input"}},
	Output:        []dataTypeNameStruct{{DataType: "bool", Name: "Output"}},
}

func init() {
	RegisterNodeCreator("NotNode", func() (Node, error) {
		return &NotNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, notDescription)
}

func (n *NotNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = 0
		return
	}

	if *n.input[0].Input == 0 {
		n.output[0].Output = 1
	} else {
		n.output[0].Output = 0
	}
}
func (n *NotNode) GetNodeType() string {
	return n.nodeType
}

func (n *NotNode) GetId() int {
	return n.id
}

func (n *NotNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *NotNode) GetInput() []InputHandle {
	return n.input
}

func (n *NotNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
